package functions

import (
	"encoding/json"
	"fmt"
	"project/variables"
)

func GetPhonkPlaylists(id_artist string, token string) (variables.Playlist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%v", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.Playlist{}, fmt.Errorf("failed to fetch playlist: %v", resErr) // Remarquez Playlist{} au lieu de nil
	}
	defer res.Body.Close()

	var response variables.Playlist

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return variables.Playlist{}, fmt.Errorf("failed to decode playlist response: %v", err) // Remarquez Playlist{} au lieu de nil
	}
	return response, nil
}

func GetRapPlaylists(id_artist string, token string) (variables.Playlist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%v", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.Playlist{}, fmt.Errorf("failed to fetch playlist: %v", resErr) // Remarquez Playlist{} au lieu de nil
	}
	defer res.Body.Close()

	var response variables.Playlist

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return variables.Playlist{}, fmt.Errorf("failed to decode playlist response: %v", err) // Remarquez Playlist{} au lieu de nil
	}
	return response, nil
}

func GetPopPlaylists(id_artist string, token string) (variables.Playlist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%v", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.Playlist{}, fmt.Errorf("failed to fetch playlist: %v", resErr) // Remarquez Playlist{} au lieu de nil
	}
	defer res.Body.Close()

	var response variables.Playlist

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return variables.Playlist{}, fmt.Errorf("failed to decode playlist response: %v", err) // Remarquez Playlist{} au lieu de nil
	}
	return response, nil
}
