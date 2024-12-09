package config

import (
    "os"
    "time"
)

// Config contiene toda la configuración de la aplicación
type Config struct {
    Firebase FirebaseConfig
    Server   ServerConfig
    Auth     AuthConfig
}

// FirebaseConfig contiene la configuración de Firebase
type FirebaseConfig struct {
    ProjectID              string
    PrivateKeyID          string
    PrivateKey            string
    ClientEmail           string
    ClientID              string
    AuthURI               string
    TokenURI              string
    AuthProviderCertURL   string
    ClientX509CertURL     string
}

// ServerConfig contiene la configuración del servidor
type ServerConfig struct {
    Port         string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
    IdleTimeout  time.Duration
}

// AuthConfig contiene la configuración de autenticación
type AuthConfig struct {
    JWTSecret     string
    TokenDuration time.Duration
}

// LoadConfig carga la configuración desde variables de entorno
func LoadConfig() (*Config, error) {
    return &Config{
        Firebase: FirebaseConfig{
            ProjectID:              os.Getenv("FIREBASE_PROJECT_ID"),
            PrivateKeyID:          os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
            PrivateKey:            os.Getenv("FIREBASE_PRIVATE_KEY"),
            ClientEmail:           os.Getenv("FIREBASE_CLIENT_EMAIL"),
            ClientID:              os.Getenv("FIREBASE_CLIENT_ID"),
            AuthURI:               os.Getenv("FIREBASE_AUTH_URI"),
            TokenURI:              os.Getenv("FIREBASE_TOKEN_URI"),
            AuthProviderCertURL:   os.Getenv("FIREBASE_AUTH_PROVIDER_CERT_URL"),
            ClientX509CertURL:     os.Getenv("FIREBASE_CLIENT_CERT_URL"),
        },
        Server: ServerConfig{
            Port:         getEnvOrDefault("SERVER_PORT", "8080"),
            ReadTimeout:  getDurationOrDefault("SERVER_READ_TIMEOUT", 15*time.Second),
            WriteTimeout: getDurationOrDefault("SERVER_WRITE_TIMEOUT", 15*time.Second),
            IdleTimeout:  getDurationOrDefault("SERVER_IDLE_TIMEOUT", 60*time.Second),
        },
        Auth: AuthConfig{
            JWTSecret:     os.Getenv("JWT_SECRET"),
            TokenDuration: getDurationOrDefault("TOKEN_DURATION", 24*time.Hour),
        },
    }, nil
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
    if value := os.Getenv(key); value != "" {
        if duration, err := time.ParseDuration(value); err == nil {
            return duration
        }
    }
    return defaultValue
}
