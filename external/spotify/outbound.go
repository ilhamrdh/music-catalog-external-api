package spotify

import (
	"time"

	"github.com/ilhamrdh/music-catalog-external-api/internal/configs"
	"github.com/ilhamrdh/music-catalog-external-api/pkg/httpclient"
)

type outbound struct {
	Cfg         *configs.Config
	Client      httpclient.HTTPClient
	AccessToken string
	TokenType   string
	ExpiredAt   time.Time
}

func NewSpotifyOutbound(cfg *configs.Config, client httpclient.HTTPClient) *outbound {
	return &outbound{
		Cfg:    cfg,
		Client: client,
	}
}
