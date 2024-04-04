package functions

import (
	"encoding/json"
	"fmt"
	"project/variables"
)

func GetMusics(id_music string, token string) (variables.TrackInfo, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%v/tracks", id_music)
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
	return resultList, nil
}
