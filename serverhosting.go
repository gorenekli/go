package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("welcome")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.Handle("/docs/", http.StripPrefix("/docs", http.FileServer(http.Dir("/users/share/doc"))))
	http.Handle("/root/", http.StripPrefix("/root", http.FileServer(http.Dir("/"))))

	var err error
	err = http.ListenAndServe(":8090", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server one closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server one : %s\n", err)
	} else {
		fmt.Printf("no error : %s catched\n", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
