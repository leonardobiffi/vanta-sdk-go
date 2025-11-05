package vanta

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (v *vanta) ListResources(ctx context.Context, opts ...ListResourcesOption) (*ListResourcesOutput, error) {
	tokenType, token := v.tokenStore.GetToken()
	if token == "" {
		return nil, errors.New("no auth token present")
	}

	// handle query params
	queryParamMap := make(map[string]string)
	for _, opt := range opts {
		opt(queryParamMap)
	}
	queryParamString := ""
	queryParamCount := 0
	for k, v := range queryParamMap {
		separator := ""
		if queryParamCount > 0 {
			separator = "&"
		}
		queryParamString = fmt.Sprintf("%s%s%s=%s", queryParamString, separator, k, v)
		queryParamCount++
	}

	url := fmt.Sprintf("%s/v1/integrations/gitlab/resource-kinds/GitlabRepo/resources", v.baseURL)
	if queryParamCount > 0 {
		url = fmt.Sprintf("%s?%s", url, queryParamString)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build http request: %v", err)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, token))

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %v", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 http response status code (%d), body: %s", resp.StatusCode, string(respBodyBytes))
	}

	var listResourcesOutput *ListResourcesOutput
	if err = json.Unmarshal(respBodyBytes, &listResourcesOutput); err != nil {
		return nil, fmt.Errorf("failed to JSON-decode response body: %v", err)
	}

	return listResourcesOutput, nil
}

func (v *vanta) GetResourceByID(ctx context.Context, id string) (*Resource, error) {
	tokenType, token := v.tokenStore.GetToken()
	if token == "" {
		return nil, errors.New("no auth token present")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/resources/%s", v.baseURL, id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build http request: %v", err)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, token))

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %v", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 http response status code (%d), body: %s", resp.StatusCode, string(respBodyBytes))
	}

	var resource *Resource
	if err = json.Unmarshal(respBodyBytes, &resource); err != nil {
		return nil, fmt.Errorf("failed to JSON-decode response body: %v", err)
	}

	return resource, nil
}
func (v *vanta) UpdateResource(ctx context.Context, id string, input UpdateResourceInput) error {
	tokenType, token := v.tokenStore.GetToken()
	if token == "" {
		return errors.New("no auth token present")
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("failed to JSON-encode request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/integrations/gitlab/resource-kinds/GitlabRepo/resources/%s", v.baseURL, id), io.NopCloser(io.Reader(bytes.NewReader(inputBytes))))
	if err != nil {
		return fmt.Errorf("failed to build http request: %v", err)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, token))

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute http request: %v", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read http response body: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("received non-204 http response status code (%d), body: %s", resp.StatusCode, string(respBodyBytes))
	}

	return nil
}
