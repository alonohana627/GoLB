package config

import (
	"encoding/json"
	"io"
	"os"
)

type IteratorConfig struct {
	AlgorithmType string   `json:"algorithm_type"`
	Ips           []string `json:"ips"`
}

// TODO: separate to 2 functions

func Parse(filePath string) (*IteratorConfig, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var body *IteratorConfig
	err = json.Unmarshal(byteValue, &body)
	return body, err
}

func (lc *IteratorConfig) ToString() string {
	ret := ""
	ret += "Algorithm Type: " + lc.AlgorithmType + "\n"
	for _, ip := range lc.Ips {
		ret += "IP: " + ip + "\n"
	}

	return ret
}
