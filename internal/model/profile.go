package model

type Profile struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	ProfileImageURL *string `json:"profile_image_url, omitempty"`
	Title           *string `json:"title, omitempty"`
	Bio             *string `json:"bio, omitempty"`
	Description     *string `json:"description, omitempty"`
	Email           *string `json:"email, omitempty"`
	Location        *string `json:"location, omitempty"`
	InstagramURL    *string `json:"instagram_url, omitempty"`
	LinkedInURL     *string `json:"linked_in_url, omitempty"`
	BehanceURL      *string `json:"behance_url, omitempty"`
}
