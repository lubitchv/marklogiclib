package marklogiclib

import "net/http"

type MarklogicClient struct {
	Url        string
	Username   string
	Password   string
	AuthDigest bool
	HttpClient *http.Client
}
