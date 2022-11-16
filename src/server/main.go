package main

import (
	"fmt"
	"httpChatRoom/src/server/pkg/backend"
	"io"
	"log"
	"net/http"
	"time"
)

const max = 50

var roomCnt = 0

func main() {
	diction := [max]string{""}
	isUsed := [max]bool{false}
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		//get uuid
		res, UUIDErr := http.Get("https://www.uuidtools.com/api/generate/v1")
		defer func() { _ = res.Body.Close() }()
		if UUIDErr != nil {
			log.Fatal(UUIDErr)
		}
		body, _ := io.ReadAll(res.Body)
		uuid := string(body)[2 : len(string(body))-2]

		//add dictionary
		for i, n := 0, 0; i < max; i++ {
			if isUsed[i] == false {
				isUsed[i] = true
				diction[i] = uuid
				break
			} //when maxed
			if i == max-1 {
				if n > 50 {
					break
				}
				i = 0
				n++
				time.Sleep(1 * time.Second)
			}
		}

		//answer the client
		_, _ = fmt.Fprintf(w, uuid)
	})
	http.HandleFunc("/getDates", func(w http.ResponseWriter, r *http.Request) {
		backend.Dates(w, r, roomCnt)
	})
	http.HandleFunc("/need/", backend.Start)

	//start services
	_ = http.ListenAndServe(":18080", nil)

}
