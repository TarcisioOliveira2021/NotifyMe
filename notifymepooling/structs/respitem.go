package structs

type Item struct {
	AlbumType            AlbumGroup           `json:"album_type"`
	TotalTracks          int64                `json:"total_tracks"`
	IsPlayable           bool                 `json:"is_playable"`
	ExternalUrl 		 ExternalUrls         `json:"external_urls"`
	ID                   string               `json:"id"`
	Images               []Image              `json:"images"`
	Name                 string               `json:"name"`
	ReleaseDate          string               `json:"release_date"`
	Type                 AlbumGroup           `json:"type"`
	URI                  string               `json:"uri"`
	Artists              []ArtistElement      `json:"artists"`
	AlbumGroup           AlbumGroup           `json:"album_group"`
}