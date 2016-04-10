/*

 */
package main

import (
	"fmt"
	"net/http"
	db "w3af_results_organizer/w3af_sqlite"
)

func main() {
	fmt.Println("Hier")
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8000", mux)

}

func index(w http.ResponseWriter, r *http.Request) {
	vulns := db.GetTest()
	//vulns := db.GetVulnerabilities()
	fmt.Print(vulns, r)
}
