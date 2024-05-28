import subprocess

def run_command(command):
    try:
        subprocess.run(command, check=True, text=True, stderr=subprocess.PIPE)
        return True, ""
    except subprocess.CalledProcessError as e:
        return False, e.stderr

users = [subprocess.getoutput("whoami"), "postgres"]
dbname = "postgres"
target_db = "barkdog"
target_user = "barkadmin"
use_existing_user_command = f"DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_roles WHERE rolname='{target_user}') THEN CREATE ROLE {target_user} LOGIN; END IF; END $$;"

found_db = False
user = None

for dbuser in users:
    command = ["psql", "-U", dbuser, "-d", dbname, "-c", "\\conninfo"]
    success, error = run_command(command)
    if success:
        user = dbuser
        found_db = True
        break
    else:
        print(f"Failure for user: {dbuser}")
        print(f"Error: {error}")

if found_db:
    commands = [
        ["psql", "-U", user, "-d", dbname, "-c", use_existing_user_command],
        ["psql", "-U", user, "-d", dbname, "-c", f"SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '{target_db}' AND pid <> pg_backend_pid();"],
        ["psql", "-U", user, "-d", dbname, "-c", f"DROP DATABASE IF EXISTS {target_db};"],
        ["psql", "-U", user, "-d", dbname, "-c", f"CREATE DATABASE {target_db};"],
        ["psql", "-U", user, "-d", dbname, "-c", f"ALTER DATABASE {target_db} OWNER TO {target_user};"]
    ]

    for cmd in commands:
        success, error = run_command(cmd)
        if not success:
            print(f"Failed to execute command: {' '.join(cmd)}")
            print(f"Error: {error}")
            break
else:
    print("Could not establish connection to the DB")
