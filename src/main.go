package main
import (
        "net/http"
        "html/template")

func index(w http.ResponseWriter, r *http.Request){
  tmpl, _ := template.ParseFiles("views/index")
  tmpl.Execute(w, "hello")

}

func products(w http.ResponseWriter, r *http.Request){
  tmpl, _ := template.ParseFiles("views/products")
  tmpl.Execute(w, "hello")

  }

func contacts(w http.ResponseWriter, r *http.Request){
  tmpl, _ := template.ParseFiles("views/contacts")
  tmpl.Execute(w, "hello")

  }

func handleRequest(){
//  fmt.Println("Go Rulit!")
  http.HandleFunc("/", index)
  http.HandleFunc("/products", products)
  http.HandleFunc("/contacts", contacts)
  http.ListenAndServe(":8080",nil)
}

func main(){
http.Handle("/static/",
  http.StripPrefix("/static/",
    http.FileServer(http.Dir("static"))))
    
  handleRequest()

}
