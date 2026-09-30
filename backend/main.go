package main

import(
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "BuyAndSell API.")
}

func main(){
	http.HandleFunc("/", HomeHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}