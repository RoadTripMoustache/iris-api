package response

// Configurations defines the API response structure for configuration which can be shared to the frontend.
type Configurations struct {
	MaxImagesPerIdea    int      `json:"max_images_per_idea"`
	MaxImagesPerComment int      `json:"max_images_per_comment"`
	MaxSize             int64    `json:"max_size"`
	AcceptedExtensions  []string `json:"accepted_extensions"`
}
