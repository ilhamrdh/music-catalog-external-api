package spotify

import (
	"context"

	"github.com/ilhamrdh/music-catalog-external-api/external/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=spotify
type spotifyOutboung interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type service struct {
	spotifyOutboung spotifyOutboung
}

func NewService(spotifyOutboung spotifyOutboung) *service {
	return &service{spotifyOutboung: spotifyOutboung}
}
