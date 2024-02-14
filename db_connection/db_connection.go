package db_connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liberopassadorneto/clean-arch/configs"
)

func ConnectDB(conf *configs.Conf) (*sql.DB, error) {
	db, err := sql.Open(conf.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName))
	if err != nil {
		panic(err)
	}
	return db, nil
}
