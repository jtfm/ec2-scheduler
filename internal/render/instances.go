package render

import (
	"html/template"
	"net/http"
)

func Instances(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"./internal/render/instances.gohtml",
	)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "instances", nil)
	if err != nil {
		panic(err)
	}
}
