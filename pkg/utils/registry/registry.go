package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Yunsang-Jeong/terraform-mcp-server/pkg/utils"
)

const (
	HTTP_TIMEOUT = 10 // seconds
)

func GetSomethingFromPublicRegistry(path string, query map[string]string) ([]byte, error) {
	u, err := url.Parse("https://registry.terraform.io")
	if err != nil {
		return nil, fmt.Errorf("error parsing terraform registry URL: %w", err)
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	u.Path = path

	q := u.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "terraform-mcp/0.1 (+public-registry)")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: HTTP_TIMEOUT * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("registry error: status=%d body=%s", resp.StatusCode, string(body))
	}

	return body, nil
}

func GetProvider(namespace, name string) (RegistryV1Provider, error) {
	resp := RegistryV1Provider{}

	path := fmt.Sprintf("/v1/providers/%s/%s", url.PathEscape(namespace), url.PathEscape(name))
	query := map[string]string{}

	data, err := GetSomethingFromPublicRegistry(path, query)
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}

func GetProviderLatestVersion(namespace, name string) (string, error) {
	path := fmt.Sprintf("/v1/providers/%s/%s", url.PathEscape(namespace), url.PathEscape(name))
	query := map[string]string{}

	data, err := GetSomethingFromPublicRegistry(path, query)
	if err != nil {
		return "", err
	}

	var resp RegistryV1Provider
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}

	return resp.Version, nil
}

func GetProviderVersionId(namespace, name, version string) (string, error) {
	path := fmt.Sprintf("/v2/providers/%s/%s", url.PathEscape(namespace), url.PathEscape(name))
	query := map[string]string{
		"include": "provider-versions",
	}

	data, err := GetSomethingFromPublicRegistry(path, query)
	if err != nil {
		return "", err
	}

	var resp RegistryV2Provider
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}

	for i := len(resp.Included) - 1; i >= 0; i-- {
		if resp.Included[i].Attributes.Version == version {
			return resp.Included[i].ID, nil
		}
	}

	return "", fmt.Errorf("fail to find provider verions id: %s/%s %s", namespace, name, version)
}

func GetProviderDocsId(versionId, category, slug string) (string, error) {
	if !utils.IsInList(category, []string{"overview", "resources", "data-sources"}) {
		return "", fmt.Errorf("invalid category: %s", category)
	}

	path := "/v2/provider-docs"
	query := map[string]string{
		"filter[provider-version]": versionId,
		"filter[category]":         category,
		"filter[slug]":             slug,
		"filter[language]":         "hcl",
	}

	data, err := GetSomethingFromPublicRegistry(path, query)
	if err != nil {
		return "", err
	}

	var resp RegistryV2ProviderDocs
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}

	if len(resp.Data) != 1 {
		return "", fmt.Errorf("response must be one: %v", resp.Data)
	}

	return resp.Data[0].ID, nil
}

func GetProviderDocsContent(docsId string) (string, error) {
	path := fmt.Sprintf("/v2/provider-docs/%s", docsId)
	query := map[string]string{}

	data, err := GetSomethingFromPublicRegistry(path, query)
	if err != nil {
		return "", err
	}

	var resp RegistryV2ProviderDocsId
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}

	return resp.Data.Attributes.Content, nil
}
