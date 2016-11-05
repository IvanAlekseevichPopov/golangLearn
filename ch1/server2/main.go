//Server2 - минимальный echo-сервер со счетчиком запросов
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)      //Каждый запрос вызывает обработчик
	http.HandleFunc("/count", counter) //Каждый запрос вызывает обработчик
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

//Обработчик, возвращающий компонент пути запрашиваемого URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()f
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Printf("Matched /.\n")
}

//Счетчик, возвращающий количество сделанных запросов
func counter (w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count %d\n", count)
	mu.Unlock()
	fmt.Printf("Matched /counter.\n")
}
