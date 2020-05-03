package eureka

import (
	"context"
	"net/http"
	"net/url"

	"github.com/micro/go-micro/v2/registry"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type contextHttpClient struct{}

var newOAuthClient = func(c clientcredentials.Config) *http.Client {
	return c.Client(oauth2.NoContext)
}

// Enable OAuth 2.0 Client Credentials Grant Flow
func OAuth2ClientCredentials(clientID, clientSecret, tokenURL string) registry.Option {
	return func(o *registry.Options) {
		c := clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     tokenURL,
		}

		o.Context = context.WithValue(o.Context, contextHttpClient{}, newOAuthClient(c))
	}
}

// Enable BasicAuth Client
func BasicAuth(username, password string) registry.Option {
	return func(o *registry.Options) {
		o.Context = context.WithValue(o.Context, contextHttpClient{}, &http.Client{
			Transport: &http.Transport{
				Proxy: func(req *http.Request) (*url.URL, error) {
					req.SetBasicAuth(username, password)
					return nil, nil
				},
			},
		})
	}
}
