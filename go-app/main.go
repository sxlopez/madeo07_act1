package main  
  
import (  
	"fmt"  
	"net/http"  
)  
  
func helloHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprintf(w, "¡Hola, mundo desde Go!")  
}  
  
func main() {  
	http.HandleFunc("/", helloHandler)  
	fmt.Println("Servidor escuchando en el puerto 8080...")  
	if err := http.ListenAndServe(":8080", nil); err != nil {  
		fmt.Println(err)  
	}  
}  
