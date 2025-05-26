package utils

import (
    "bytes"
    "errors"
    "fmt"
    "io"
    "net/http"
    "time"
)

// HTTPClient provides a shared client with timeout settings.
var HTTPClient = &http.Client{
    Timeout: 10 * time.Second,
}

// HTTPRequest performs a generic HTTP request with method, headers, and optional body.
func HTTPRequest(method, url string, headers map[string]string, body []byte) ([]byte, int, error) {
    req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
    if err != nil {
        return nil, 0, err
    }

    for key, val := range headers {
        req.Header.Set(key, val)
    }

    resp, err := HTTPClient.Do(req)
    if err != nil {
        return nil, 0, err
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, resp.StatusCode, err
    }

    if resp.StatusCode >= 400 {
        return responseBody, resp.StatusCode, errors.New(fmt.Sprintf("HTTP error %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
    }

    return responseBody, resp.StatusCode, nil
}

// HTTPGet performs an HTTP GET request.
func HTTPGet(url string, headers map[string]string) ([]byte, int, error) {
    return HTTPRequest(http.MethodGet, url, headers, nil)
}

// HTTPPost performs an HTTP POST request with a JSON body.
func HTTPPost(url string, headers map[string]string, body []byte) ([]byte, int, error) {
    return HTTPRequest(http.MethodPost, url, headers, body)
}

// HTTPPut performs an HTTP PUT request with a JSON body.
func HTTPPut(url string, headers map[string]string, body []byte) ([]byte, int, error) {
    return HTTPRequest(http.MethodPut, url, headers, body)
}

// HTTPPatch performs an HTTP PATCH request with a JSON body.
func HTTPPatch(url string, headers map[string]string, body []byte) ([]byte, int, error) {
    return HTTPRequest(http.MethodPatch, url, headers, body)
}

// HTTPDelete performs an HTTP DELETE request.
func HTTPDelete(url string, headers map[string]string) ([]byte, int, error) {
    return HTTPRequest(http.MethodDelete, url, headers, nil)
}
