package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strings"
	"unsplash-recap/unsplash"
)

func FilterByYear(photos []*unsplash.UserPhoto, year int) []*unsplash.UserPhoto {
	var filteredPhotos []*unsplash.UserPhoto
	for _, photo := range photos {
		if photo.CreatedAt.Year() == year {
			filteredPhotos = append(filteredPhotos, photo)
		}
	}
	return filteredPhotos
}

func CheckLastPhotoYear(photos []*unsplash.UserPhoto, year int) bool {
	return photos[len(photos)-1].CreatedAt.Year() == year
}

func SortByLikes(photos []*unsplash.UserPhoto) []*unsplash.UserPhoto {
	sortedPhotos := make([]*unsplash.UserPhoto, len(photos))
	copy(sortedPhotos, photos)
	for i := 0; i < len(sortedPhotos); i++ {
		for j := 0; j < len(sortedPhotos)-1; j++ {
			if sortedPhotos[j].Likes < sortedPhotos[j+1].Likes {
				sortedPhotos[j], sortedPhotos[j+1] = sortedPhotos[j+1], sortedPhotos[j]
			}
		}
	}
	return sortedPhotos
}

func GetRecapFromPhotos(photos []*unsplash.UserPhoto) *unsplash.Recap {
	var recap unsplash.Recap
	recap.TotalPhotos = len(photos)
	recap.TotalLikes = 0
	for _, photo := range photos {
		recap.TotalLikes += photo.Likes
		recap.TotalViews += photo.Statistics.Views.Total
		recap.TotalDownloads += photo.Statistics.Downloads.Total
	}

	sortedList := SortByLikes(photos)
	recap.TopPhotos = sortedList[:5]

	return &recap
}

func GetUsernameFromBody(body string) (username string, err error) {
	jsonString := body

	// Replace "\\n" with actual newline characters
	jsonString = strings.ReplaceAll(jsonString, "\\n", "")
	jsonString = strings.ReplaceAll(jsonString, "\\", "")

	// Unmarshal the modified JSON string into a User struct
	var user struct {
		Username string `json:"username"`
	}
	
	err = json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		log.Errorf("error unmarshalling username: %v", err)
		return
	}

	// Access the parsed data
	return user.Username, nil
}

func JSONResponse(statusCode int, message string, body *unsplash.Recap) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    message,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		IsBase64Encoded: false,
	}
}