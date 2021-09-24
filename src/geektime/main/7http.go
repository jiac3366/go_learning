package main

import (
	"fmt"
	"log"
	"net/http"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("viewHandler111")
}
func viewHandler2(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("viewHandler2")
}


func main() {
	name := "tom"
	//http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/view2/", viewHandler2)

	str := fmt.Sprintf("hello, can I fuck %s", name)
	println(str)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
