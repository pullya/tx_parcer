package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pullya/tx_parcer/internal/config"
)

type CloudflareEthClient struct {
}

func CloudflareCall(payload string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", config.CloudflareEndpoint, strings.NewReader(payload))
	if err != nil {
		fmt.Println("CloudflareCall: Error while request create", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("CloudflareCall: Error while send request", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("CloudflareCall: Error while reading response", err)
		return nil, err
	}

	return body, nil
}
