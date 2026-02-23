package marklogiclib

import (
	"bytes"
	"fmt"
	"github.com/icholy/digest"
	"net/http"
)

func PutDocument(client *MarklogicClient, uri string, data []byte, collections []string) error {

	url := client.Url + "/v1/documents?uri=" + uri
	for _, coll := range collections {
		url += "&collection=" + coll
	}
	req, err := http.NewRequest(
		"PUT",
		url,
		bytes.NewReader(data),
	)
	if err != nil {
		return err
	}
	if client.HttpClient == nil {
		client.HttpClient = &http.Client{}
	}
	if client.AuthDigest {
		client.HttpClient.Transport = &digest.Transport{
			Username: client.Username,
			Password: client.Password,
		}
	} else {
		req.SetBasicAuth(client.Username, client.Password)
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error putting document to Marklogic: %d", resp.StatusCode)
	}
	fmt.Println("Status:", resp.StatusCode)
	return nil
}
