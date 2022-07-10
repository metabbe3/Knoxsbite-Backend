package database

import "github.com/jmoiron/sqlx"

//ConnConfig
type ConnConfig struct {
	Host          string
	DBName        string
	SSLMode       string
	RetryInterval int
	MaxConn       int
	MaxIdle       int
	MaxLifeTime   int
	Type          string
}

//DB Primary Database Object
type DB struct {
	DBConnection  *sqlx.DB
	Conn          string
	RetryInterval int
	MaxConn       int
	MaxIdle       int
	MaxLifeTime   int
	Type          string
}

var MasterDB *DB

func init() {

	// MasterDB = &DB{
	// 	Conn: ,
	// }
}
