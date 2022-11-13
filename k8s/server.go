package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
);

func main(){
	http.HandleFunc("/", Hello);
	http.HandleFunc("/configmap", ConfigMap);
	http.HandleFunc("/secret", Secret);
	http.ListenAndServe(":80", nil);
}

func Hello(w http.ResponseWriter, r *http.Request){
	name := os.Getenv("NAME");
	age := os.Getenv("AGE");

	fmt.Fprintf(w, "Hello, I'm %s. I'm %s", name, age);
}

func Secret(w http.ResponseWriter, r *http.Request){
	user := os.Getenv("USER");
	password := os.Getenv("PASSWORD");

	fmt.Fprintf(w, "User: %s. Password: %s", user, password);
}

func ConfigMap(w http.ResponseWriter, r *http.Request){
	data, err := ioutil.ReadFile("/go/myfamily/myfamily.txt");
	if err != nil {
		log.Fatalf("Error reading file: ", err);
	}
	fmt.Fprintf(w, "Hy Family: %s.", string(data));
}