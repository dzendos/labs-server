package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	imgParts := strings.Split(parts[len(parts)-1], ".")
	if len(imgParts) != 2 {
		w.Write([]byte("incorrect file format"))
		return
	}

	ext := imgParts[1]
	w.Header().Set("Content-Type", "image/"+ext)

	b, err := os.ReadFile("./images/" + parts[len(parts)-1])
	if err != nil {
		w.Write([]byte("unknown file"))
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.Write([]byte("unexpected error"))
	}
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
