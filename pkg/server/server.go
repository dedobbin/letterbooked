package server

import (
    "fmt"
    "net/http"

	"github.com/dedobbin/letterbooked/pkg/booksapi"
)

func main_handler (w http.ResponseWriter, r *http.Request){
	query, err := booksapi.QueryTitle("salo", -1)
	var content string
	if err != nil {
		content = "error";
		
	} else {
		content = query.Docs[3].Title;
	}
	w.Header().Set("Content-Type", "text/html");
	fmt.Fprintf(w, "<h1>books</h1>%s", content);
}

func api_handler(w http.ResponseWriter, r *http.Request){
	resp := []byte(`{"status": "ok"}`)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	w.Write(resp)
}

func Start(){
	http.HandleFunc("/", main_handler);
	http.HandleFunc("/api.json", api_handler);
	fmt.Println("starting serber on port 8080");
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println("Error starting serber", err);
	}
}