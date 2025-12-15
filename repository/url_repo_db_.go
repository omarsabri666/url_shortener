package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/go-sql-driver/mysql"
	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/model/url"
)

type SqlUrlRepository struct {
	db *sql.DB
}

func NewSqlUrlRepository(db *sql.DB) *SqlUrlRepository {
	return &SqlUrlRepository{db: db}
}
func (repo *SqlUrlRepository) CreateURL(req url.URL, c context.Context) error {
	_, err := repo.db.ExecContext(c, "insert into urls (short_url,long_url,user_id) values(?,?,?)", req.ShortUrl, req.LongUrl, req.UserId)
	if err != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// extract key from error message
			msg := mysqlErr.Message
			re := regexp.MustCompile(`for key '(.+)'`)
			matches := re.FindStringSubmatch(msg)
			key := "unknown"
			if len(matches) == 2 {
				key = matches[1]
				log.Println(key)
				if key == "urls.PRIMARY" {
					key = "short url" // map primary key to column name
				}
			}
			return fmt.Errorf("duplicate entry for %s", key)
		}
		log.Println(err)
		return errs.InternalServerError("failed to create url")
	}

	return nil
}
func (repo *SqlUrlRepository) GetURL(shortURL string) (*url.URL, error) {
	var u url.URL
	err := repo.db.QueryRow("select short_url,long_url,user_id , created_at from urls where short_url = ?", shortURL).Scan(&u.ShortUrl, &u.LongUrl, &u.UserId, &u.CreatedAt)

	if err == sql.ErrNoRows {
		// log.Println("No user found with this email")

		return nil, errs.NotFound("url not found")
	}

	if err != nil {
		log.Println(err)
		return nil, errs.InternalServerError("failed to get url")
	}

	return &u, nil
}
func (repo *SqlUrlRepository) IncrementCounter() (int64, error) {
	var count int64

	// Increment counter and get the new value in one query
	_, err := repo.db.Exec(`
    UPDATE counters 
    SET id = LAST_INSERT_ID(id + 1)
`)
	if err != nil {
		log.Println(err)
		return 0, errs.InternalServerError("failed to increment counter")
	}

	// Retrieve the value that was set by LAST_INSERT_ID()
	err = repo.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&count)
	if err != nil {
		log.Println(err)
		return 0, errs.InternalServerError("failed to get counter value")
	}

	return count, nil
}
