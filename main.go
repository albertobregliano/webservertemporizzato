package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	timeout := *flag.Duration("t", 10*time.Second, "timeout del server web")
	port := *flag.String("p", "8080", "porta tcp da usare")
	flag.Parse()

	log.Println(timeout)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	router := http.DefaultServeMux
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
		cancel()
	})

	go http.ListenAndServe(":"+port, router)
	log.Printf("Avviato web server su porta: %s, resterà attivo per %v\n", port, timeout)
	<-ctx.Done()
	log.Println(ctx.Err())
}
