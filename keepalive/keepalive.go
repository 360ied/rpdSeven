package keepalive

import (
	"fmt"
	"net/http"

	"github.com/ztrue/tracerr"
)

func KeepAlive() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var _, err = fmt.Fprintf(w, "I'm alive")

		if err != nil {
			tracerr.PrintSourceColor(err)
		}
	})

	panic(http.ListenAndServe(":8080", nil))
}
