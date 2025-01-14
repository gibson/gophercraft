package home

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gibson/gophercraft/home/config"
	"github.com/gibson/gophercraft/home/rpcnet"
	"github.com/gibson/gophercraft/utils/log"
	"github.com/gibson/gophercraft/vsn"
)

func getConfigPath() string {
	gc := config.DefaultLocation()
	path := filepath.Join(gc, "Home")
	return path
}

func getConfig() (*config.Home, error) {
	path := getConfigPath()

	conf, err := config.LoadHome(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warn("No Homeserver configuration found at ", path)
			log.Warn("You can create one using gcraft_wizard")
		}
		return nil, err
	}

	fp, err := rpcnet.GetCertFileFingerprint(filepath.Join(conf.Dir, "cert.pem"))
	if err != nil {
		panic(err)
	}

	log.Println("This server's fingerprint is", fp)

	if conf.HostExternal == "" {
		conf.HostExternal = "localhost"
	}

	return conf, nil
}

func Run() {
	vsn.PrintBanner()

	conf, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(RunServer(conf))
}
