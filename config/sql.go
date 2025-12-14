package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println(".env file not found, relying on system environment variables")
    }
}
// Connect connects to MySQL, creates tables if they don't exist, and returns the *sql.DB
func Connect() *sql.DB {
	LoadEnv()

user := os.Getenv("DB_USER")
password := os.Getenv("MYSQL_ROOT_PASSWORD")
host := os.Getenv("DB_HOST")
dbname := os.Getenv("MYSQL_DATABASE")
dbPort := os.Getenv("DB_PORT") // <-- read DB_PORT from .env

// Data Source Name with port
dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, dbPort, dbname)

// Open connection
db, err := sql.Open("mysql", dsn)
if err != nil {
    log.Fatal("Failed to connect to DB:", err)
}

    // Test the connection
    if err := db.Ping(); err != nil {
        log.Fatal("Failed to ping DB:", err)
    }

    // Create users table
    usersTable := `CREATE TABLE IF NOT EXISTS users (
        id CHAR(36) PRIMARY KEY,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL
    );`

    if _, err := db.Exec(usersTable); err != nil {
        log.Fatal("Failed to create users table:", err)
    }

    // Create urls table
    urlsTable := `CREATE TABLE IF NOT EXISTS urls (
    short_url VARCHAR(7) PRIMARY KEY,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    clicks INT DEFAULT 0,
    user_id CHAR(36) NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);`
 
counterTable := `CREATE TABLE IF NOT EXISTS counters (
        id BIGINT NOT NULL

)`
 refreshTokenTable := `CREATE TABLE IF NOT EXISTS tokens (
    token VARCHAR(512) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_token FOREIGN KEY (user_id) REFERENCES users(id)
);`

 

    

    if _, err := db.Exec(urlsTable); err != nil {
        log.Fatal("Failed to create urls table:", err)
    }
    if _, err := db.Exec(counterTable); err != nil {
        log.Fatal("Failed to create counters table:", err)
    }
    if _, err := db.Exec(refreshTokenTable); err != nil {
        log.Fatal("Failed to create refresh tokens table:", err)
    }

    fmt.Println("Connected to DB and tables ensured.")
    return db
}
