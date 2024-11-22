package spotify

import (
	"context"
	"reflect"
	"testing"

	extSpotify "github.com/ilhamrdh/music-catalog-external-api/external/spotify"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/spotify"
	"go.uber.org/mock/gomock"
)

func Test_service_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSpotifyOutbound := NewMockspotifyOutboung(mockCtrl)
	next := "https://api.spotify.com/v1/search?offset=10&limit=10&query=bernadya&type=track&market=ID&locale=id-ID,id;q%3D0.9,en-US;q%3D0.8,en;q%3D0.7,de;q%3D0.6"
	type args struct {
		query     string
		pageSize  int
		pageIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.SearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:     "bernadya",
				pageSize:  10,
				pageIndex: 1,
			},
			want: &spotify.SearchResponse{
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
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(
					&extSpotify.SpotifySearchResponse{
						Tracks: extSpotify.SpotifyTracks{
							Href:   "https://api.spotify.com/v1/search?offset=0&limit=10&query=bernadya&type=track&market=ID&locale=id-ID,id;q%3D0.9,en-US;q%3D0.8,en;q%3D0.7,de;q%3D0.6",
							Limit:  10,
							Next:   &next,
							Offset: 0,
							Total:  818,
							Items: []extSpotify.SpotifyTrackObject{
								{
									Album: extSpotify.SpotifyAlbumObject{
										AlbumType:   "single",
										TotalTracks: 5,
										Images: []extSpotify.SpotifyAlbumImage{
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
									Artists: []extSpotify.SpotifyArtistObject{
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
									Album: extSpotify.SpotifyAlbumObject{
										AlbumType:   "album",
										TotalTracks: 8,
										Images: []extSpotify.SpotifyAlbumImage{
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
									Artists: []extSpotify.SpotifyArtistObject{
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
					}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				spotifyOutboung: mockSpotifyOutbound,
			}
			got, err := s.Search(context.Background(), tt.args.query, tt.args.pageSize, tt.args.pageIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
