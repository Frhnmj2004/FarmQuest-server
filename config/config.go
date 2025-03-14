package config

import (
	"log"
	"os"

	"github.com/Netflix/go-env"
)

type Config struct {
	Solana struct {
		RpcUrl             string `env:"SOLANA_RPC_URL,default=https://api.mainnet-beta.solana.com"`
		FeePayerPrivateKey string `env:"SOLANA_FEE_PAYER_PRIVATE_KEY,default=BETHNQZirmcNfGYfJYWYboxgHfg98kLfjvjPixSXwAXnV8f4gWcFhxzMe6DZs3MBG5MNAomkevTh2x8gcEwwfyW"`
		FeePayerPublicKey  string `env:"SOLANA_FEE_PAYER_PUBLIC_KEY,default=CjWL3teESBJ5bXnH6c2XoQmuHZ84Rm3K1R8WBxP7no5E"`
	}

	Relayer struct {
		Port int `env:"RELAYER_PORT,default=8080"`
	}

	Encoder struct {
		Port int `env:"ENCODER_PORT,default=8081"`
	}

	Rugcheck struct {
		BaseUrl string `env:"RUGCHECK_BASE_URL,default=https://api.rugcheck.xyz/v1"`
	}

	WebApi struct {
		Name string `env:"WEB_API_NAME,default=plena-api"`
		Port int    `env:"WEB_API_PORT,default=8888"`
	}

	Environment string `env:"ENVIRONMENT,default=development"`

	Log struct {
		Enable bool `env:"LOG_ENABLE,default=true"`
	}

	Version struct {
		Name string `env:"API_VERSION,default=unset"`
		Hash string `env:"GIT_COMMIT_HASH,default=unset"`
	}

	Redis struct {
		URL string `env:"REDIS_URL,default=redis://localhost:6379/11"`
	}

	Extras env.EnvSet
}

func LoadConfig() *Config {
	var cfg Config
	es, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Remaining environment variables.
	cfg.Extras = es
	return &cfg
}
