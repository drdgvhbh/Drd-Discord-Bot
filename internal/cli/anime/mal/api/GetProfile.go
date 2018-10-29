package api

import (
	"drdgvhbh/discordbot/internal/cli/anime/mal/api/response"
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func GetProfile(username string) (*response.UserProfileResponse, error) {
	var url = fmt.Sprintf("https://api.jikan.moe/v3/user/%s/", username)
	_, body, errs := gorequest.New().Post(url).End()
	if errs != nil {
		return nil, fmt.Errorf(
			"Failed to retrieve MAL profile for user: %s", username)
	}

	var userProfile = &response.UserProfileResponse{}
	json.Unmarshal([]byte(body), userProfile)
	return userProfile, nil
}
