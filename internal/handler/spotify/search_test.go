package spotify

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/spotify"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := NewMockservice(mockCtrl)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       spotify.SearchResponse
		wantErr            bool
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			wantErr:            false,
			expectedBody: spotify.SearchResponse{
				Limit:  10,
				Offset: 0,
				Items: []spotify.SpotifyTrackObject{
					{
						ID:       "7zOVh5fGpEwSbZd0g9z80B",
						Name:     "Satu Bulan",
						Explicit: false,
						ArtistsName: []string{
							"Bernadya",
						},
						AlbumType:        "single",
						AlbumTotalTracks: 5,
						AlbumImagesURL: []string{
							"https://i.scdn.co/image/ab67616d0000b273cb1e5f7d0942bf9196c1e711",
							"https://i.scdn.co/image/ab67616d00001e02cb1e5f7d0942bf9196c1e711",
							"https://i.scdn.co/image/ab67616d00004851cb1e5f7d0942bf9196c1e711",
						},
						AlbumName: "Terlintas",
					},
					{
						ID:       "2gcMYiZzzmzoF8PPAfL3IO",
						Name:     "Untungnya, Hidup Harus Tetap Berjalan",
						Explicit: false,
						ArtistsName: []string{
							"Bernadya",
						},
						AlbumType:        "album",
						AlbumTotalTracks: 8,
						AlbumImagesURL: []string{
							"https://i.scdn.co/image/ab67616d0000b27327693aaf059002bba3a7655a",
							"https://i.scdn.co/image/ab67616d00001e0227693aaf059002bba3a7655a",
							"https://i.scdn.co/image/ab67616d0000485127693aaf059002bba3a7655a",
						},
						AlbumName: "Sialnya, Hidup Harus Tetap Berjalan",
					},
				},
				Total: 818,
			},
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "bernadya", 10, 1).Return(&spotify.SearchResponse{
					Limit:  10,
					Offset: 0,
					Items: []spotify.SpotifyTrackObject{
						{
							ID:       "7zOVh5fGpEwSbZd0g9z80B",
							Name:     "Satu Bulan",
							Explicit: false,
							ArtistsName: []string{
								"Bernadya",
							},
							AlbumType:        "single",
							AlbumTotalTracks: 5,
							AlbumImagesURL: []string{
								"https://i.scdn.co/image/ab67616d0000b273cb1e5f7d0942bf9196c1e711",
								"https://i.scdn.co/image/ab67616d00001e02cb1e5f7d0942bf9196c1e711",
								"https://i.scdn.co/image/ab67616d00004851cb1e5f7d0942bf9196c1e711",
							},
							AlbumName: "Terlintas",
						},
						{
							ID:       "2gcMYiZzzmzoF8PPAfL3IO",
							Name:     "Untungnya, Hidup Harus Tetap Berjalan",
							Explicit: false,
							ArtistsName: []string{
								"Bernadya",
							},
							AlbumType:        "album",
							AlbumTotalTracks: 8,
							AlbumImagesURL: []string{
								"https://i.scdn.co/image/ab67616d0000b27327693aaf059002bba3a7655a",
								"https://i.scdn.co/image/ab67616d00001e0227693aaf059002bba3a7655a",
								"https://i.scdn.co/image/ab67616d0000485127693aaf059002bba3a7655a",
							},
							AlbumName: "Sialnya, Hidup Harus Tetap Berjalan",
						},
					},
					Total: 818,
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			h := &Handler{
				engine:  gin.New(),
				service: mockSvc,
			}
			h.SpotifyRoute()
			w := httptest.NewRecorder()

			endpoint := `/tracks/search?query=bernadya&pageSize=10&pageIndex=1`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			h.engine.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.SearchResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
