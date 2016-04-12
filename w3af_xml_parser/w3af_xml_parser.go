/*
This
*/
package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	db "github.com/lalelunet/w3af_results_organizer/w3af_sqlite"
)

// define the xml struct informations from the vulnerabilities
type vuln struct {
	XMLName     xml.Name `xml:"vulnerability"`
	Description string   `xml:"description"`
	Name        string   `xml:"name,attr"`
	Severity    string   `xml:"severity,attr"`
	URL         string   `xml:"url,attr"`
	Plugin      string   `xml:"plugin,attr"`
}

type plugin struct {
	Name string `xml:"name,attr"`
}

// main open and parse the xml file.
func main() {
	args := os.Args[1:]

	// check if the neccesary informations are given
	if len(args) < 1 {
		fmt.Printf("Usage: xml_parser path-to-w3af-xml-file.xml")
	}

	r := parseXML(args[0])
	if r {
		fmt.Println("##########\n Parsing from the xml file finished.\n ##########\n ")
	} else {
		fmt.Println("##########\n We found new plugins so we need to reparse the file again\n ##########\n ")
		parseXML(args[0])
	}

}

func parseXML(path string) bool {
	var cat, check string
	// is the path to the xml file correct?
	xmli, err := os.Open(path)
	checkErr(err)
	defer xmli.Close()

	project := getProject(path)
	scan := getScan(project)
	plugins := db.GetPlugins()
	categories := db.GetCategories()
	severities := db.GetSeverities()
	dc := xml.NewDecoder(xmli)

	// first we read and update the categories and the plugis
	for {
		v, _ := dc.Token()
		if v == nil {
			break
		}
		switch r := v.(type) {
		case xml.StartElement:
			_, ok := categories[r.Name.Local]
			if ok {
				// we found a categoriy
				cat = r.Name.Local
				fmt.Printf("Found category %s. Cat ID: %s Start reading plugins\n ", cat, categories[cat]["id"])
			}
			if r.Name.Local == "plugin" {
				var pn plugin
				dc.DecodeElement(&pn, &r)
				_, ok := plugins[pn.Name]
				if ok {
					// plugin already known
					fmt.Printf("Found already known plugin: %s. Doing nothing\n ", pn.Name)
				} else {
					// unknown plugin.
					fmt.Printf("Found not known plugin: %s. Insert it into the database. Categorie is %s (ID:%s)\n ", pn.Name, cat, categories[cat]["id"])
					id := db.PluginInsert(categories[cat]["id"], pn.Name)
					fmt.Printf("Plugin inserted. ID: %v\n ", id)
					check = "reload"
				}
			}

			if r.Name.Local == "vulnerability" {
				// we reach the vuln part of the xml so we are finishred with the categories and plugins
				// exit the function and restart it to parsing just the the vulnerabilities
				if check == "reload" {
					return false
				}
				var vn vuln
				dc.DecodeElement(&vn, &r)
				db.VulnerabilityInsert(vn.URL, vn.Description, severities[vn.Severity], plugins[vn.Plugin]["id"], plugins[vn.Plugin]["cat"], project, scan)

			}

		}

	}
	fmt.Println("vulnerabilities done. Exiting now")
	return true
}

// getProject request all projects from the db and delivers the requested project id back.
// If the requeseted project does not exist it will be created.
func getProject(arg string) (id int64) {
	projects := db.GetProjects()
	// parse file name to define the scan project
	var pr string
	prn := strings.Split(arg, ".")
	if len(prn) >= 2 {
		pr = prn[0]
		fmt.Printf("Project name found. Using: %s \n ", pr)
	} else {
		pr = "undefined"
		fmt.Println("Project name not found. Using: undefined\n ")
	}

	_, ok := projects[pr]
	if ok {
		// project exists
		fmt.Printf("Project already exists.  ID: %d\n ", projects[pr])
		return projects[pr]
	}
	prid := db.ProjectInsert(pr)
	fmt.Printf("New project added.  ID: %d\n ", prid)
	return prid
}

// getScan insert a new scan entry in the db and returns the id
func getScan(project int64) (id int64) {
	sid := db.ScanInsert(project)
	return sid
}

// checkErr check for a error
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
