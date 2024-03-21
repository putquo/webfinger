package jrd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Jrd struct {
	Subject string `json:"subject"`
	Links   []Link `json:"links"`
}

type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href,omitempty"`
	Issuer string `json:"issuer,omitempty"`
}

type Resource string

func (r Resource) Jrd() (*Jrd, error) {
	res := string(r)
	filePath := filepath.Join("resources", res)
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("resource not found")
		}
		return nil, fmt.Errorf("error reading resource file: %v", err)
	}

	var jrd Jrd
	err = json.Unmarshal(data, &jrd)
	if err != nil {
		return nil, fmt.Errorf("error parsing resource: %v", err)
	}

	return &jrd, nil
}
