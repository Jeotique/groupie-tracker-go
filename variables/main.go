package variables

import (
	"text/template"
	"time"
)

var (
	ClientID     = "b46d215af0cc40d28a63b64d810c2e2c"
	ClientSecret = "113870a84d914a2393a8e31719eb241c"
	AccessToken  string
	Tmpl         *template.Template
)

type ExternalURLsSearch struct {
	Spotify string `json:"spotify"`
}

// FollowersSearch représente les abonnés.
type FollowersSearch struct {
	Href  interface{} `json:"href"`
	Total int         `json:"total"`
}

// ImageSearch représente une image.
type ImageSearch struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Playlist struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ExternalUrls  struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		Height interface{} `json:"height"`
		URL    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	Name  string `json:"name"`
	Owner struct {
		DisplayName  string `json:"display_name"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"owner"`
	PrimaryColor string `json:"primary_color"`
	Public       bool   `json:"public"`
	SnapshotID   string `json:"snapshot_id"`
	Tracks       struct {
		Href  string `json:"href"`
		Items []struct {
			AddedAt time.Time `json:"added_at"`
			AddedBy struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"added_by"`
			IsLocal      bool        `json:"is_local"`
			PrimaryColor interface{} `json:"primary_color"`
			Track        struct {
				Album struct {
					AlbumType string `json:"album_type"`
					Artists   []struct {
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href string `json:"href"`
						ID   string `json:"id"`
						Name string `json:"name"`
						Type string `json:"type"`
						URI  string `json:"uri"`
					} `json:"artists"`
					AvailableMarkets []string `json:"available_markets"`
					ExternalUrls     struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href   string `json:"href"`
					ID     string `json:"id"`
					Images []struct {
						Height int    `json:"height"`
						URL    string `json:"url"`
						Width  int    `json:"width"`
					} `json:"images"`
					Name                 string `json:"name"`
					ReleaseDate          string `json:"release_date"`
					ReleaseDatePrecision string `json:"release_date_precision"`
					TotalTracks          int    `json:"total_tracks"`
					Type                 string `json:"type"`
					URI                  string `json:"uri"`
				} `json:"album"`
				Artists []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets []string `json:"available_markets"`
				DiscNumber       int      `json:"disc_number"`
				DurationMs       int      `json:"duration_ms"`
				Episode          bool     `json:"episode"`
				Explicit         bool     `json:"explicit"`
				ExternalIds      struct {
					Isrc string `json:"isrc"`
				} `json:"external_ids"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href        string `json:"href"`
				ID          string `json:"id"`
				IsLocal     bool   `json:"is_local"`
				Name        string `json:"name"`
				Popularity  int    `json:"popularity"`
				PreviewURL  string `json:"preview_url"`
				Track       bool   `json:"track"`
				TrackNumber int    `json:"track_number"`
				Type        string `json:"type"`
				URI         string `json:"uri"`
			} `json:"track"`
			VideoThumbnail struct {
				URL interface{} `json:"url"`
			} `json:"video_thumbnail"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

// ArtistSearch représente un artiste.
type ArtistSearch struct {
	ExternalURLs ExternalURLsSearch `json:"external_urls"`
	Followers    FollowersSearch    `json:"followers"`
	Genres       []string           `json:"genres"`
	Href         string             `json:"href"`
	ID           string             `json:"id"`
	Images       []ImageSearch      `json:"images"`
	Name         string             `json:"name"`
	Popularity   int                `json:"popularity"`
	Type         string             `json:"type"`
	URI          string             `json:"uri"`
}

// ArtistsResponseSearch représente la réponse de l'API Spotify pour les artistes.
type ArtistsResponseSearch struct {
	Artists struct {
		Href     string         `json:"href"`
		Limit    int            `json:"limit"`
		Next     string         `json:"next"`
		Offset   int            `json:"offset"`
		Previous string         `json:"previous"`
		Total    int            `json:"total"`
		Items    []ArtistSearch `json:"items"`
	} `json:"artists"`
}

type Image struct {
	URL string `json:"url"`
}

type TrackInfo struct {
	Items []struct {
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"artists"`
		DiscNumber   int  `json:"disc_number"`
		DurationMs   int  `json:"duration_ms"`
		Explicit     bool `json:"explicit"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"items"`
}

type SearchResults struct {
	Albums []Album `json:"items"`
}

type SearchArtists struct {
	Artists struct {
		Items []struct {
			External_urls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`

			Followers struct {
				Total int `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Id     string   `json:"id"`
			Image  []struct {
				URL string `json:"url"`
			} `json:"images"`
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"items"`
	} `json:"artists"`
}

type Album struct {
	Items []struct {
		Artist []struct {
			Name          string `json:"name"`
			Id            string `json:"id"`
			External_urls struct {
				Spotify string `json:" spotify "`
			} `json:"external_urls"`
		} `json:"artists"`
		External_urls struct {
			Spotify string `json:" spotify "`
		} `json:"external_urls"`
		Id     string `json:"id"`
		Images []struct {
			Url string `json:"url"`
		} `json:"images"`
		Name         string `json:"name"`
		Release_date string `json:"release_date"`
		Total_tracks int    `json:"total_tracks"`
	} `json:"items"`
}

type Albums struct {
	AlbumType string `json:"album_type"`
	Artists   []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	Copyrights       []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"copyrights"`
	ExternalIds struct {
		Upc string `json:"upc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Genres []interface{} `json:"genres"`
	Href   string        `json:"href"`
	ID     string        `json:"id"`
	Images []struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"images"`
	Label                string `json:"label"`
	Name                 string `json:"name"`
	Popularity           int    `json:"popularity"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks          int    `json:"total_tracks"`
	Tracks               struct {
		Href  string `json:"href"`
		Items []struct {
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			DiscNumber       int      `json:"disc_number"`
			DurationMs       int      `json:"duration_ms"`
			Explicit         bool     `json:"explicit"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			IsLocal     bool   `json:"is_local"`
			Name        string `json:"name"`
			PreviewURL  string `json:"preview_url"`
			TrackNumber int    `json:"track_number"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type PopularArtists struct {
	Artists []Artist `json:"artists"`
}

type Artist struct {
	Followers struct {
		Total int `json:"total"`
	} `json:"followers"`
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
}

type Favorites struct {
	IDs []string `json:"ids"`
}