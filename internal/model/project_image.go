package model

type ProjectImage struct {
	ID        int     `json:"id"`
	ImageURL  string  `json:"image_url"`
	ImageType *string `json:"image_type,omitempty"`
	SortOrder *int    `json:"sort_order,omitempty"`
}
