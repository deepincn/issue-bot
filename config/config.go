package config

// Yaml struct of yaml
type Yaml struct {
	Version uint `yaml:"version"`
	Github  *struct {
		Token *string `yaml:"token,omitempty"`
	} `yaml:"github"`
	Database *struct {
		URL  *string `yaml:"url,omitempty"`
		Auth *struct {
			Name     *string `yaml:"name,omitempty"`
			Password *string `yaml:"password,omitempty"`
		}
	} `yaml:"database"`
}
