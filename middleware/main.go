package main

import (
	"github.com/wlgq2/meerkat"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (template *Template) Render(writer io.Writer, name string, data interface{}, context *meerkat.Context) error {
	return template.templates.ExecuteTemplate(writer, name, data)
}

func main() {
	server := meerkat.New()
	temp1 :=""
	temp2 :=""
	server.GET("/*",func(context *meerkat.Context) error {
		return context.String(http.StatusOK,temp1+temp2)
	},
	func(callback meerkat.HttpHandler)  meerkat.HttpHandler {
		temp2 = "test. "
		return callback
	})

	server.Use(func(callback meerkat.HttpHandler)  meerkat.HttpHandler {
		temp1 = "middleware "
		return callback
	})
	meerkat.LogInstance().Fatalln(server.Start(":8001"))
}
