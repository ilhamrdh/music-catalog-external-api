package spotify

import (
	"context"

	extSpotify "github.com/ilhamrdh/music-catalog-external-api/external/spotify"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	tracks, err := s.spotifyOutboung.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}

	return modelToResponse(tracks), nil
}

func modelToResponse(data *extSpotify.SpotifySearchResponse) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)
	for _, item := range data.Tracks.Items {
		artistsName := make([]string, len(item.Artists))
		for index, artist := range item.Artists {
			artistsName[index] = artist.Name
		}
		imageUrls := make([]string, len(item.Album.Images))
		for index, image := range item.Album.Images {
			imageUrls[index] = image.URL
		}

		items = append(items, spotify.SpotifyTrackObject{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL:   imageUrls,
			AlbumName:        item.Album.Name,
			ArtistsName:      artistsName,
			Explicit:         item.Explicit,
			ID:               item.ID,
			Name:             item.Name,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Items:  items,
		Total:  data.Tracks.Total,
	}
}
