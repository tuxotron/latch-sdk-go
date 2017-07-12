package latch

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var client http.Client

func init() {
	client = http.Client{
		Timeout: time.Second * 5,
	}
}

func sendRequest(httpMethod, urlPath string, xpathHeaders, parameters map[string]string, credentials Credentials) *LatchResponse {

	var body string
	url := LatchHost + urlPath
	currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")

	if len(parameters) > 0 {
		body = serializeParameters(parameters)
	}

	authorization := buildAuthorizationHeader(httpMethod, currentTime, urlPath, body, xpathHeaders, credentials)

	req, err := http.NewRequest(httpMethod, url, bytes.NewBufferString(body))
	if err != nil {
		log.Fatal(err)
	}

	addRequestHttpHeaders(httpMethod, authorization, string(currentTime), *req)

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data := LatchResponse{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return &data

}

func addRequestHttpHeaders(httpMethod, authorization, currentTime string, req http.Request) {

	const (
		AuthHeader                 = "Authorization"
		X11PathDateHeader          = "X-11Paths-Date"
		UserAgentHeader            = "User-Agent"
		ContentTypeHeader          = "Content-type"
		ContentTypeFormHeaderValue = "application/x-www-form-urlencoded"
		UserAgentHeaderValue       = "Golang library"
	)

	req.Header.Set(AuthHeader, authorization)
	req.Header.Set(X11PathDateHeader, currentTime)
	req.Header.Set(UserAgentHeader, UserAgentHeaderValue)

	if httpMethod == http.MethodPut || httpMethod == http.MethodPost {
		req.Header.Set(ContentTypeHeader, ContentTypeFormHeaderValue)
	}
}

func buildAuthorizationHeader(httpMethod, currentTime, urlPath, body string, xpathHeaders map[string]string, credentials Credentials) string {

	const (
		ElevenPathAuthHeader = "11PATHS"
		HeaderSeparator      = " "
		RequestSeparator     = "\n"
	)

	var request bytes.Buffer
	var header bytes.Buffer

	request.WriteString(httpMethod)
	request.WriteString(RequestSeparator)
	request.WriteString(currentTime)
	request.WriteString(RequestSeparator)
	request.WriteString(serializeXPathHeaders(xpathHeaders))
	request.WriteString(RequestSeparator)
	request.WriteString(urlPath)
	if len(body) > 0 {
		request.WriteString(RequestSeparator)
		request.WriteString(body)
	}

	header.WriteString(ElevenPathAuthHeader)
	header.WriteString(HeaderSeparator)
	header.WriteString(credentials.Id)
	header.WriteString(HeaderSeparator)
	header.WriteString(sign([]byte(request.String()), []byte(credentials.Secret)))

	return header.String()

}

func sign(request, key []byte) string {

	mac := hmac.New(sha1.New, key)
	mac.Write(request)
	signed := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString([]byte(signed))

}

func serializeParameters(parameters map[string]string) string {

	var serialized bytes.Buffer

	keys := []string{}
	for key, _ := range parameters {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		serialized.WriteString(url.QueryEscape(key))
		serialized.WriteString("=")
		serialized.WriteString(url.QueryEscape(parameters[key]))
		serialized.WriteString("&")
	}

	return strings.TrimSuffix(serialized.String(), "&")

}

func serializeXPathHeaders(xpathHeaders map[string]string) string {

	var serialized bytes.Buffer

	keys := []string{}
	for key, _ := range xpathHeaders {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		serialized.WriteString(strings.ToLower(key))
		serialized.WriteString(":")
		serialized.WriteString(xpathHeaders[key])
		serialized.WriteString(" ")
	}

	return strings.TrimSuffix(serialized.String(), " ")

}
