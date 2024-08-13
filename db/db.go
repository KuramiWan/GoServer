package db

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var handler *sql.DB

func Init() {
	proto := flag.String("proto", "tcp", "protocol for api to use (tcp, unix)")
	username := flag.String("username", "pokerogue", "username for database login")
	password := flag.String("password", "pokerogue", "password for database login")
	host := flag.String("host", "127.0.0.1:3306", "database host")
	pokeroguedb := flag.String("dbname", "pokeroguedb", "database name")
	flag.Parse()
	err := execSql(*username, *password, *host, *proto, *pokeroguedb)
	if err != nil {
		panic(err)
	}

}

func execSql(username, password, host, proto, dbname string) error {
	var err error
	handler, err = sql.Open("mysql", username+":"+password+"@"+proto+"("+host+")/"+dbname)
	if err != nil {
		return err
	}
	defer handler.Close()
	handler.SetMaxOpenConns(64)
	handler.SetMaxIdleConns(64)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := handler.QueryContext(ctx, "SELECT uuid,username,hash,salt,registered,lastLoggedIn FROM pokeroguedb.accounts t LIMIT 501")
	if err != nil {
		return err
	}
	defer res.Close()
	if err != nil {
		return err
	}

	for res.Next() {
		var uuid []byte
		var username string
		var hash []byte
		var salt []byte
		var registered string
		var lastLoggedIn string
		err := res.Scan(&uuid, &username, &hash, &salt, &registered, &lastLoggedIn)
		if err != nil {
			panic(err)
		}
		fmt.Println("UUID:", uuid)
		fmt.Println("Username:", username)
		fmt.Println("Hash:", hash)
		fmt.Println("Salt:", salt)
		fmt.Println("Registered:", registered)
		fmt.Println("LastLoggedIn:", lastLoggedIn)
	}

	return res.Err()
}
