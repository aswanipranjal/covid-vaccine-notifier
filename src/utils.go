package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func prepareURL(urlStr string, query map[string]string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("prepareURL: error while parsing url : %v", err)
	}
	q := u.Query()
	for key, value := range query {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func DoSecureGet(endpoint, bearerToken string, query, headers map[string]string) ([]byte, error) {
	endpoint, err := prepareURL(endpoint, query)
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("server returned non OK response %s", body)
	}
	return body, nil
}

func DoSecureSend(method, endpoint, authToken string, query map[string]string, requestBody interface{}) ([]byte, error) {
	endpoint, err := prepareURL(endpoint, query)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(payload))
	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("server returned non OK response %s", body)
	}
	return body, nil
}
