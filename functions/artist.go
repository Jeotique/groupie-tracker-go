package functions

import (
	"encoding/json"
	"fmt"
	"project/variables"
)

func GetArtists(id_artist string, token string) (variables.Artist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=genre:pop&type=artist")
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.Artist{}, fmt.Errorf(resErr.Error())
	}

	defer res.Body.Close()
	var resultList variables.Artist
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return variables.Artist{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil
}

func GetbyArtists(id_track string, token string) (variables.TrackInfo, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%v/tracks", id_track)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.TrackInfo{}, fmt.Errorf(resErr.Error())
	}

	defer res.Body.Close()
	var resultList variables.TrackInfo
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return variables.TrackInfo{}, fmt.Errorf(errDecode.Error())
	}
	fmt.Println(resultList)
	return resultList, nil
}

func SearchArtists(q string, token string) (variables.SearchArtists, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%v&type=artist&market=FR", q)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.SearchArtists{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList variables.SearchArtists
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return variables.SearchArtists{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil
}
