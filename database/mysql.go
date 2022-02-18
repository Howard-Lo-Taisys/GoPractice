package database

import (
 "database/sql"
 "time"
 "fmt"

 _ "github.com/go-sql-driver/mysql"
)

const (
	UserName     string = "root"
	Password     string = "123456"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "mysql"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func CreateTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT NOT NULL AUTO_INCREMENT,
        numb INT DEFAULT NULL,
        address VARCHAR(64) DEFAULT NULL,
        amount VARCHAR(64) DEFAULT NULL,
        percentage FLOAT DEFAULT NULL,
		PRIMARY KEY ( id )
    );`

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Create Table Fail:", err)
		return err
	}
	fmt.Println("Create Table Success")
	return nil
}

var SqlDB *sql.DB

func init() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	var err error
	SqlDB, err = sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	CreateTable(SqlDB)

	SqlDB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	SqlDB.SetMaxOpenConns(MaxOpenConns)
	SqlDB.SetMaxIdleConns(MaxIdleConns)

	if err := SqlDB.Ping(); err != nil {
		fmt.Println("connection to mysql failed:", err.Error())
		return
	}
}