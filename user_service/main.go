//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"shared/repository"

	_ "shared/startup"
)

func main() {
	GET("", func(i int) string {
		return "asdsadsa"
	})

	var d repository.Repository

	_ = d

	return

	// cfg := repository.NewCon

	// cfg.SetHost("HOST").
	// 	SetDbname("DBNAME").
	// 	SetPort(5432).
	// 	SetUser("USER").
	// 	SetPassword("123456")

	m := http.ServeMux{}

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET / 1" + r.PathValue("id")))
	})

	m.HandleFunc("GET /user/...", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET user"))
	})

	m.HandleFunc("GET /user/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET user " + r.PathValue("id")))
	})

	m.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST"))
	})

	http.ListenAndServe(":8080", &m)

	GET("", func(w http.ResponseWriter, args ...any) {
	})

	// log.Println(cfg)
}

type Any any

func GET(pat string, fns ...Any) {
	fnf := func() {
		// var actions []func(w http.ResponseWriter, r *http.Request)

		for _, fn := range fns {

			fnValue := reflect.ValueOf(fn)

			fnType := fnValue.Type()

			if fnType.Kind() != reflect.Func {
				err := fmt.Sprintf("type not supported %v", fnType)
				panic(err)
			}

			outsCount := fnType.NumOut()
			_ = outsCount

			arguments(fnType)

		}

		// http.HandleFunc(pat, func(w http.ResponseWriter, r *http.Request) {
		// })
	}

	fnf()
}

func arguments(arg reflect.Type) {
	argsCount := arg.NumIn()

	for i := 0; i < argsCount; i++ {
		arg := arg.In(i)

		log.Println(arg)
	}
}
