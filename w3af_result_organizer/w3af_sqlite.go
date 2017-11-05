/*
Package sqlite offers db functionality
*/
package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" //
)

var db, err = sql.Open("sqlite3", "w3af_results_organizer.db")

// GetPlugins returns all rows from the plugins table
// Format: map[plugin_name]map[key] = val
func GetPlugins() map[string]map[string]string {
	rows, err := db.Query("SELECT plugins_id, plugins_categories, plugins_name, plugins_description FROM plugins")
	checkErr(err)
	m := make(map[string]map[string]string)
	defer rows.Close()
	for rows.Next() {
		var (
			id, cat, name, desc string
		)
		rows.Scan(&id, &cat, &name, &desc)
		_, ok := m[name]
		if !ok {
			mm := make(map[string]string)
			m[name] = mm
		}
		m[name]["id"] = id
		m[name]["cat"] = cat
		m[name]["name"] = name
		m[name]["desc"] = desc

	}
	return m
}

// GetCategories returns all rows from the categories table
func GetCategories() map[string]map[string]string {
	rows, err := db.Query("SELECT categories_id, categories_name, categories_description FROM categories")
	checkErr(err)
	m := make(map[string]map[string]string)
	defer rows.Close()
	for rows.Next() {
		var (
			id, name, desc string
		)
		rows.Scan(&id, &name, &desc)
		_, ok := m[name]
		if !ok {
			mm := make(map[string]string)
			m[name] = mm
		}
		m[name]["id"] = id
		m[name]["name"] = name
		m[name]["desc"] = desc

	}
	return m
}

// GetSeverities returns all rows from the categories table
func GetSeverities() map[string]string {
	rows, err := db.Query("SELECT severity_id, severity_name FROM severity")
	checkErr(err)
	m := make(map[string]string)
	defer rows.Close()
	for rows.Next() {
		var (
			id, name string
		)
		rows.Scan(&id, &name)
		m[name] = id
	}
	return m
}

// GetSeveritiesById returns all rows from the categories table with id as key
func GetSeveritiesById() map[string]string {
	rows, err := db.Query("SELECT severity_id, severity_name FROM severity")
	checkErr(err)
	m := make(map[string]string)
	defer rows.Close()
	for rows.Next() {
		var (
			id, name string
		)
		rows.Scan(&id, &name)
		m[id] = name
	}
	return m
}

// PluginInsert insert a new Plugin into the db and returns the id from the new dataset
func PluginInsert(pc, pn string) (id int64) {
	stmt, err := db.Prepare("INSERT INTO plugins(plugins_categories, plugins_name) values(?,?)")
	checkErr(err)
	res, err := stmt.Exec(pc, pn)
	checkErr(err)
	id, err = res.LastInsertId()
	checkErr(err)

	return id
}

// GetProjects returns all scan projects
func GetProjects() map[string]int64 {
	rows, err := db.Query("SELECT project_id, project_name FROM projects")
	checkErr(err)
	m := make(map[string]int64)
	defer rows.Close()
	for rows.Next() {
		var (
			id   int64
			name string
		)
		rows.Scan(&id, &name)
		m[name] = id
	}
	return m
}

// ProjectInsert insert a new Plugin into the db and returns the id from the new dataset
func ProjectInsert(pn string) (id int64) {
	stmt, err := db.Prepare("INSERT INTO projects(project_name) values(?)")
	checkErr(err)
	res, err := stmt.Exec(pn)
	checkErr(err)
	id, err = res.LastInsertId()
	checkErr(err)

	return id
}

// ScanInsert insert a new Scan into the db and returns the scan_id
func ScanInsert(sp int64) (id int64) {
	stmt, err := db.Prepare("INSERT INTO scans(scan_project,scan_date) values(?,DateTime('now'))")
	checkErr(err)
	res, err := stmt.Exec(sp)
	checkErr(err)
	id, err = res.LastInsertId()
	checkErr(err)

	return id
}

// VulnerabilityInsert insert a new vulnerability. If it already exist the insert will be ignored
func VulnerabilityInsert(url, desc, sev, plugin, cat string, project, scan int64) (id int64) {
	stmt, err := db.Prepare("INSERT OR IGNORE INTO vulns(vuln_url,vuln_description,vuln_scan,vuln_severity,vuln_plugin,vuln_project,vuln_categorie,vuln_state) values(?,?,?,?,?,?,?,?)")
	checkErr(err)
	res, err := stmt.Exec(url, desc, scan, sev, plugin, project, cat, 0)
	checkErr(err)
	id, err = res.LastInsertId()
	checkErr(err)
	return id
}

// GetVulnerabilities returns vulnerabilities
func GetVulnerabilities(scan, project string) map[int]map[string]string {
	stmt, err := db.Prepare("SELECT vuln_id,vuln_url,vuln_description,vuln_scan,vuln_severity,vuln_plugin,vuln_project,vuln_categorie,vuln_state,vuln_comment from vulns where vuln_state = ?")
	parameter := "0"
	checkErr(err)
	switch {
	case scan != "0":
		stmt, err = db.Prepare("SELECT vuln_id,vuln_url,vuln_description,vuln_scan,vuln_severity,vuln_plugin,vuln_project,vuln_categorie,vuln_state,vuln_comment from vulns where vuln_scan = ?")
		parameter = scan
	case project != "0":
		stmt, err = db.Prepare("SELECT vuln_id,vuln_url,vuln_description,vuln_scan,vuln_severity,vuln_plugin,vuln_project,vuln_categorie,vuln_state,vuln_comment from vulns where vuln_project = ?")
		parameter = project
	}
	cnt := 0
	m := make(map[int]map[string]string)

	rows, err := stmt.Query(parameter)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var (
			id, url, desc, scan, sev, plugin, project, cat, state, comment string
		)
		rows.Scan(&id, &url, &desc, &scan, &sev, &plugin, &project, &cat, &state, &comment)

		// create a map of each Vulnerability
		mm := make(map[string]string)
		cnt++
		m[cnt] = mm
		m[cnt]["id"] = id
		m[cnt]["url"] = url
		m[cnt]["desc"] = desc
		m[cnt]["sev"] = sev
		m[cnt]["plugin"] = plugin
		m[cnt]["project"] = project
		m[cnt]["cat"] = cat
		m[cnt]["state"] = state
		m[cnt]["comment"] = comment
	}

	return m
}

/*

// GetVulnerabilities returns all vulnerabilities
func GetVulnerabilities() map[string]map[string]string {
	rows, err := db.Query("select vuln_id,vuln_url,vuln_description,vuln_scan,vuln_severity,vuln_plugin,vuln_project,vuln_categorie,vuln_state,vuln_comment from vulns")
	checkErr(err)
	m := make(map[string]map[string]string)
	defer rows.Close()

	for rows.Next() {
		var (
			id, url, desc, scan, sev, plugin, project, cat, state, comment string
		)
		rows.Scan(&id, &url, &desc, &scan, &sev, &plugin, &project, &cat, &state, &comment)
		// create a map of each Vulnerability
		_, ok := m[id]
		if !ok {
			mm := make(map[string]string)
			m[id] = mm
		}

		m[id]["url"] = url
		m[id]["desc"] = desc
		m[id]["sev"] = sev
		m[id]["plugin"] = plugin
		m[id]["project"] = project
		m[id]["cat"] = cat
		m[id]["state"] = state
		m[id]["comment"] = comment
	}

	return m
}


*/

// GetScanData returns project and details from the scan
func GetScanData(scan string) map[string]string {
	var (
		date, name, id string
	)

	err := db.QueryRow("select scan_date, project_name, project_id from scans sc inner join projects pr on sc.scan_project =  pr.project_id where scan_id = ?", scan).Scan(&date, &name, &id)
	checkErr(err)
	m := make(map[string]string)
	m["date"] = date
	m["name"] = name
	m["id"] = id

	return m
}

// VulnerabilityUpdate update a vulnerability. (vuln, state, desc string) int64
func VulnerabilityUpdate(vuln, state, comment string) int64 {
	stmt, err := db.Prepare("UPDATE vulns set vuln_comment=?, vuln_state=? where vuln_id=?")
	checkErr(err)
	ret, err := stmt.Exec(comment, state, vuln)
	checkErr(err)
	num, err := ret.RowsAffected()
	checkErr(err)

	return num
}

func checkErr(err error) {
	if err != nil {
		// empty rows in queries should not panic
		if err != sql.ErrNoRows {
			fmt.Printf("Error: %v", err)
			panic(err)
		}

	}
}
