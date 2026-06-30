package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"
)

func main() {
	var port = "8080"
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on PORT: ", port)
	// Recover will apply globally on every request specific to this main goroutine only
	if err := http.ListenAndServe("localhost:"+port, Recover(http.DefaultServeMux)); err != nil {
		panic(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	type result struct {
		name string
		err  error
	}

	ch := make(chan result)
	id := strings.TrimSpace(r.URL.Query().Get("id"))

	// if u look carefully there is no recover block here for this goroutine
	// go func(id string, ch chan string) {
	// 	var u *User = getUser(id)
	// 	ch <- u.Name
	// }(id, ch)

	go func(id string, ch chan result) {
		defer func() {
			if rec := recover(); rec != nil {
				recoverExecution(w, r)
				ch <- result{
					name: "",
					err:  fmt.Errorf("failuer in goroutine"),
				}
			}
		}()

		var u *User = getUser(id)
		ch <- result{
			name: u.Name,
			err:  nil,
		}
	}(id, ch)

	res := <-ch
	if res.err != nil {
		http.Error(w, "internal goroutine error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res.name + "\n"))
}

type User struct {
	Id   string
	Name string
}

var userData = map[string]*User{
	"1": {Id: "1", Name: "Hasan"},
	"2": {Id: "2", Name: "Umer"},
}

func getUser(id string) *User {
	user, ok := userData[id]
	if !ok {
		return nil
	}

	return user
}

func Recover(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer recoverExecution(w, r)
		next.ServeHTTP(w, r)
	})
}

func recoverExecution(w http.ResponseWriter, r *http.Request) {
	func() {
		if rec := recover(); rec != nil {
			slog.Error("panic recovered",
				"err", rec,
				"stack", string(debug.Stack()),
				"path", r.URL.Path,
			)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}()
}

// Panic Concept :-
// Every goroutine need to recover their own panic separately; if recover correctly happen there will be no process crash.
// Like if u spin up some custom goroutine so u need to handle the recovery their as well.
