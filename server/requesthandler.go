package main

import "fmt";
import "log";
import "net/http";

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	firstParam, secondParam := parseUrl(r.URL.Path[1:])
	response := ""
	switch(r.Method) {
	case http.MethodGet: 
		if (firstParam == "definition_TD") {
			response = fetchTableDefinition(secondParam)
		} else {
			response = fetchTableContent(firstParam, secondParam);
		}
		enableCors(&w)
		fmt.Fprintf(w, fmt.Sprint(response))
	case http.MethodPost: fmt.Println("Post");
	default: fmt.Println("Default");
	}	
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", 
		"*")
    (*w).Header().Set("Access-Control-Allow-Methods", 
		"POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", 
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}