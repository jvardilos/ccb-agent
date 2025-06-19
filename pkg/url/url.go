package url

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Requester interface {
	Request() (*http.Request, error)
}

func GetFromDB(method, url string, params map[string]any) (map[string]any, error) {
	pm, err := ParseMethod(method, url, params)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	r, err := pm.Request()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	c := http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		return nil, fmt.Errorf("could not do the request %w", err)
	}

	// will spit out XML, but have a if nil chain to see if there is a way to decode
	in, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read bytes %w", err)
	}

	fmt.Println(string(in))

	var out map[string]any
	if err := json.Unmarshal(in, &out); err != nil {
		if err = xml.Unmarshal(in, &out); err != nil {
			return nil, fmt.Errorf("could not unmarshal with json or xml: %w", err)
		}
	}

	return out, nil
}
