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
	fs := http.FileServer(http.Dir("templates/js"))
	mux.Handle("/js/", http.StripPrefix("/js/", fs))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/vuln/status/", vulnchange)
	http.ListenAndServe(":8000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	vulns := db.GetVulnerabilities()
	tpl := template.Must(template.ParseFiles("templates/vuln.html"))
	err := tpl.ExecuteTemplate(w, "vuln.html", vulns)
	checkErr(err)
}

func vulnchange(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("fff%s", r.RequestURI)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		panic(err)
	}
}
