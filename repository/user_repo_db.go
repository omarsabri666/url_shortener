package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	// "log"

	"github.com/go-sql-driver/mysql"
	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/model/user"
)

type dbUserRepository struct {
	db *sql.DB
}
func NewDbUserRepository(db *sql.DB) *dbUserRepository {
	return &dbUserRepository{db: db}
}

func (repo *dbUserRepository) CreateUser(user user.User) error {
_,err :=	repo.db.Exec("insert into users (id,email,password,first_name,last_name) values(?,?,?,?,?)", user.Id,user.Email,user.Password,user.FirstName,user.LastName)
if err != nil {
	log.Println(err)
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
		   parts := strings.SplitN(key, ".", 2)
        if len(parts) == 2 {
            key = parts[1]
        }
    }
    return fmt.Errorf("duplicate entry for %s", key)
}

	return errs.InternalServerError("failed to create user")
}
	return nil

}


func (repo *dbUserRepository) GetUser(req user.UserSignin)( *user.User,error) {
	var u user.User
	
	err:= repo.db.QueryRow("select id,email,password from users where email = ?" , req.Email).Scan(&u.Id,&u.Email,&u.Password)
	if err == sql.ErrNoRows {
		// log.Println("No user found with this email")

		return nil, errs.Unauthorized("invalid email or password")
	}
	if err != nil {


		return nil, errs.InternalServerError("failed to get user")
	}

	return &u,nil
}