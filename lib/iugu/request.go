package iugu

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
)

func getAuthorization() string {
	token := "d286f8ce21ad448ccbfc028b9344fe8d"
	buffer := []byte(token)
	converted := base64.StdEncoding.EncodeToString(buffer)
	return "Basic" + " " + converted
}

func doRequest(method string, path string, body io.Reader) ([]byte, error) {
	scheme := "https"
	host := "api.iugu.com"
	req, err := http.NewRequest(method, scheme+"://"+host+path, body)
	req.Header.Add("Authorization", getAuthorization())
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err == nil {
		data, _ := ioutil.ReadAll(resp.Body)
		return data, nil
	}
	return nil, err
}
