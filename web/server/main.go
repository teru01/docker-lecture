package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

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
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	Db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/addusers", addUsers)
	mux.HandleFunc("/showusers", showUsers)
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
	html := `<!DOCTYPE html><html><input type="text" name="name" /><input type="submit" action="/postuser" value="send"/></html>`
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
		usersHTML += `<th>` + user.id + `</th>` + `<td>` + user.name + `</td>`
	}

	html := `<!DOCTYPE html><html><table>` + usersHTML + `</table></html>`
	fmt.Fprintf(w, html)
}
