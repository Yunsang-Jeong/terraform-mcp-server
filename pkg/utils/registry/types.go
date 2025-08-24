package registry

import "time"

type RegistryV1Provider struct {
	ID          string                  `json:"id"`
	Owner       string                  `json:"owner"`
	Namespace   string                  `json:"namespace"`
	Name        string                  `json:"name"`
	Alias       string                  `json:"alias"`
	Version     string                  `json:"version"`
	Tag         string                  `json:"tag"`
	Description string                  `json:"description"`
	Source      string                  `json:"source"`
	PublishedAt time.Time               `json:"published_at"`
	Downloads   int64                   `json:"downloads"`
	Tier        string                  `json:"tier"`
	LogoURL     string                  `json:"logo_url"`
	Versions    []string                `json:"versions"`
	Docs        []RegistryV1ProviderDoc `json:"docs"`
}

type RegistryV1ProviderDoc struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Path        string `json:"path"`
	Slug        string `json:"slug"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Language    string `json:"language"`
}

type RegistryV2Provider struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Alias         string `json:"alias"`
			Description   string `json:"description"`
			Downloads     int64  `json:"downloads"`
			Featured      bool   `json:"featured"`
			FullName      string `json:"full-name"`
			LogoURL       string `json:"logo-url"`
			Name          string `json:"name"`
			Namespace     string `json:"namespace"`
			OwnerName     string `json:"owner-name"`
			RobotsNoindex bool   `json:"robots-noindex"`
			Source        string `json:"source"`
			Tier          string `json:"tier"`
			Unlisted      bool   `json:"unlisted"`
			Warning       string `json:"warning"`
		} `json:"attributes"`
		Relationships struct {
			ProviderVersions struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"provider-versions"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Description string    `json:"description"`
			Downloads   int       `json:"downloads"`
			PublishedAt time.Time `json:"published-at"`
			Tag         string    `json:"tag"`
			Version     string    `json:"version"`
		} `json:"attributes"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"included"`
}

type RegistryV2ProviderDocs struct {
	Data []RegistryV2ProviderDoc `json:"data"`
}

type RegistryV2ProviderDoc struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Category    string      `json:"category"`
		Language    string      `json:"language"`
		Path        string      `json:"path"`
		Slug        string      `json:"slug"`
		Subcategory interface{} `json:"subcategory"`
		Title       string      `json:"title"`
		Truncated   bool        `json:"truncated"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type RegistryV2ProviderDocsId struct {
	Data RegistryV2ProviderDocsIdData `json:"data"`
}

type RegistryV2ProviderDocsIdData struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Category    string      `json:"category"`
		Content     string      `json:"content"`
		Language    string      `json:"language"`
		Path        string      `json:"path"`
		Slug        string      `json:"slug"`
		Subcategory interface{} `json:"subcategory"`
		Title       string      `json:"title"`
		Truncated   bool        `json:"truncated"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
