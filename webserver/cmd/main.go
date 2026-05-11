package main

import (
	"fmt"
	"log"
	"net/http"
)

 func formhandler(w http.ResponseWriter, r *http.Request) {
	if err:= r.ParseForm();err!=nil{
		fmt.Fprintf(w,"Parseform err: %v",err)
		return
	}
	fmt.Fprintln(w,"POST happened succesfully")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	age:=r.FormValue("age")
	fmt.Fprintf(w, "name: %s\n", name)
fmt.Fprintf(w, "address: %s\n", address)
fmt.Fprintf(w, "age: %s\n", age)	
	}
	func homehandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/hello"{
		fmt.Fprint(w,"404 not fond",http.StatusNotFound) 
		return
	}
	if r.Method!="GET"{
		fmt.Fprint(w,"method is not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"Hello!")
 }
 func main(){
	fileserver:=http.FileServer(http.Dir("../static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/hello",homehandler)
	http.HandleFunc("/form",formhandler)
	log.Println("Server running on port 8080")

	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}
 }