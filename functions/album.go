package functions

import (
	"encoding/json"
	"fmt"
	"project/variables"
)

func GetRecentAlbums(token string) ([]variables.Album, error) {
	id_artist := "3IW7ScrzXmPvZhB27hmfgy"
	url := "https://api.spotify.com/v1/artists/" + id_artist + "/albums?include_groups=single&limit=50"
	resp, err := requestSpotifyAPI(url, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	var searchResults variables.SearchResults
	if err := json.NewDecoder(resp.Body).Decode(&searchResults); err != nil {
		return nil, err
	}

	return searchResults.Albums, nil
}

func GetPopularAlbums(id_artist string, token string) (variables.Albums, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/artists/%v/albums", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.Albums{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList variables.Albums
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return variables.Albums{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil

}

func GetAlbums(id_jul string, token string) (variables.Album, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/artists/%v/albums", id_jul)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.Album{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList variables.Album
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return variables.Album{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil

}

func GetAlbum(id_album string, token string) (variables.Album, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%v", id_album)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return variables.Album{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList variables.Album
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return variables.Album{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil
}

func GetReleased(token string) (variables.Album, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/browse/new-releases")
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		println("error")
		return variables.Album{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList variables.Album
	println(json.NewDecoder(res.Body))
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		fmt.Println("Decode error")
		return variables.Album{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil
}
