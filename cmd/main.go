package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	img, err := os.Open("." + r.URL.Path)
	if err != nil {
		w.Write([]byte("unknown file: " + "." + r.URL.Path))
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, img)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/images/", ImageHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
