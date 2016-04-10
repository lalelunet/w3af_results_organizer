/*
Package sqlite offers db functionality
*/
package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // path where you clone the git package to // git clone https://github.com/mattn/go-sqlite3.git
)

var db, err = sql.Open("sqlite3", "w3af_results_organizer.db")

// GetPlugins returns all rows from the plugins table
// Format: map[plugin_name]map[key] = val
func GetPlugins() map[string]map[string]string {
	rows, err := db.Query("SELECT plugins_id, plugins_categories, plugins_name, plugins_description FROM plugins")
	checkErr(err)
	m := make(map[string]map[string]string)
	defer rows.Close()
	var (
		id, cat, name, desc string
	)
	for rows.Next() {
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
	var (
		id, name, desc string
	)
	for rows.Next() {
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
	var (
		id, name string
	)
	for rows.Next() {
		rows.Scan(&id, &name)
		m[name] = id
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
	var (
		id   int64
		name string
	)
	for rows.Next() {
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

// GetVulnerabilities returns all vulnerabilities
func GetVulnerabilities() map[string]map[string]string {
	rows, err := db.Query("select vuln_id,vuln_url,vuln_description,vuln_scan,vuln_severity,vuln_plugin,vuln_project,vuln_categorie,vuln_state from vulns")
	//rows, err := db.Query("SELECT * FROM vulns")
	checkErr(err)
	r := make(map[string]map[string]string)
	defer rows.Close()

	var (
		//id,cat,sev,project,state,plugin int64
		id, request, cat, sev, project, state, plugin, url, desc, comment string
	)

	for rows.Next() {

		rows.Scan(&id, &request, &cat, &sev, &project, &state, &plugin, &url, &desc, &comment)
		fmt.Println("das sind doch viele?", id, request, cat, sev, project, state, plugin, url, desc, comment)

		_, ok := r[id]
		if !ok {
			mm := make(map[string]string)
			r[id] = mm
		}

		r[id]["cat"] = cat
		r[id]["request"] = request
		r[id]["sev"] = sev
		r[id]["project"] = project
		r[id]["state"] = state
		r[id]["plugin"] = plugin
		r[id]["url"] = url
		r[id]["desc"] = desc
		r[id]["comment"] = comment
		fmt.Println(r)
		fmt.Println("\n ")
	}
	return r
}

// GetTest dsds
func GetTest() bool {

	rows, err := db.Query("SELECT vuln_id, vuln_url FROM vulns")
	checkErr(err)
	m := make(map[string]string)
	defer rows.Close()
	var (
		id, name string
	)
	for rows.Next() {

		rows.Scan(&id, &name)
		//fmt.Println(id, name)
		m[id] = name
	}
	fmt.Println(m)

	/*
		rows,err := db.Query("select vuln_url,vuln_description,vuln_scan,vuln_severity,vuln_plugin,vuln_project,vuln_categorie,vuln_state from vulns")
		//rows, err := db.Query("SELECT * FROM vulns")
		checkErr(err)
		r := make(map[string]map[string]string)
		defer rows.Close()

		var (
			//id,cat,sev,project,state,plugin int64
			id,request,cat,sev,project,state,plugin,url,desc,comment string
		)

		for rows.Next() {


			rows.Scan(&id, &request, &cat, &sev, &project, &state, &plugin, &url, &desc, &comment)
			fmt.Println("das sind doch viele?",id,request,cat, sev,project, state, plugin, url, desc, comment)

			_, ok := r[id]
			if !ok {
				mm := make(map[string]string)
				r[id] = mm
			}

			r[id]["cat"] = cat
			r[id]["request"] = request
			r[id]["sev"] = sev
			r[id]["project"] = project
			r[id]["state"] = state
			r[id]["plugin"] = plugin
			r[id]["url"] = url
			r[id]["desc"] = desc
			r[id]["comment"] = comment
			fmt.Println(r)
			fmt.Println("\n ")
		}
	*/
	return true
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		panic(err)
	}
}

/*/ Dbmanager connect and read / write to the sqlite db
func Dbmanager() {

	stmt, err := db.Prepare("SELECT name FROM sqlite_master WHERE type = 'table'")

	res, err := stmt.Exec()

	// insert
	stmt, err = db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")

	res, err = stmt.Exec("astaxie", "研发部门", "2012-12-09")

	id, err := res.LastInsertId()

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")

	res, err = stmt.Exec("astaxieupdate", id)

	affect, err := res.RowsAffected()

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")

	res, err = stmt.Exec(id)

	affect, err = res.RowsAffected()

	fmt.Println(affect)

	db.Close()

}
*/
