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

user := os.Getenv("MYSQLUSER")
password := os.Getenv("MYSQL_ROOT_PASSWORD")
host := os.Getenv("MYSQLHOST")
dbname := os.Getenv("MYSQL_DATABASE")
dbPort := os.Getenv("MYSQLPORT") // <-- read DB_PORT from .env
log.Printf("DEBUG dbPort : %v" , dbPort)
log.Printf("DEBUG dbname : %v" , dbname)
log.Printf("DEBUG host : %v" , host)
log.Printf("DEBUG user : %v" , user)
log.Printf("DEBUG password : %v" , password)


// mysqlUrl := os.Getenv("MYSQL_URL")

// // Data Source Name with port
dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, dbPort, dbname)
// log.Println(dsn)

// mysqlUrl := os.Getenv("MYSQL_URL") // from Railway

// db, err := sql.Open("mysql", mysqlUrl)
// mysqlUrl := os.Getenv("MYSQL_URL")
// mysqlUrl = strings.Trim(mysqlUrl, "\"") // remove quotes
// mysqlUrl = strings.TrimPrefix(mysqlUrl, "mysql://")

// parts := strings.SplitN(mysqlUrl, "@", 2)
// userPass := parts[0]
// hostDb := parts[1]

// hostParts := strings.SplitN(hostDb, "/", 2)
// host := hostParts[0]
// dbName := hostParts[1]

// dsn := userPass + "@tcp(" + host + ")/" + dbName
log.Println(dsn)
log.Printf("DEBUG dsn : %v" , dsn)

db, err := sql.Open("mysql", dsn)
// Open connection
// db, err := sql.Open("mysql", mysqlUrl)
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
