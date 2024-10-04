package server

import (
    "fmt"
    "net/http"
	"encoding/json"

	"github.com/dedobbin/letterbooked/pkg/booksapi"
)

func main_page (w http.ResponseWriter, r *http.Request){
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

func featured(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var titles = [4]string{"The jungle", "The test", "bazinga", "shelldon"}
	var all_docs []interface{}

	for _, title := range titles {
		//TODO: better API???
		query, err := booksapi.QueryTitle(title, 1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			error_response := struct {
				ErrorMsg string `json:"error"`
			}{ErrorMsg : err.Error()}
			json.NewEncoder(w).Encode(error_response)
			return
		}
		
		all_docs = append(all_docs, query.Docs[0])
	}

	json.NewEncoder(w).Encode(all_docs)
}

func Start(){
	http.HandleFunc("/", main_page);
	http.HandleFunc("/featured", featured);
	fmt.Println("starting serber on port 8080");
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println("Error starting serber", err);
	}
}