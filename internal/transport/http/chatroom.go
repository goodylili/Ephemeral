package http

import (
	"net/http"
)

func ChatPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "path_to_your_html_file.html")
}
