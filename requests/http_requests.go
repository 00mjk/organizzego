package requests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpRequests struct {
	Username  string
	Password  string
	UserAgent string
}

func (httpRequests HttpRequests) PostToOrganizze(url string, body []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("User-Agent", httpRequests.UserAgent)
	req.SetBasicAuth(httpRequests.Username, httpRequests.Password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(responseBody))
}
