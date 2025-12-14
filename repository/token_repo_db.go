package repository

import (
	"database/sql"
	"errors"
	"log"
)

type TokenRepositoryDb struct {
	db *sql.DB
}

func NewTokenRepositoryDb(db *sql.DB) *TokenRepositoryDb {
	return &TokenRepositoryDb{db: db}
}

func (r *TokenRepositoryDb) InsertToken(token TokenStruct) error {
	log.Println(token.Exp)
	_, err := r.db.Exec("insert into tokens (token,user_id,expires_at) values(?,?,?)", token.RefreshToken, token.UserId,token.Exp)

	if err != nil {
		return err
	}

	return nil
}
func (r *TokenRepositoryDb) GetToken(token TokenStruct) (*TokenStruct, error) {
	var t TokenStruct


	err := r.db.QueryRow("select token,user_id , expires_at from tokens where token = ? and expires_at > unix_timestamp() and user_id  = ?  ", token.RefreshToken, token.UserId).Scan(&t.RefreshToken, &t.UserId,&t.Exp)
	if err != nil {
		return nil, err
	}
    if err == sql.ErrNoRows {
        return nil, errors.New("invalid or expired refresh token")
    }
		return &t, nil
}
func (r *TokenRepositoryDb) DeleteToken(token TokenStruct) error {
	
	_, err := r.db.Exec("delete from tokens where token = ? and user_id = ?", token.RefreshToken, token.UserId)
	if err != nil {
		return err
	}
	return nil
}