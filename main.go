package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
// requestdumpTemp = "\n----------------------------REQUEST---------------------------\n%s-------------------------------------------------------"
)

func main() {

	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	idellconnection := make(chan struct{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//	log.Printf(requestdumpTemp, printHeader(r.Header))
		log.Println("recvived request:- ", r.RequestURI)
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "<h1>Hello world<h1>")

	})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		s := <-sigint
		log.Println("stopping server due to: ", s.String())
		time.Sleep(1 * time.Minute)

		tctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()
		if err := server.Shutdown(tctx); err != nil {
			log.Fatal(err.Error())

		}
		close(idellconnection)
	}()

	log.Println("starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}

	<-idellconnection
}

/*func printHeader(h http.Header) (requestdump string) {

	for k, v := range h {
		requestdump += fmt.Sprintf("%s:[%s]\n", k, listToString(v))
	}
	return
}

func listToString(s []string) (ans string) {

	for _, v := range s {
		ans = fmt.Sprintf("%s,%s", v, ans)
	}
	return
}*/
