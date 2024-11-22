package configs

type (
	Config struct {
		Service       Service       `mapstructure:"service"`
		Database      Database      `mapstructure:"database"`
		SpotifyConfig SpotifyConfig `mapstructure:"spotify_config"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJWT string `mapstructure:"secret_jwt"`
	}

	Database struct {
		DatabaseSourceName string `mapstructure:"db_source_name"`
	}

	SpotifyConfig struct {
		ClientID     string `mapstructure:"client_id"`
		ClientSecret string `mapstructure:"client_secret"`
	}
)
