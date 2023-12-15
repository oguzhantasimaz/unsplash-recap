package unsplash

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type PhotoService struct {
	client *Client
}

func (c *Client) Photo() (photo *PhotoService) {
	photo = &PhotoService{client: c}
	return
}

func (s *PhotoService) GetUserPhotos(username string, page int) (photos []*UserPhoto, err error) {
	var queryParameters = map[string]string{
		"stats":      "true",
		"resolution": "days",
		"order_by":   "latest",
		"quantity":   "1",
		"per_page":   "30",
		"page":       fmt.Sprintf("%d", page),
	}

	var userPhotos []*UserPhoto

	url := fmt.Sprintf("%s/users/%s/photos", s.client.BaseURL, username)
	body, err := s.client.Get(url, queryParameters)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(body, &userPhotos)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	photos = append(photos, userPhotos...)
	return
}

func (s *PhotoService) GetPhoto(id string) (photo *Photo, err error) {
	url := fmt.Sprintf("%s/photos/%s", s.client.BaseURL, id)
	body, err := s.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &photo)
	if err != nil {
		return nil, err
	}

	return
}
