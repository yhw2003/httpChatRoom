package backend

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
	buf := make([]byte, 1024)
	length := 0
	for {
		n, readErr := fp.Read(buf)
		length += n
		if readErr == io.EOF {
			text := string(buf[:length])
			if strings.HasSuffix(requireFile, ".css") {
				w.Header().Add("Content-Type", "text/css;charset=utf-8")
			}
			_, _ = fmt.Fprint(w, text)
			break
		}
	}

}
