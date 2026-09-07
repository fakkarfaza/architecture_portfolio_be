package model

type ProjectDetail struct {
	ID                int            `json:"id"`
	Slug              string         `json:"slug"`
	Title             string         `json:"title"`
	Category          *string        `json:"category,omitempty"`
	Location          *string        `json:"location,omitempty"`
	Year              *int           `json:"year,omitempty"`
	ThumbnailImageURL *string        `json:"thumbnail_image_url,omitempty"`
	HeroImageURL      *string        `json:"hero_image_url,omitempty"`
	ConceptImageURL   *string        `json:"concept_image_url,omitempty"`
	Brief             *string        `json:"brief,omitempty"`
	ConceptTitle      *string        `json:"concept_title,omitempty"`
	Concept           *string        `json:"concept,omitempty"`
	Gallery           []ProjectImage `json:"gallery"`
}
