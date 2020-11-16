package main

// Do not write this config to logs or stdout, it contains secrets!
type ImageBuilderConfig struct {
	ListenAddress   string  `env:"LISTEN_ADDRESS"`
	LogLevel        string  `env:"LOG_LEVEL"`
	LogGroup        *string `env:"CW_LOG_GROUP"`
	Region          *string `env:"CW_AWS_REGION"`
	AccessKeyID     *string `env:"CW_AWS_ACCESS_KEY_ID"`
	SecretAccessKey *string `env:"CW_AWS_SECRET_ACCESS_KEY"`
	OsbuildURL      string  `env:"OSBUILD_URL"`
	OsbuildCert     *string  `env:"OSBUILD_CERT_PATH"`
	OsbuildKey      *string  `env:"OSBUILD_KEY_PATH"`
	OsbuildCA       *string  `env:"OSBUILD_CA_PATH"`
}
