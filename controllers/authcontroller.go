package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInput struct {
	Username string
	Password string
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}

}

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/index.html")
	CheckError(err)
	temp.Execute(w, nil)

}
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/login.html")
		CheckError(err)
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		Userinput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
		fmt.Println(Userinput)
		// db := config.DBConn()
		// db.AutoMigrate(&Product{})

	}
}
