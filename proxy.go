package proxy

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/proxy", new(Proxy))
}

// Proxy is the k6 extension
type Proxy struct{}

// Set new http proxy
func (*Proxy) SetEnvHTTP(proxy string) {
	os.Setenv("HTTP_PROXY", proxy)
}

// Set new https proxy
func (*Proxy) SetEnvHTTPS(proxy string) {
	os.Setenv("HTTPS_PROXY", proxy)
}

// Get the current https proxy in use
func (*Proxy) GetCurrentEnvHTTP(test string) string {
	return os.Getenv("HTTP_PROXY")
}

// Get the current http proxy in use
func (*Proxy) GetCurrentEnvHTTPS(test string) string {
	return os.Getenv("HTTPS_PROXY")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type options struct {
	Headers []map[string]string `json:"headers"`
	Body    string              `json:"body"`
}

func getProxyUrl(proxyUrl string) *url.URL {
	proxyUrlParsed, err := url.Parse(proxyUrl)
	check(err)
	return proxyUrlParsed
}

func prepareHTTPRequest(method string, targetUrl string, proxyUrl string, options options) *http.Request {
	// prepare body
	jsonBody := []byte(options.Body)
	bodyReader := bytes.NewReader(jsonBody)
	req, reqErr := http.NewRequest(method, targetUrl, bodyReader)
	check(reqErr)

	// prepare headers
	for k, v := range options.Headers[0] {
		req.Header.Set(k, v)
	}

	return req
}

func captureResponseData(res *http.Response) string {
	dumped, err := httputil.DumpResponse(res, true)
	check(err)

	dumpedString := string(dumped)
	return strings.ReplaceAll(dumpedString, "\r", "")
}

func sendRequest(req *http.Request, proxyUrlParsed *url.URL, options options) *http.Response {
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrlParsed), TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	res, err := myClient.Do(req)
	check(err)

	return res
}

// submit a request with the indicated method, url, proxy url and optional options
func (*Proxy) Request(method string, targetUrl string, proxyUrl string, options options) string {
	method = strings.ToUpper(method)
	request := prepareHTTPRequest(method, targetUrl, proxyUrl, options)
	proxyUrlParsed := getProxyUrl(proxyUrl)
	response := sendRequest(request, proxyUrlParsed, options)

	return captureResponseData(response)
}
