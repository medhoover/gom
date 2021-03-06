/*
config package
Define Configuration structure and related handlers
*/
package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/medhoover/gom/logger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type anyType interface{}

type launcher interface {
	Execute(args []string)
}

type envSetter interface {
	Set()
}

type parser interface {
	parse(path string) (*ConfigInstance, error)
}

// Defining the Configuration type
type ConfigInstance struct {
	Name         string                 `yaml:"name,omitempty"`
	Commands     map[string]command     `yaml:"commands,flow,omitempty"`
	Environments map[string]environment `yaml:"env,flow,omitempty"`
}

// New exports a configuration instance
func New(path string) *ConfigInstance {
	// create a new configuration instance
	var ci *ConfigInstance
	ci, err := ci.parse(path)
	if err != nil {
		logger.Error(err)
	}
	return ci
}

// Read, parse and validate the config file
func (ci *ConfigInstance) parse(path string) (*ConfigInstance, error) {

	// Read config file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		absPath, _ := filepath.Abs(path)
		return nil, errors.Errorf(
			"Unable to read file %s\nCreate a configuration file first",
			absPath,
		)
	}

	// Parse the file to the ConfigInstance value
	err = yaml.Unmarshal([]byte(data), &ci)
	if err != nil {
		return nil, errors.Wrap(err, "File structure")
	}

	return ci, nil
}

// Execute a command by the name as passed in arguments
func (ci *ConfigInstance) Execute(args []string) {

	if command, exist := ci.Commands[args[0]]; exist {
		if err := command.Execute(args[1:]); err != nil {
			logger.Error(
				errors.Wrapf(
					err,
					"Command '%s' Failed",
					strings.Join(args, " "),
				),
			)
		}
	} else {
		logger.Error(errors.Errorf("Command '%s' Failed: Command is not defined", args[0]))
	}
}

// Set an environment by the name
func (ci *ConfigInstance) Set(name string) {

	if env, exist := ci.Environments[name]; exist {
		if err := env.Set(); err != nil {
			logger.Error(
				errors.Wrapf(
					err,
					"Environment '%s' Failed",
					name,
				),
			)
		}
	} else {
		logger.Error(errors.Errorf("Environment '%s' Failed: Environment not defined", name))
	}
}
