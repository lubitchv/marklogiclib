package marklogiclib

import (
	"bytes"
	"fmt"
	"github.com/icholy/digest"
	"net/http"
	"io"
	"strings"
)

func PutDocument(client *MarklogicClient, uri string, data []byte, collections []string) error {

	url := client.Url + "/v1/documents?uri=" + uri 
	for _, coll := range collections {
		co := strings.Replace(coll, "#", "%23", -1)
		co = strings.Replace(co, "/", "%2F", -1)
		url += "&collection=" + co
	}
	
	fmt.Println(url)
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
		bodyBytes, err := io.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))
		if err != nil {
			fmt.Println("error reading response body: %v", err)
			return fmt.Errorf("error reading response body: %v", err)
		}
		fmt.Println(string(bodyBytes))
	}
	fmt.Println("Status:", resp.StatusCode)
	return nil
}
