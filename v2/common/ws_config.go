package common

type WsConfig struct {
	IsTestnet bool
	Proxy     *string
}

type WsOption func(*WsConfig)

func ResolveWsConfig(opts ...WsOption) *WsConfig {
	cfg := new(WsConfig)
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func WithTestnet(testnet bool) WsOption {
	return func(cfg *WsConfig) {
		cfg.IsTestnet = testnet
	}
}

func WithProxy(proxy string) WsOption {
	return func(cfg *WsConfig) {
		if proxy == "" {
			cfg.Proxy = nil
			return
		}
		cfg.Proxy = &proxy
	}
}
