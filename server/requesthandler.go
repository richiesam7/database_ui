package main

import "fmt";
import "log";
import "net/http";

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := ""
	switch(r.Method) {
	case http.MethodGet: fmt.Println("Get"); 
	case http.MethodPost: fmt.Println("Post");
	default: fmt.Println("Default");
	}
    tableName := r.URL.Path[1:];
	queryResponse := executeQuery(tableName, params);
	response := marshallResponse(queryResponse);
	enableCors(&w)
	fmt.Fprintf(w,fmt.Sprint(response))	
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", 
		"*")
    (*w).Header().Set("Access-Control-Allow-Methods", 
		"POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", 
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}