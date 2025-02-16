package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

const (
	url = "http://localhost:7001"

	authEndpoint     = "/api/auth"
	buyMerchEndpoint = "/api/buy"
	sendCoinEndpoint = "/api/sendCoin"
)

type authResponse struct {
	Token string `json:"token"`
}

func Test_BuyMerch(t *testing.T) {
	creds := `{"login": "login1", "password": "password1"}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", url, authEndpoint), strings.NewReader(creds))
	if err != nil {
		t.Fatalf("failed to create request: %s", err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("failed to send %s request: %s", authEndpoint, err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got unexpected status code %s in %s response", resp.Status, authEndpoint)
	}

	var authResp authResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Fatalf("failed to decode body: %s", err)
	}

	request, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s/%s", url, buyMerchEndpoint, "cup"), strings.NewReader(creds))
	if err != nil {
		t.Fatalf("failed to create request: %s", err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authResp.Token))

	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("failed to send %s request: %s", buyMerchEndpoint, err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got unexpected status code %s in %s response", resp.Status, buyMerchEndpoint)
	}

	// todo: validate via /api/info
}

func Test_SendCoin(t *testing.T) {
	creds := `{"login": "login2", "password": "password2"}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", url, authEndpoint), strings.NewReader(creds))
	if err != nil {
		t.Fatalf("failed to create request: %s", err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("failed to send %s request: %s", authEndpoint, err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got unexpected status code %s in %s response", resp.Status, authEndpoint)
	}

	var authResp authResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Fatalf("failed to decode body: %s", err)
	}

	body := struct {
		ToUser string `json:"toUser"`
		Amount int64  `json:"amount"`
	}{
		ToUser: "password1",
		Amount: 1,
	}

	out, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to encode body: %s", err)
	}

	request, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", url, sendCoinEndpoint), bytes.NewBuffer(out))
	if err != nil {
		t.Fatalf("failed to create request: %s", err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authResp.Token))
	request.Header.Set("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("failed to send %s request: %s", sendCoinEndpoint, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got unexpected status code %s in %s response", resp.Status, sendCoinEndpoint)
	}
}
