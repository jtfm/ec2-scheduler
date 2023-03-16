package render

import (
	"html/template"
	"net/http"
)

func Instances(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"./dist/template/instances.gohtml",
	)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "transactions", nil)
	if err != nil {
		panic(err)
	}
}
