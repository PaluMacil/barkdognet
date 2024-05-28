package main

import (
	"fmt"
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
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get current user: %s\n", err)
		return
	}

	users := []string{currentUser.Username, "postgres"}
	dbName := "postgres"
	targetDB := "barkdog"
	targetUser := "barkadmin"
	useExistingUserCommand := fmt.Sprintf("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_roles WHERE rolname='%s') THEN CREATE ROLE %s LOGIN; END IF; END $$;", targetUser, targetUser)

	foundAdmin := false
	var successfulUser string

	for _, dbUser := range users {
		success, output := runCommand("psql", []string{"-U", dbUser, "-d", dbName, "-c", "\\conninfo"})
		if success {
			successfulUser = dbUser
			fmt.Printf("Found admin user: %s\n", dbUser)
			foundAdmin = true
			break
		} else {
			fmt.Printf("Failure for user: %s\nError: %s\n", dbUser, output)
		}
	}

	if !foundAdmin {
		fmt.Println("Could not establish connection to the DB")
		os.Exit(1)
	}
	commands := [][]string{
		{"-U", successfulUser, "-d", dbName, "-c", useExistingUserCommand},
		{"-U", successfulUser, "-d", dbName, "-c", fmt.Sprintf("SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '%s' AND pid <> pg_backend_pid();", targetDB)},
		{"-U", successfulUser, "-d", dbName, "-c", fmt.Sprintf("DROP DATABASE IF EXISTS %s;", targetDB)},
		{"-U", successfulUser, "-d", dbName, "-c", fmt.Sprintf("CREATE DATABASE %s;", targetDB)},
		{"-U", successfulUser, "-d", dbName, "-c", fmt.Sprintf("ALTER DATABASE %s OWNER TO %s;", targetDB, targetUser)},
	}

	for _, args := range commands {
		success, output := runCommand("psql", args)
		if !success {
			fmt.Printf("Failed to execute command: psql %s\nError: %s\n", strings.Join(args, " "), output)
			os.Exit(1)
		}
	}
	fmt.Printf("Successfully reset database!")
}
