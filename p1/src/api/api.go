package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type API struct {
	ConnectDB *sql.DB
}
type Member struct {
	ID    string `json"id"`
	Fname string `json"fmane"`
	Lname string `json"lname"`
	Age   int    `json"age"`
}

func (a *API) AddMember(writer http.ResponseWriter, request *http.Request) {
	member := Member{}
	err := json.NewDecoder(request.Body).Decode(&member)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(member)
	insertStatement, err := a.ConnectDB.Prepare("INSERT INTO member(id,fname,lname,age) VALUES( ?, ?, ?, ? )")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer insertStatement.Close()
	_, err = insertStatement.Exec(member.ID, member.Fname, member.Lname, member.Age)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
