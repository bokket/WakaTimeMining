package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type Author struct {
	Name    string
	Subject string
}
type Subject struct {
	Apriori	string
	DBSCAN    string
	KMEANS string
}

type Rooster struct {
	Author   Author
	Subjects []Subject
}
func ShowIndexView(response http.ResponseWriter, request *http.Request) {

	author := Author{
		Name:    "bokket",
		Subject: ".....",
	}
	subjects := []Subject{
		{Apriori: "Apriori", DBSCAN: "DBSCAN",KMEANS: "KMEANS"},
	}
	rooster := Rooster{
		Author:author,
		Subjects: subjects,
	}

	tmpl, err := template.ParseFiles("./views/layout.gohtml", "./views/nav.gohtml", "./views/content.gohtml","./views/footer.gohtml")
	if err != nil {
		fmt.Println("Error " +  err.Error())
	}
	tmpl.Execute(response, rooster)
}
