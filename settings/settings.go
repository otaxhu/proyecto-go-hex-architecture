package settings

import (
	_ "embed"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var file []byte

func NewDatabase() (*Database, error) {
	settings := &Database{}
	if err := yaml.Unmarshal(file, settings); err != nil {
		return nil, errors.Wrap(err, "settings.NewDatabase()")
	}
	return settings, nil
}

func NewServer() (*Server, error) {
	settings := &Server{}
	if err := yaml.Unmarshal(file, settings); err != nil {
		return nil, errors.Wrap(err, "settings.NewServer()")
	}
	return settings, nil
}
