package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/types"
)

var ErrNoToken = fmt.Errorf("no token provided")
var ErrUnauthorized = fmt.Errorf("unauthorized")

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteText(w http.ResponseWriter, status int, v string) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)

	_, err := w.Write([]byte(v))
	return err
}

func WriteJS(w http.ResponseWriter, status int, v string) error {
	w.Header().Set("Content-Type", "application/javascript")
	w.WriteHeader(status)

	_, err := w.Write([]byte(v))
	return err
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error response from server: %v", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(data)
}

func ParseRequest(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func ParseQueryParams(r *http.Request, inParams ...types.Param) (map[string]string, error) {
	query := r.URL.Query()

	params := make(map[string]string)

	for _, param := range inParams {
		value := query.Get(param.Name)
		if value == "" && !param.Optional {
			return nil, fmt.Errorf("missing required query parameter: %s", param.Name)
		}

		if value == "" {
			value = param.DefaultValue
		}

		params[param.Name] = value
	}

	return params, nil
}

func ParseJSON(in string, data interface{}) error {
	return json.Unmarshal([]byte(in), data)
}

func HTTPGet(url string, queryParams map[string]string, timeout time.Duration, logger logger.Logger) (*http.Response, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	q := req.URL.Query()

	for key, value := range queryParams {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()
	start := time.Now()
	resp, err := client.Do(req)
	logger.Request("%v GET %s %s", resp.StatusCode, fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path), time.Since(start))
	return resp, err
}

func HTTPPost(url string, body []byte, headers map[string]string, ctx context.Context, logger logger.Logger) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req = req.WithContext(ctx)

	start := time.Now()
	resp, err := client.Do(req)
	logger.Request("%v POST %s %s", resp.StatusCode, url, time.Since(start))
	return err
}

func SetCookie(w http.ResponseWriter, name string, value string, expires time.Duration, httpOnly bool) {
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time.Now().Add(expires),
		HttpOnly: httpOnly,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
	}

	// fmt.Printf("Setting cookie: %+v\n", cookie)

	http.SetCookie(w, &cookie)
}