// Package main offers a convenience script for local development.
// The script attempts to connect to the database using either the 'postgres'
// user or a user named after the current user (useful in case Postgres is installed with brew on a Mac).
//
// The script implements the following operations:
// 1. Attempts to disconnect any users from the application database.
// 2. Drops the application database. WARNING: this operation will lead to the loss of all data in the database.
// 3. Recreates the application database.
// 4. Attempts to recreate the application user.
// 5. If a password is set in the configuration, it attempts to assign that password to the application user.
//
// NOTE: This script is designed to work in environments where the postgres database accepts connections
// without a password from the two potential admin usernames tried (namely, 'postgres' and the current user's name).
// If your environment requires a password or uses a different admin username, it is recommended to drop and recreate
// the database manually and then ensure the user and password for the application are set as needed.
package main

import (
	"fmt"
	"github.com/PaluMacil/barkdognet/configuration"
	"github.com/PaluMacil/barkdognet/logger"
	"log/slog"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func runCommand(command string, args []string) (bool, string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Sprintf("%s\n%s", err, output)
	}
	return true, ""
}

func main() {
	configProvider := configuration.DefaultProvider{}
	config, err := configProvider.Config()
	if err != nil {
		fmt.Printf("loading configuration: %v", err)
		os.Exit(1)
	}
	log := logger.NewLogger(config.Env)

	currentUser, err := user.Current()
	if err != nil {
		log.Error("Failed to get current user", slog.String("error", err.Error()))
		os.Exit(1)
	}

	users := []string{currentUser.Username, "postgres"}
	dbName := "postgres"
	targetDB := config.Database.Database
	targetUser := config.Database.User
	targetPassword := config.Database.Password
	useExistingUserCommand := fmt.Sprintf("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_roles WHERE rolname='%s') THEN CREATE ROLE %s LOGIN; END IF; END $$;", targetUser, targetUser)

	foundAdmin := false
	var successfulUser string

	for _, dbUser := range users {
		success, output := runCommand("psql", []string{"-U", dbUser, "-d", dbName, "-c", "\\conninfo"})
		if success {
			successfulUser = dbUser
			log.Info("found admin db user", slog.String("user", dbUser))
			foundAdmin = true
			break
		} else {
			log.Warn("failure for user", slog.String("user", dbUser), slog.String("error", output))
		}
	}

	if !foundAdmin {
		log.Error("could not establish connection with an admin user; consider manually dropping and creating db")
		os.Exit(1)
	}
	commands := [][]string{
		{"-U", successfulUser, "-d", dbName, "-c", useExistingUserCommand},
		{"-U", successfulUser, "-d", dbName, "-c", fmt.Sprintf(
			`SELECT pg_terminate_backend(pg_stat_activity.pid) 
                FROM pg_stat_activity
                WHERE pg_stat_activity.datname = '%s' AND pid <> pg_backend_pid();`,
			targetDB)},
		{"-U", successfulUser, "-d", dbName, "-c", fmt.Sprintf("DROP DATABASE IF EXISTS %s;", targetDB)},
		{"-U", successfulUser, "-d", dbName, "-c", fmt.Sprintf("CREATE DATABASE %s;", targetDB)},
		{"-U", successfulUser, "-d", dbName, "-c", fmt.Sprintf("ALTER DATABASE %s OWNER TO %s;", targetDB, targetUser)},
	}
	if targetPassword != "" {
		commands = append(commands, []string{
			"-U", successfulUser,
			"-d", dbName,
			"-c", fmt.Sprintf("ALTER USER barkadmin WITH PASSWORD '%s';", targetPassword),
		})
	}

	for _, args := range commands {
		success, output := runCommand("psql", args)
		if !success {
			command := strings.Replace(fmt.Sprintf("psql %s", strings.Join(args, " ")), targetPassword, "****", -1)
			log.Error("failed to execute command", slog.String("command", command), slog.String("error", output))
			os.Exit(1)
		}
	}

	log.Info("Successfully reset database!")
}
