package facebook

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func currentId(accessToken string) (string, error) {
	res, err := fb.Get("/me", fb.Params{
		"fields":       "id",
		"access_token": accessToken,
	})
	if err != nil {
		return "", err
	}

	return res["id"].(string), nil
}

type LiveVideo struct {
	Id              string
	StreamUrl       string
	SecureStreamUrl string
}

func CreateLiveVideo(accessToken string) (vid LiveVideo, err error) {
	id, err := currentId(accessToken)
	if err != nil {
		return vid, err
	}

	res, err := fb.Post(fmt.Sprintf("/%s/live_videos", id), fb.Params{
		"access_token": accessToken,
	})
	if err != nil {
		return vid, err
	}

	vid.Id = res["id"].(string)
	vid.StreamUrl = res["stream_url"].(string)
	vid.SecureStreamUrl = res["secure_stream_url"].(string)

	return vid, nil
}
