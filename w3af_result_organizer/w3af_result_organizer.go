/*

 */
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	dbg "github.com/davecgh/go-spew/spew"
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
	m := parseQueryString(r.RequestURI)
	dbg.Dump(m)
	num := db.VulnerabilityUpdate(m["vuln"], m["state"], m["comment"])
	dbg.Dump(num)
}

// parseQueryString takes query string from a url and returns the key values as a map[string]string
func parseQueryString(s string) map[string]string {
	sp := strings.Split(s, "?")
	m := make(map[string]string)
	//there are get parameters at the url
	if len(sp) == 2 {
		params := strings.Split(sp[1], "&")
		for _, v := range params {
			pair := strings.Split(v, "=")
			m[pair[0]] = pair[1]
		}
	}
	return m
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		panic(err)
	}
}
