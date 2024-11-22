package spotify

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (o *outbound) GetTokenDetail() (string, string, error) {
	if o.AccessToken == "" || time.Now().After(o.ExpiredAt) {
		err := o.generateToken()
		if err != nil {
			return "", "", err
		}
	}
	return o.AccessToken, o.TokenType, nil
}

func (o *outbound) generateToken() error {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", o.Cfg.SpotifyConfig.ClientID)
	formData.Set("client_secret", o.Cfg.SpotifyConfig.ClientSecret)

	encodedURL := formData.Encode()

	req, err := http.NewRequest(http.MethodPost, `https://accounts.spotify.com/api/token`, strings.NewReader(encodedURL))
	if err != nil {
		log.Error().Err(err).Msg("error create request for sptify")
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := o.Client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("error execute request for sptify")
		return err
	}
	defer res.Body.Close()

	var response SpotifyTokenResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshal response from sptify")
		return err
	}

	o.AccessToken = response.AccessToken
	o.TokenType = response.TokenType
	o.ExpiredAt = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)

	return nil
}
