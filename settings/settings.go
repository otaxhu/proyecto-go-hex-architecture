package settings

import (
	_ "embed"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var file []byte

func NewDatabase() *Database {
	settings := &Database{}
	if err := yaml.Unmarshal(file, settings); err != nil {
		panic(errors.Wrap(err, "settings.NewDatabase()"))
	}
	return settings
}

func NewServer() *Server {
	settings := &Server{}
	if err := yaml.Unmarshal(file, settings); err != nil {
		panic(errors.Wrap(err, "settings.NewServer()"))
	}
	return settings
}
