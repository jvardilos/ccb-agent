package url

import "fmt"

func ParseMethod(method, url string) (Requester, error) {

	switch method {
	case "GET":
		return Get{URL: url}, nil
	case "POST":
		return Post{URL: url}, nil
	}

	return nil, fmt.Errorf("could not parse the URL Method")
}
