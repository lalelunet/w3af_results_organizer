/*

 */
package main

import (
	"fmt"
	"html/template"
	"net/http"

	db "github.com/lalelunet/w3af_results_organizer/w3af_sqlite"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	vulns := db.GetVulnerabilities()
	tpl := template.Must(template.ParseFiles("templates/vuln.html"))
	err := tpl.ExecuteTemplate(w, "vuln.html", vulns)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		panic(err)
	}
}
