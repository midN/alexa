package helpers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func DebugRequest(req *http.Request, resp *http.Response) {
	if os.Getenv("LOG_LEVEL") == "debug" {
		requestDump, _ := httputil.DumpRequest(req, true)
		responseDump, _ := httputil.DumpResponse(resp, true)
		log.Print(string(requestDump))
		log.Print(string(responseDump))
	}
}
