package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HttpRequests struct {
	Username  string
	Password  string
	UserAgent string
}

func (httpRequests HttpRequests) PostToOrganizze(url string, body []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.WithError(err).Error("Error while building the POST request")
	}
	req.Header.Set("User-Agent", httpRequests.UserAgent)
	req.SetBasicAuth(httpRequests.Username, httpRequests.Password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("Error while executing the POST request")
	}
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)
}

func (httpRequests HttpRequests) GetFromOrganizze(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.WithError(err).Error("Error while building the GET request")
	}
	req.Header.Set("User-Agent", httpRequests.UserAgent)
	req.SetBasicAuth(httpRequests.Username, httpRequests.Password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("Error while executing the GET request")
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	return responseBody
}
