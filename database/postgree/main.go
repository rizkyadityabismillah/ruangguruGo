package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "090302"
	dbname = "test_db"
)

func connectDB() (*sql.DB, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db, err := sql.Open("postgres",psqlInfo)
	if err != nil{
		return nil,err
	}
	return db, nil
}

func main(){
	db, err := connectDB()
	if err != nil{
		panic(err)
	}
	_, err = db.Exec(`Create TABLE human(
		id INT,
		name varchar(200) not null,
		age INT not null,
		address varchar(100) not null
	)`)
	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("Afrida maharani")
}