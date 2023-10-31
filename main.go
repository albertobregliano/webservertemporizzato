package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
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

	str := randomString()

	router := http.DefaultServeMux
	router.HandleFunc("/"+str, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
		cancel()
	})

	go http.ListenAndServe(":"+port, router)
	log.Printf("Avviato web server su porta: %s, resterà attivo per %v\n", port, timeout)
	log.Println("http://127.0.0.1" + ":" + port + "/" + str)
	<-ctx.Done()
	log.Println(ctx.Err())
}

func randomString() string {
	b := make([]byte, 200)
	rand.Read(b)
	h := sha256.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
