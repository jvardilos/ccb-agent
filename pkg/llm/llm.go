package llm

import (
	"fmt"

	"github.com/jvardilos/ccb-agent/pkg/url"
)

func Ask(prompt string) (map[string]any, error) {
	payload := map[string]any{
		"Model":  "mistral",
		"Prompt": prompt,
	}

	out, err := url.GetFromDB("POST", "http://localhost:11434/api/generate", payload)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return out, nil

}
