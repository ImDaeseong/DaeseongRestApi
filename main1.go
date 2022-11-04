// main
package main

import (
	"fmt"
	"html"
	"net/http"
)

func handler_EscapeString(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Method : ", request.Method)
	fmt.Println("URL : ", request.URL)
	fmt.Println("Header : ", request.Header)
	fmt.Println("Body : ", request.Body)

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

	http.HandleFunc("/", handler_EscapeString)
	http.HandleFunc("/path", handler_Path)
	http.HandleFunc("/param", handler_Param)
	http.ListenAndServe(":8080", nil)

	fmt.Println("http://localhost:8080")
	fmt.Println("http://localhost:8080/path")
	fmt.Println("http://localhost:8080/param?id=아이디&password=비밀번호")
}
