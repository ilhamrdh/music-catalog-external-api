package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

type SpotifySearchResponse struct {
	Tracks SpotifyTracks `json:"tracks"`
}

type SpotifyTracks struct {
	Href     string               `json:"href"`
	Limit    int                  `json:"limit"`
	Next     *string              `json:"next"`
	Offset   int                  `json:"offset"`
	Previous *string              `json:"previous"`
	Total    int                  `json:"total"`
	Items    []SpotifyTrackObject `json:"items"`
}

type (
	SpotifyTrackObject struct {
		ID       string                `json:"id"`
		Album    SpotifyAlbumObject    `json:"album"`
		Artists  []SpotifyArtistObject `json:"artists"`
		Explicit bool                  `json:"explicit"`
		Href     string                `json:"href"`
		Name     string                `json:"name"`
	}

	SpotifyAlbumObject struct {
		AlbumType   string              `json:"album_type"`
		TotalTracks int                 `json:"total_tracks"`
		Images      []SpotifyAlbumImage `json:"images"`
		Name        string              `json:"name"`
	}
	SpotifyAlbumImage struct {
		URL string `json:"url"`
	}

	SpotifyArtistObject struct {
		Href string `json:"href"`
		Name string `json:"name"`
	}
)

func (o *outbound) Search(ctx context.Context, query string, limit, offset int) (*SpotifySearchResponse, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("type", "track")
	params.Set("limit", strconv.Itoa(limit))
	params.Set("offset", strconv.Itoa(offset))

	endpoint := fmt.Sprintf("%s?%s", `https://api.spotify.com/v1/search`, params.Encode())

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Error().Err(err).Msg("error create search request for sptify")
		return nil, err
	}

	accessToken, tokenType, err := o.GetTokenDetail()
	if err != nil {
		log.Error().Err(err).Msg("error get token")
		return nil, err
	}
	bearerToken := fmt.Sprintf("%s %s", tokenType, accessToken)
	req.Header.Set("Authorization", bearerToken)
	res, err := o.Client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("error execute search request for sptify")
		return nil, err
	}
	defer res.Body.Close()

	var response SpotifySearchResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshal response from sptify")
		return nil, err
	}

	return &response, nil
}
