package url

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Get struct {
	URL    string
	Params map[string]any
}

func (g Get) Request() (*http.Request, error) {
	req, err := http.NewRequest("GET", g.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create the request: %w", err)
	}

	return req, nil
}

type Post struct {
	URL    string
	Params map[string]any
}

func (g Post) Request() (*http.Request, error) {

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(g.Params); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.URL, buf)
	if err != nil {
		return nil, fmt.Errorf("could not create the request: %w", err)
	}

	return req, nil
}
