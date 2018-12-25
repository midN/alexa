package helpers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GenerateRequest(method string, path string, msg interface{}, returnBody bool) (*http.Response, []byte) {
	var data io.Reader
	if msg == nil {
		data = nil
	} else {
		data = generateRequestData(msg)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest(method, getURL(path), data)
	req.Header.Add("AUTH", getAUTH())
	req.Header.Add("Content-Type", "application/json")

	if returnBody {
		resp := handleRequest(client, req)
		return resp, readBody(resp)
	} else {
		return handleRequest(client, req), nil
	}
}

func readBody(resp *http.Response) []byte {
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func generateRequestData(msg interface{}) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(msg)
	return buffer
}

func handleRequest(client *http.Client, req *http.Request) *http.Response {
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	DebugRequest(req, resp)
	return resp
}

func getURL(path string) string {
	return fmt.Sprintf("https://%s:%s/%s", os.Getenv("HOST"), os.Getenv("PORT"), path)
}

func getAUTH() string {
	return os.Getenv("AUTH")
}
