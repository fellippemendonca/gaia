package openWeatherMap

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
)

func getAuthorization() string {
	token := ""
	buffer := []byte(token)
	converted := base64.StdEncoding.EncodeToString(buffer)
	return "Basic" + " " + converted
}

func doRequest(method string, path string, body io.Reader) ([]byte, error) {
	scheme := "https"
	host := "api.openweathermap.org"
	req, _ := http.NewRequest(method, scheme+"://"+host+path, body)
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
