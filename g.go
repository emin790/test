package main
import (
  "fmt"
  "log"
  "net/http"
)

func swap(a,b *int){					//функция принимает указатели a и b
	*a = *b
	*b = *a
}

func main() {
  http.HandleFunc("/", postHandler)
  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello")
   
   a := 111
	b := 222
	
	fmt.Println("До замены:")
	fmt.Println("a =", a, "b =", b)
	
	swap(&a,&b) 						//передаем в функцию адреса a и b
	
	fmt.Println("После замены:")
	fmt.Println("a =", a, "b =", b)
}