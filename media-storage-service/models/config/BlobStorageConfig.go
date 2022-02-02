package config

type BlobStorageConfig struct {
	Endpoint  string `json:"endpoint"`
	Key       string `json:"key"`
	Container string `json:"container"`
}
