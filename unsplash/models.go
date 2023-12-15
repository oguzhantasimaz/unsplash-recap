package unsplash

import (
	"time"
)

type Recap struct {
	TotalPhotos    int          `json:"total_photos"`
	TotalViews     int          `json:"total_views"`
	TotalDownloads int          `json:"total_downloads"`
	TotalLikes     int          `json:"total_likes"`
	TopPhotos      []*UserPhoto `json:"top_photos"`
}

type UserPhoto struct {
	Id             string    `json:"id"`
	Slug           string    `json:"slug"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Width          int       `json:"width"`
	Height         int       `json:"height"`
	Color          string    `json:"color"`
	BlurHash       string    `json:"blur_hash"`
	Description    string    `json:"description"`
	AltDescription string    `json:"alt_description"`
	Urls           struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
		SmallS3 string `json:"small_s3"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		Html             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	Likes      int `json:"likes"`
	Statistics struct {
		Downloads struct {
			Total      int `json:"total"`
			Historical struct {
				Change     int    `json:"change"`
				Resolution string `json:"resolution"`
				Quantity   int    `json:"quantity"`
				Values     []struct {
					Date  string `json:"date"`
					Value int    `json:"value"`
				} `json:"values"`
			} `json:"historical"`
		} `json:"downloads"`
		Views struct {
			Total      int `json:"total"`
			Historical struct {
				Change     int    `json:"change"`
				Resolution string `json:"resolution"`
				Quantity   int    `json:"quantity"`
				Values     []struct {
					Date  string `json:"date"`
					Value int    `json:"value"`
				} `json:"values"`
			} `json:"historical"`
		} `json:"views"`
		Likes struct {
			Total      int `json:"total"`
			Historical struct {
				Change     int    `json:"change"`
				Resolution string `json:"resolution"`
				Quantity   int    `json:"quantity"`
				Values     []struct {
					Date  string `json:"date"`
					Value int    `json:"value"`
				} `json:"values"`
			} `json:"historical"`
		} `json:"likes"`
	} `json:"statistics"`
}

type Photo struct {
	Id             string    `json:"id"`
	Slug           string    `json:"slug"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Width          int       `json:"width"`
	Height         int       `json:"height"`
	BlurHash       string    `json:"blur_hash"`
	Description    string    `json:"description"`
	AltDescription string    `json:"alt_description"`
	Urls           struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
		SmallS3 string `json:"small_s3"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		Html             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	Likes            int `json:"likes"`
	TopicSubmissions struct {
		Wallpapers struct {
			Status     string    `json:"status"`
			ApprovedOn time.Time `json:"approved_on"`
		} `json:"wallpapers"`
		ArchitectureInterior struct {
			Status     string    `json:"status"`
			ApprovedOn time.Time `json:"approved_on"`
		} `json:"architecture-interior"`
	} `json:"topic_submissions"`
	Location struct {
		Name    *string `json:"name"`
		City    *string `json:"city"`
		Country *string `json:"country"`
	} `json:"location"`
	TagsPreview []struct {
		Type  string `json:"type"`
		Title string `json:"title"`
	} `json:"tags_preview"`
	Views     int `json:"views"`
	Downloads int `json:"downloads"`
	Topics    []struct {
		Id         string `json:"id"`
		Title      string `json:"title"`
		Slug       string `json:"slug"`
		Visibility string `json:"visibility"`
	} `json:"topics"`
}
