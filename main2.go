// main
package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

func handler_EscapeString(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(response, "%q", html.EscapeString(request.URL.Path))
	//fmt.Fprintf(response, "%s", request.URL.Path[0:])
}

func handler_Path(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(response, "%s", request.URL.Path)
}

func handler_Param(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(response, "아이디:%s password:%s", request.URL.Query().Get("id"), request.URL.Query().Get("password"))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", handler_EscapeString)
	router.HandleFunc("/path", handler_Path)
	router.HandleFunc("/param", handler_Param)
	http.ListenAndServe(":8080", router)

	fmt.Println("http://localhost:8080")
	fmt.Println("http://localhost:8080/path")
	fmt.Println("http://localhost:8080/param?id=아이디&password=비밀번호")
}
