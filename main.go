package main

import (
	"fmt"
	"net/http"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("Xin chao Hang Nga1"))
	case "POST":
		w.Write([]byte("Post method"))
	case "PUT":
		w.Write([]byte("Put method"))
	case "DELETE":
		w.Write([]byte("Detele method"))
	}
}

func main() {
	http.HandleFunc("/", homeHandle)

	fmt.Println("Server listenning on port 3000 ...")
	fmt.Println(http.ListenAndServe(":3000", nil))
}
