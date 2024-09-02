package hellofunc

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err == nil {
			fmt.Println("Received body: " + string(body))
		}
	}

	w.Write([]byte("HelloFunc says hello " + r.Method))
}
