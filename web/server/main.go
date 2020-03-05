package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type User struct {
	id         int
	name       string
	password   string
	created_at time.Time
}

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	Db, err = sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/addusers", addUsers)
	mux.HandleFunc("/showusers", showUsers)
	mux.HandleFunc("/postuser", postUser)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func addUsers(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html><html>
	<form action="/postuser" method="POST">
		<input type="text" name="name" />
		<input type="submit" value="send"/>
	</form>
	<a href="/showusers">showuser</a>
	</html>`
	fmt.Println("add")
	fmt.Fprintf(w, html)
}

func showUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := Db.Query("SELECT id, name FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	var users []User
	usersHTML := ""
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.id, &user.name); err != nil {
			return
		}
		users = append(users, user)
		usersHTML += `<tr><td>` + user.name + `</td></tr>`
	}

	html := `<!DOCTYPE html><html><table>` + usersHTML + `</table></html>`
	fmt.Fprintf(w, html)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	fmt.Println(r.PostFormValue("name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	query := "INSERT INTO users (name) values (?)"
	stmt, err := Db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	stmt.QueryRow(r.PostFormValue("name"))
}
