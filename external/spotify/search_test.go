package spotify

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/ilhamrdh/music-catalog-external-api/internal/configs"
	"github.com/ilhamrdh/music-catalog-external-api/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_outbound_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHTTPClient := httpclient.NewMockHTTPClient(mockCtrl)
	next := "https://api.spotify.com/v1/search?offset=10&limit=10&query=bernadya&type=track&market=ID&locale=id-ID,id;q%3D0.9,en-US;q%3D0.8,en;q%3D0.7,de;q%3D0.6"
	type args struct {
		query  string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    *SpotifySearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:  "bernadya",
				limit:  10,
				offset: 0,
			},
			want: &SpotifySearchResponse{
				Tracks: SpotifyTracks{
					Href:   "https://api.spotify.com/v1/search?offset=0&limit=10&query=bernadya&type=track&market=ID&locale=id-ID,id;q%3D0.9,en-US;q%3D0.8,en;q%3D0.7,de;q%3D0.6",
					Limit:  10,
					Next:   &next,
					Offset: 0,
					Total:  818,
					Items: []SpotifyTrackObject{
						{
							Album: SpotifyAlbumObject{
								AlbumType:   "single",
								TotalTracks: 5,
								Images: []SpotifyAlbumImage{
									{
										URL: "https://i.scdn.co/image/ab67616d0000b273cb1e5f7d0942bf9196c1e711",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d00001e02cb1e5f7d0942bf9196c1e711",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d00004851cb1e5f7d0942bf9196c1e711",
									},
								},
								Name: "Terlintas",
							},
							Artists: []SpotifyArtistObject{
								{
									Href: "https://api.spotify.com/v1/artists/47z98pKd71yIbgXwe9LPVC",
									Name: "Bernadya",
								},
							},
							ID:       "7zOVh5fGpEwSbZd0g9z80B",
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/7zOVh5fGpEwSbZd0g9z80B",
							Name:     "Satu Bulan",
						},
						{
							Album: SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 8,
								Images: []SpotifyAlbumImage{
									{
										URL: "https://i.scdn.co/image/ab67616d0000b27327693aaf059002bba3a7655a",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d00001e0227693aaf059002bba3a7655a",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d0000485127693aaf059002bba3a7655a",
									},
								},
								Name: "Sialnya, Hidup Harus Tetap Berjalan",
							},
							Artists: []SpotifyArtistObject{
								{
									Href: "https://api.spotify.com/v1/artists/47z98pKd71yIbgXwe9LPVC",
									Name: "Bernadya",
								},
							},
							ID:       "2gcMYiZzzmzoF8PPAfL3IO",
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/2gcMYiZzzmzoF8PPAfL3IO",
							Name:     "Untungnya, Hidup Harus Tetap Berjalan",
						},
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				endpoint := fmt.Sprintf("%s?%s", `https://api.spotify.com/v1/search`, params.Encode())

				req, err := http.NewRequest(http.MethodGet, endpoint, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
				}, nil)
			},
		},
		{
			name: "Fail",
			args: args{
				query:  "bernadya",
				limit:  10,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				endpoint := fmt.Sprintf("%s?%s", `https://api.spotify.com/v1/search`, params.Encode())

				req, err := http.NewRequest(http.MethodGet, endpoint, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(bytes.NewBufferString(`Internal Server Error`)),
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			o := &outbound{
				Cfg:         &configs.Config{},
				Client:      mockHTTPClient,
				AccessToken: "accessToken",
				TokenType:   "Bearer",
				ExpiredAt:   time.Now().Add(1 * time.Hour),
			}

			got, err := o.Search(context.Background(), tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("outbound.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
