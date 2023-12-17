package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"os"
	"unsplash-recap/unsplash"
	"unsplash-recap/upstash"
	"unsplash-recap/utils"
)

type UnsplashRecapEvent struct {
	Body string `json:"body"`
}

func (u *UnsplashRecapEvent) Validate() error {
	if u.Body == "" {
		return fmt.Errorf("body is empty")
	}
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event UnsplashRecapEvent) (*utils.Response, error) {
	// Validate event
	err := event.Validate()
	if err != nil {
		return utils.JSONResponse(400, err.Error(), err.Error()), nil
	}

	accessKey := os.Getenv("UNSPLASH_ACCESS_KEY")
	if accessKey == "" {
		return utils.JSONResponse(500, "unsplash access key is empty", utils.ErrorResponseBody{Message: "unsplash access key is empty"}), fmt.Errorf("unsplash access key is empty")
	}
	redisUrl := os.Getenv("UPSTASH_REDIS_REST_URL")
	if redisUrl == "" {
		return utils.JSONResponse(500, "redis url is empty", utils.ErrorResponseBody{Message: "redis url is empty"}), fmt.Errorf("redis url is empty")
	}
	redisToken := os.Getenv("UPSTASH_REDIS_REST_TOKEN")
	if redisToken == "" {
		return utils.JSONResponse(500, "redis token is empty", utils.ErrorResponseBody{Message: "redis token is empty"}), fmt.Errorf("redis token is empty")
	}

	username, err := utils.GetUsernameFromBody(event.Body)
	if err != nil {
		return utils.JSONResponse(400, err.Error(), utils.ErrorResponseBody{Message: err.Error()}), nil
	}

	// Create Upstash Redis REST client
	client := upstash.NewClient(redisUrl, redisToken)

	// Check if username is cached
	cached, err := client.Get(username)
	if err != nil {
		log.Error(err)
		return utils.JSONResponse(500, err.Error(), utils.ErrorResponseBody{Message: err.Error()}), fmt.Errorf("failed to get cached username: %s", err.Error())
	}
	if cached != nil {
		log.Printf("username: %s is cached", username)

		// Unmarshal cached username
		var recap *unsplash.Recap
		err = json.Unmarshal([]byte(*cached), &recap)
		if err != nil {
			return utils.JSONResponse(500, err.Error(), utils.ErrorResponseBody{Message: err.Error()}), fmt.Errorf("failed to unmarshal cached username: %s", err.Error())
		}

		return utils.JSONResponse(200, "Success", recap), nil
	}

	// Get recap
	var recap *unsplash.Recap
	recap, err = getRecap(username, accessKey)
	if err != nil {
		return utils.JSONResponse(500, err.Error(), utils.ErrorResponseBody{Message: err.Error()}), nil
	}

	if recap == nil {
		return utils.JSONResponse(500, "recap is nil", utils.ErrorResponseBody{Message: "recap is nil"}), nil
	}

	if recap.TotalPhotos == 0 {
		return utils.JSONResponse(500, "user has no photos", utils.ErrorResponseBody{Message: "user has no photos"}), nil
	}

	err = client.Set(username, recap)
	if err != nil {
		log.Error(err)
		return utils.JSONResponse(500, err.Error(), utils.ErrorResponseBody{Message: err.Error()}), fmt.Errorf("failed to set cached username: %s", err.Error())
	}

	return utils.JSONResponse(200, "Success", recap), nil
}

func getRecap(username, accessKey string) (*unsplash.Recap, error) {
	client := unsplash.NewClient(accessKey)

	var err error
	var photos []*unsplash.UserPhoto

	// Get first list of photos
	photos, err = client.Photo().GetUserPhotos(username, 1)
	if err != nil {
		return nil, err
	}

	if len(photos) == 0 {
		return nil, fmt.Errorf("user has no photos")
	}

	if photos[0].CreatedAt.Year() < 2023 {
		return nil, fmt.Errorf("user has no photos in 2023")
	}

	if len(photos) == 30 {
		// Get users photo if last photo is still in 2023 get next page
		for i := 2; utils.CheckLastPhotoYear(photos, 2023); i++ {
			newPagePhotos, err := client.Photo().GetUserPhotos(username, i)
			if err != nil {
				return nil, err
			}
			if len(newPagePhotos) < 30 {
				photos = append(photos, newPagePhotos...)
				break
			}
			photos = append(photos, newPagePhotos...)
		}
	}

	// Filter photos by year
	photos = utils.FilterByYear(photos, 2023)

	// Get recap
	return utils.GetRecapFromPhotos(photos), nil
}
