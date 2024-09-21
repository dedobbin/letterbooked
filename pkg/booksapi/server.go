package booksapi

import (
    "fmt"
    "net/http"
	"os"
)

func handler (w http.ResponseWriter, r *http.Request){
	query, err := QueryTitle("salo", -1)
	if err != nil {
		fmt.Println("Query failed!", err)
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "text/html");
	fmt.Fprintf(w, "<h1>%s</h1>", query.Docs[3].Title);
}

func Start_server(){
	http.HandleFunc("/", handler);
	fmt.Println("starting serber on port 8080");
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println("Error starting serber", err);
	}
}