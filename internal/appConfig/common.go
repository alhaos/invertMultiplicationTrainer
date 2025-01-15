package appConfig

import (
	"flag"
	"github.com/alhaos/invertMultiplicationTrainer/internal/webServer"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config general appConfig struct
type Config struct {
	WebServer webServer.Config
}

// New creates the Config from the configuration file specified in the appConfig parameter
// return pointer to Config
func New() (*Config, error) {

	filenamePointer := flag.String("appConfig", "appConfig.yml", "invertMultiplicationTrainer appConfig file")
	flag.Parse()
	filename := *filenamePointer

	var c Config

	err := cleanenv.ReadConfig(filename, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
