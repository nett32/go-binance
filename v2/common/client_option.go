package common

import (
	"net/http"
	"net/url"
)

type ProxyFunc func(*http.Request) (*url.URL, error)

type ClientConfig struct {
	UseTestnet bool
	httpClient *http.Client
	ProxyFunc  ProxyFunc
}

func (cfg *ClientConfig) HTTPClient() *http.Client {
	if cfg.httpClient != nil {
		return cfg.httpClient
	}
	if cfg.ProxyFunc != nil {
		return &http.Client{Transport: &http.Transport{Proxy: cfg.ProxyFunc}}
	}
	return http.DefaultClient
}

func ParseClientConfig(opts ...ClientOptionFunc) ClientConfig {
	cfg := ClientConfig{UseTestnet: false}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

type ClientOptionFunc func(*ClientConfig)

func UseTestnet(useTestnet bool) ClientOptionFunc {
	return func(cfg *ClientConfig) {
		cfg.UseTestnet = useTestnet
	}
}

func WithHttpClient(client *http.Client) ClientOptionFunc {
	return func(cfg *ClientConfig) {
		cfg.httpClient = client
	}
}

func WithProxyFunc(proxyFunc ProxyFunc) ClientOptionFunc {
	return func(cfg *ClientConfig) {
		cfg.ProxyFunc = proxyFunc
	}
}
