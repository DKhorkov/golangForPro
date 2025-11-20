package config

import (
	"time"

	"github.com/DKhorkov/libs/loadenv"
)

func New() Config {
	return Config{
		HTTP: HTTPConfig{
			Host: loadenv.GetEnv("HOST", "0.0.0.0"),
			Port: loadenv.GetEnvAsInt("PORT", 8080),
			ReadTimeout: time.Second * time.Duration(
				loadenv.GetEnvAsInt("HTTP_READ_TIMEOUT", 1),
			),
			IdleTimeout: time.Second * time.Duration(
				loadenv.GetEnvAsInt("HTTP_IDLE_TIMEOUT", 10),
			),
			WriteTimeout: time.Second * time.Duration(
				loadenv.GetEnvAsInt("HTTP_WRITE_TIMEOUT", 1),
			),
		},
		CORS: CORSConfig{
			AllowedOrigins:   loadenv.GetEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}, ", "),
			AllowedMethods:   loadenv.GetEnvAsSlice("CORS_ALLOWED_METHODS", []string{"*"}, ", "),
			AllowedHeaders:   loadenv.GetEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"*"}, ", "),
			AllowCredentials: loadenv.GetEnvAsBool("CORS_ALLOW_CREDENTIALS", true),
			MaxAge:           loadenv.GetEnvAsInt("CORS_MAX_AGE", 600),
		},
		Filepath: FilepathConfig{
			Path:   loadenv.GetEnv("PHONEBOOK_PATH", "phonebook.json"),
			Source: loadenv.GetEnv("PHONEBOOK_SOURCE", "json"),
		},
	}
}

type HTTPConfig struct {
	Host         string
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	MaxAge           int
	AllowCredentials bool
}

type FilepathConfig struct {
	Path   string
	Source string
}

type Config struct {
	HTTP     HTTPConfig
	CORS     CORSConfig
	Filepath FilepathConfig
}
