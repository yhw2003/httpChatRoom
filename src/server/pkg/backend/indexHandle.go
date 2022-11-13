package backend

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func HtmlServer(w http.ResponseWriter, r *http.Request) {
	//init file path
	_url := r.RequestURI
	requireFile := _url[6:len(_url)]
	requireFile = "./sources/" + requireFile

	//read file
	fp, err := os.Open(requireFile)
	if err != nil {
		fp, err = os.Open(requireFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func() { _ = fp.Close() }()
	//init buffer
	buf := make([]byte, 1024000)
	len := 0
	for {
		n, readErr := fp.Read(buf)
		len += n
		if readErr == io.EOF {
			text := string(buf[:len])
			_, _ = fmt.Fprintf(w, text)
			n = n
			break
		}
	}

}
