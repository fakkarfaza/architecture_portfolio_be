package model

type Project struct {
	ID                int    `json:"id"`
	Slug              string `json:"slug"`
	Title             string `json:"title"`
	Category          string `json:"category"`
	Location          string `json:"location"`
	Year              int    `json:"year"`
	ThumbnailImageURL string `json:"thumbnail_image_url"`
	HeroImageURL      string `json:"hero_image_url"`
}
