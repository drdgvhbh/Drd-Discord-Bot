package mal

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/parnurzeal/gorequest"
)

func GetProfile(username string) *UserProfile {
	var url = fmt.Sprintf("https://api.jikan.moe/v3/user/%s/", username)
	_, body, errs := gorequest.New().Post(url).End()
	if errs != nil {
		log.Fatalln(errs)
	}

	var userProfile = &UserProfile{}
	json.Unmarshal([]byte(body), userProfile)
	return userProfile
}
