package api

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func GetProfile(username string) (*UserProfileResponse, error) {
	var url = fmt.Sprintf("https://api.jikan.moe/v3/user/%s/", username)
	_, body, errs := gorequest.New().Post(url).End()
	if errs != nil {
		return nil, fmt.Errorf(
			"Failed to retrieve MAL profile for user: %s", username)
	}

	var userProfile = &UserProfileResponse{}
	json.Unmarshal([]byte(body), userProfile)
	return userProfile, nil
}
