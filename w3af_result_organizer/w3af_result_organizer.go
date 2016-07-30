/*

 */
package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
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
	mux.HandleFunc("/vulns/", vulns)
	mux.HandleFunc("/vuln/status/", vulnChange)
	mux.HandleFunc("/scan/result/export", exportScanResult)
	http.ListenAndServe(":8000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	vulns := db.GetVulnerabilities("0", "0")
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tpl.ExecuteTemplate(w, "index.html", vulns)
	//err := tpl.ExecuteTemplate(w, "vuln.html", vulns)
	checkErr(err)
}

func vulns(w http.ResponseWriter, r *http.Request) {
	vulns := db.GetVulnerabilities("0", "0")
	data, err := json.Marshal(vulns)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	//tpl := template.Must(template.ParseFiles("templates/index.html"))
	//err := tpl.ExecuteTemplate(w, "index.html", vulns)
	//err := tpl.ExecuteTemplate(w, "vuln.html", vulns)
	checkErr(err)
}

func vulnChange(w http.ResponseWriter, r *http.Request) {
	m := parseQueryString(r.RequestURI)
	dbg.Dump(m)
	num := db.VulnerabilityUpdate(m["vuln"], m["state"], m["comment"])
	dbg.Dump(num)
}

func exportScanResult(w http.ResponseWriter, r *http.Request) {
	m := parseQueryString(r.RequestURI)
	dbg.Dump(m)
	_, ok := m["scan"]
	if !ok {
		return
	}
	vulns := db.GetVulnerabilities(m["scan"], "0")

	// create a new buffer
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)

	// insert some format stuff to make the cvs looking better
	wr.Write([]string{" ", " ", " ", " ", " "})

	for _, row := range vulns {
		//dat = []string{row["plugin"], row["state"], row["url"]}
		fmt.Println(row["url"])

		wr.Write([]string{" ", row["plugin"], row["state"], row["url"]})
	}
	wr.Flush()

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=report.csv")
	w.Write(b.Bytes())

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
