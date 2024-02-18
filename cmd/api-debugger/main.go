package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	homeTemplate := template.Must(template.ParseFiles("templates/pages/index.html"))
// 	homeTemplate.Execute(w, nil)
// }

// func clickHandler(w http.ResponseWriter, r *http.Request) {
// 	clickTemplate := template.Must(template.ParseFiles("templates/components/click.html"))
// 	clickTemplate.Execute(w, nil)
// }

// type Config struct {
// 	services
// }

type Substruct struct {
	Description string
}

type Test struct {
	A int8
	B string
	C struct {
		D string
		E string
	}
	List []struct {
		Name  string
		Value int8
	}
	Implicit map[string]Substruct
}

type Config struct {
}

func main() {
	// fmt.Println("working")

	// http.HandleFunc("/hello", rootHandler)

	// http.HandleFunc("/", homeHandler)

	// http.HandleFunc("/click", clickHandler)

	// log.Fatal(http.ListenAndServe(":8080", nil))

	// defer fmt.Println("what")

	var test Test
	dat, _ := os.ReadFile("easy.yaml")
	yaml.Unmarshal(dat, &test)
	fmt.Print(test)
}
