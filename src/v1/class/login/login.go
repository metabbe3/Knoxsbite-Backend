package login

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const (

//selectAllUser := `SELECT * from user`
)

type (
	loginRepository struct{}

	UserData struct {
		ID       int    `db:"id"`
		Username string `db:"username"`
		Password string `db:"password"`
		Email    string `db:"email"`
	}
)

func init() {
	timeStart := time.Now()
	log.Println("Init class/login : ", timeStart)
}

func (l loginRepository) SelectUserData(ctx context.Context, tx *sqlx.Tx, userData UserData) error {

	//query := tx.Rebind(selectAllUser)
	return nil
}
