package main

import (
	"github.com/ContainerSolutions/validation-admission-controller-go/server"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)


type Config struct {
	ListenOn string `default:"0.0.0.0:8080"`
	TlsCert  string `default:"/etc/webhook/certs/cert.pem"`
	TlsKey   string `default:"/etc/webhook/certs/key.pem"`
	Debug    bool   `default:"true"`
}

func main() {
	config := &Config{}
	envconfig.Process("", config)

	if config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Infoln(config)
	nsac := server.NamespaceAdmission{}
	s := server.GetAdmissionValidationServer(&nsac, config.TlsCert, config.TlsKey, config.ListenOn)
	s.ListenAndServeTLS("", "")
}
