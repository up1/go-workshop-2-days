package api

import (
	"database/sql"
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
	Age   string `json"age"`
}

func (a *API) AddMember(writer http.ResponseWriter, request *http.Request) {
	member := Member{}
	if request.Method == "POST" {
		member.ID = request.FormValue("id")
		member.Fname = request.FormValue("fname")
		member.Lname = request.FormValue("lname")
		member.Age = request.FormValue("age")
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
