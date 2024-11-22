package spotify

type SearchResponse struct {
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Items  []SpotifyTrackObject `json:"items"`
	Total  int                  `json:"total"`
}
type (
	SpotifyTrackObject struct {
		ID               string   `json:"id"`
		Name             string   `json:"name"`
		Explicit         bool     `json:"explicit"`
		ArtistsName      []string `json:"artists_name"`
		AlbumType        string   `json:"album_type"`
		AlbumTotalTracks int      `json:"album_total_tracks"`
		AlbumImagesURL   []string `json:"album_images_url"`
		AlbumName        string   `json:"album_name"`
	}
)
