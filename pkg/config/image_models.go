package config

type ImageConfig struct {
	MaxImagesPerIdea    int      `yaml:"maxImagesPerIdea" envconfig:"IMAGE_MAX_IMAGES_PER_IDEA"`
	MaxImagesPerComment int      `yaml:"maxImagesPerComment" envconfig:"IMAGE_MAX_IMAGES_PER_COMMENT"`
	MaxSize             int64    `yaml:"maxSize" envconfig:"IMAGE_MAX_SIZE"` // In bites
	AcceptedExtensions  []string `yaml:"acceptedExtensions"`
}
