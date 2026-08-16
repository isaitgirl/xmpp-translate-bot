package lookup

import (
	"context"
	"net/http"
	"time"
)

type WeatherConfig struct {
	// Host é o domínio do serviço de clima, sem esquema. Ex.: "api.open-meteo.com".
	Host string
	// UserAgent identifica o bot. Obrigatório: alguns serviços respondem 403 pra
	// User-Agent genérico. Ex.:
	// "xmpp-translate-bot/1.0 (https://exemplo.org)".
	UserAgent  string
	Timeout    time.Duration
	MaxRetries int
	// Transport é opcional: nil usa o http.DefaultTransport. Existe pra
	// teste apontar o client pra um httptest.Server.
	Transport http.RoundTripper
}

// Weather consulta informações de clima via API.
type Weather struct {
	cfg  WeatherConfig
	http *httpJSON
}

// NewWeather cria a fonte. Não faz I/O.
func NewWeather(cfg WeatherConfig) *Weather {
	return &Weather{
		cfg:  cfg,
		http: newHTTPJSON(cfg.UserAgent, cfg.Timeout, cfg.MaxRetries, cfg.Transport),
	}
}

func (w *Weather) Look(ctx context.Context, query string) (Result, error) {

	// Call the weather API using httpjson utility
	// Parse the response and return a Result
	// This is a placeholder implementation; actual API call and parsing logic should be implemented here.

	return Result{}, nil
}
