package booksapi

import (
    "fmt"
    "net/http"
)

func handler (w http.ResponseWriter, r *http.Request){
	query, err := QueryTitle("salo", -1)
	var content string
	if err != nil {
		content = "error";
		
	} else {
		content = query.Docs[3].Title;
	}
	w.Header().Set("Content-Type", "text/html");
	fmt.Fprintf(w, "<h1>books</h1>%s", content);
}

func Start_server(){
	http.HandleFunc("/", handler);
	fmt.Println("starting serber on port 8080");
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println("Error starting serber", err);
	}
}