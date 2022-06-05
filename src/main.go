package main
import ("fmt"
        "net/http"
        "html/template")

func index(w http.ResponseWriter, r *http.Request){
  tmpl, _ := template.ParseFiles("views/index.html")
  tmpl.Execute(w, "hello")
}

func contacts(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "contacts page")
  }


func handleRequest(){
//  fmt.Println("Go Rulit!")
  http.HandleFunc("/", index)
  http.HandleFunc("/contacts", contacts)
  http.ListenAndServe(":8080",nil)
}

func main(){
  handleRequest()
}
