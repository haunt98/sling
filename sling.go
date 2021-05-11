package sling

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	slinghttp "github.com/haunt98/sling/http"
)

const (
	contentType = "Content-Type"

	defaultQueryLen = 4
)

// Sling is an HTTP request builder and response receiver
type Sling struct {
	httpClient      slinghttp.Client
	method          string
	reqURL          *url.URL
	header          http.Header
	queries         []interface{}
	bodyProvider    slinghttp.BodyProvider
	responseDecoder slinghttp.ResponseDecoder
}

// New return barebone Sling
func New(httpClient slinghttp.Client) *Sling {
	return &Sling{
		httpClient: httpClient,
		header:     make(http.Header),
		queries:    make([]interface{}, 0, defaultQueryLen),
	}
}

// Clone return Sling with same values
// All values are copied except HTTP client so that change in the clone will not affect the original
func (s *Sling) Clone() (*Sling, error) {
	// Copy request URL
	// Feel like a hack
	reqURLStr := s.reqURL.String()
	reqURL, err := url.Parse(reqURLStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	// Copy header
	header := make(http.Header)
	for k, v := range s.header {
		header[k] = v
	}

	// Copy queries
	queries := make([]interface{}, 0, len(s.queries))
	copy(queries, s.queries)

	return &Sling{
		httpClient:      s.httpClient,
		method:          s.method,
		reqURL:          reqURL,
		header:          header,
		queries:         queries,
		bodyProvider:    s.bodyProvider,
		responseDecoder: s.responseDecoder,
	}, nil
}

// HTTP client

// HTTPClient set HTTP client
// Fallback to http.DefaultClient
func (s *Sling) HTTPClient(client slinghttp.Client) *Sling {
	if client == nil {
		s.httpClient = http.DefaultClient
		return s
	}

	s.httpClient = client
	return s
}

// HTTP Method
// https://golang.org/pkg/net/http/#pkg-constants

func (s *Sling) Get(reqURL string) *Sling {
	return s.Method(http.MethodGet, reqURL)
}

func (s *Sling) Head(reqURL string) *Sling {
	return s.Method(http.MethodHead, reqURL)
}

func (s *Sling) Post(reqURL string) *Sling {
	return s.Method(http.MethodPost, reqURL)
}

func (s *Sling) Put(reqURL string) *Sling {
	return s.Method(http.MethodPut, reqURL)
}

func (s *Sling) Patch(reqURL string) *Sling {
	return s.Method(http.MethodPatch, reqURL)
}

func (s *Sling) Delete(reqURL string) *Sling {
	return s.Method(http.MethodDelete, reqURL)
}

func (s *Sling) Connect(reqURL string) *Sling {
	return s.Method(http.MethodConnect, reqURL)
}

func (s *Sling) Options(reqURL string) *Sling {
	return s.Method(http.MethodOptions, reqURL)
}

func (s *Sling) Trace(reqURL string) *Sling {
	return s.Method(http.MethodTrace, reqURL)
}

func (s *Sling) Method(method, reqURL string) *Sling {
	s.method = method
	return s.RequestURL(reqURL)
}

// Header

func (s *Sling) AddHeader(key, value string) *Sling {
	s.header.Add(key, value)
	return s
}

func (s *Sling) SetHeader(key, value string) *Sling {
	s.header.Set(key, value)
	return s
}

// URL

// RequestURL set request url
// Leave empty if error
func (s *Sling) RequestURL(reqURL string) *Sling {
	parsedReqURL, err := url.Parse(reqURL)
	if err != nil {
		return s
	}

	s.reqURL = parsedReqURL
	return s
}

// Query

func (s *Sling) AddQuery(q interface{}) *Sling {
	if q == nil {
		return s
	}

	s.queries = append(s.queries, q)
	return s
}

func (s *Sling) AddQueries(qs ...interface{}) *Sling {
	s.queries = append(s.queries, qs...)
	return s
}

// Body

func (s *Sling) BodyJSON(data interface{}) *Sling {
	if data == nil {
		return s
	}

	return s.BodyProvider(&slinghttp.JSONBodyProvider{
		Data: data,
	})
}

func (s *Sling) BodyForm(data interface{}) *Sling {
	if data == nil {
		return s
	}

	return s.BodyProvider(&slinghttp.FormBodyProvider{
		Data: data,
	})
}

func (s *Sling) BodyProvider(bodyProvider slinghttp.BodyProvider) *Sling {
	if bodyProvider == nil {
		return s
	}

	ct := bodyProvider.ContentType()
	// Ignore empty content type
	if ct == "" {
		return s
	}

	s.SetHeader(contentType, ct)
	s.bodyProvider = bodyProvider

	return s
}

// Request

// Request return HTTP request
func (s *Sling) Request() (*http.Request, error) {
	if err := addQueriesToURL(s.reqURL, s.queries); err != nil {
		return nil, fmt.Errorf("failed to add queries to url: %w", err)
	}

	var body io.Reader
	if s.bodyProvider != nil {
		var err error
		if body, err = s.bodyProvider.Body(); err != nil {
			return nil, fmt.Errorf("failed to provide body: %w", err)
		}
	}

	req, err := http.NewRequest(s.method, s.reqURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	addHeaderToRequest(req, s.header)

	return req, nil
}

func addQueriesToURL(reqURL *url.URL, qs []interface{}) error {
	oldURLValues, err := url.ParseQuery(reqURL.RawQuery)
	if err != nil {
		return fmt.Errorf("failed to parse query: %w", err)
	}

	// Combine old queries with new queries
	for _, q := range qs {
		newURLValues, err := query.Values(q)
		if err != nil {
			return fmt.Errorf("failed to get query values: %w", err)
		}

		for key, values := range newURLValues {
			for _, value := range values {
				oldURLValues.Add(key, value)
			}
		}
	}

	reqURL.RawQuery = oldURLValues.Encode()

	return nil
}

func addHeaderToRequest(req *http.Request, header http.Header) {
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}

// Response

func (s *Sling) ResponseDecoder(rspDecoder slinghttp.ResponseDecoder) *Sling {
	if rspDecoder == nil {
		return s
	}

	s.responseDecoder = rspDecoder

	return s
}

func (s *Sling) Receive(v interface{}) error {
	if s.responseDecoder == nil {
		return nil
	}

	rsp, err := s.Response()
	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code: %d", rsp.StatusCode)
	}

	if rsp.ContentLength == 0 {
		return nil
	}

	defer rsp.Body.Close()

	if err := s.responseDecoder.Decode(rsp, v); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	// https://golang.org/pkg/net/http/#Response
	// https://stackoverflow.com/questions/17948827/reusing-http-connections-in-golang
	if _, err := io.Copy(io.Discard, rsp.Body); err != nil {
		return fmt.Errorf("failed to discard response body: %w", err)
	}

	return nil
}

func (s *Sling) Response() (*http.Response, error) {
	req, err := s.Request()
	if err != nil {
		return nil, err
	}

	rsp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do http request: %w", err)
	}

	return rsp, nil
}
