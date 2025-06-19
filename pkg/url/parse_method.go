package url

import "fmt"

func ParseMethod(method, url string, params map[string]any) (Requester, error) {

	switch method {
	case "GET":
		return Get{URL: url, Params: params}, nil
	case "POST":
		return Post{URL: url, Params: params}, nil
	}

	return nil, fmt.Errorf("could not parse the URL Method")
}
