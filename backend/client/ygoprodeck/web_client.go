package ygoprodeck

const endpointCardInfo = "https://db.ygoprodeck.com/api/v7"

type JsonWebClient interface {
	// GetJsonFromTarget retrieves a json object of the given type from the target url.
	GetJsonFromTarget(targetUrl string, data any) error
}

// YgoProDeckClient is responsible to extract all necessary card information from the ygoprodeck website.
type YgoProDeckClient struct {
	BaseUrl string        `json:"base_url,omitempty"`
	Client  JsonWebClient `json:"Client" json:"client,omitempty"`
}

// NewYgoProDeckClient creates a new instance of the YgoProDeckClient.
func NewYgoProDeckClient() *YgoProDeckClient {
	rateLimitedClient := NewDefaultRateLimitedClient()

	return &YgoProDeckClient{
		BaseUrl: endpointCardInfo,
		Client:  rateLimitedClient,
	}
}
