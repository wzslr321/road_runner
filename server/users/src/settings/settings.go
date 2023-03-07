package settings

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Server struct {
	Address string `yaml:"address"`
}

type DevMode int64

const (
	Dev DevMode = iota
	Prod
)

func GetServerConf(mode DevMode) *Server {
	serv := &Server{}
	configFile := "config.yaml"
	if mode == Dev {
		configFile = "config-dev.yaml"
	}
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		log.Printf("Failed to get information about caller")
	}
	basePath := filepath.Dir(b)

	dirPath := filepath.Dir(basePath)

	if !ok {
		log.Printf("Failed to get path of config directory")
	}

	path := fmt.Sprintf("%s/conf/%s", dirPath, configFile)

	f, err := os.ReadFile(path)
	if err != nil {
		// fatal is kinda controversial, but I believe it's usage is good here
		// since this function is performed in Init function
		log.Fatalf("Failed to read server config file: %v", err)
	}
	err = yaml.Unmarshal(f, serv)
	if err != nil {
		log.Fatalf("Failed to unmarshal server config file: %v", err)
	}

	return serv
}
