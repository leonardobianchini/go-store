package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "dbteste"
)

func checkErr(err error) {
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	fmt.Println("Successfully connected!")

	content, err := os.Open("base_teste.txt")
	checkErr(err)

	createStatement := `
	DROP TABLE IF EXISTS tbUsers;
	CREATE TABLE tbUsers (
		cpf VARCHAR(16) NOT NULL PRIMARY KEY,
		private VARCHAR(16),
		incompleto VARCHAR(16),
		data_ultima_compra VARCHAR(16),
		ticket_medio VARCHAR(16),
		ticket_ultima_compra VARCHAR(16),
		loja_mais_frequente VARCHAR(18),
		loja_da_ultima_compra VARCHAR(18)
	)`
	_, err = db.Exec(createStatement)
	checkErr(err)

	defer content.Close()

	var usersStore [][]string

	scanner := bufio.NewScanner(content)

	firstLine := true

	var lineStr string

	for scanner.Scan() {
		if !firstLine {
			rgx := regexp.MustCompile(` +`)
			line := rgx.Split(scanner.Text(), -1)

			for i := 0; i < len(line); i++ {
				lineStr = lineStr + "'" + line[i] + "', "
			}

			fmt.Println(lineStr[:len(lineStr)-2])

			insertStatement := `
				INSERT INTO tbUsers(cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, 
					loja_mais_frequente, loja_da_ultima_compra) VALUES (` + lineStr[:len(lineStr)-2] + `)`
			_, err = db.Exec(insertStatement)
			checkErr(err)

			lineStr = ""

			usersStore = append(usersStore, line)
		} else {
			firstLine = false
		}
	}

	fmt.Println(usersStore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
