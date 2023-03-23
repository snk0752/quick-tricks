package netclient

import (
	"net/http"
	"net/url"
)

func NewHTTPClient(proxyString string) (http.Client, error) {
	client := http.Client{
		Transport: http.DefaultTransport,
	}

	if proxyString != "" {
		// Parse the proxy URL.
		proxyURL, err := url.Parse(proxyString)
		if err != nil {
			return client, err
		}

		// Create new transport with the proxy.
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.Proxy = http.ProxyURL(proxyURL)

		// Set the new transport as the default
		client.Transport = transport
	}

	return client, nil
}
