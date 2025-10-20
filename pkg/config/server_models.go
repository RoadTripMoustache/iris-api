package config

type ServerConfig struct {
	Expose         string   `yaml:"expose" envconfig:"SERVER_EXPOSE"`
	MetricsExpose  string   `yaml:"metrics_expose" envconfig:"SERVER_METRICS_EXPOSE"`
	OriginsAllowed []string `yaml:"origins_allowed"`
	ImagesDir      string   `yaml:"images_dir" envconfig:"SERVER_IMAGES_DIR"`
	ImagesBaseURL  string   `yaml:"images_base_url" envconfig:"SERVER_IMAGES_BASE_URL"`
}
