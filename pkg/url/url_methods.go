package url

import (
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
	req, err := http.NewRequest("POST", g.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create the request: %w", err)
	}

	return req, nil
}
