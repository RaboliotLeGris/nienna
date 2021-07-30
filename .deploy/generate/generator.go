package main

import (
	"io"

	"gopkg.in/yaml.v3"
)

/// Docker-compose file structure

type compose struct {
	Version  string             `yaml:"version"`
	Services map[string]service `yaml:"services"`
	Volumes  []string           `yaml:"volumes"`
}

type service struct {
	Image       string            `yaml:"image,omitempty"`
	Build       string            `yaml:"build,omitempty"`
	Restart     string            `yaml:"restart"`
	StdINOpen   *bool             `yaml:"stdin_open,omitempty"`
	TTY         *bool             `yaml:"tty,omitempty"`
	Command     *string           `yaml:"command,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Volumes     []string          `yaml:"volumes,omitempty"`
	DependsOn   []string          `yaml:"depends_on,omitempty"`
	Ports       []string          `yaml:"ports,omitempty"`
}

/// Generator internals
type ComposeElement struct {
	name      string
	build     bool
	buildPath string
	imageName string
	restart   bool
	dev       bool
	command   string
	env       map[string]string
	volumes   []string
	depends   []string
	ports     []string
}

type Generator struct {
	elem []ComposeElement
}

func (c *Generator) Append(elem ComposeElement) {
	c.elem = append(c.elem, elem)
}

func (c *Generator) Generate(params *Params, w io.Writer) error {
	g := compose{
		Version:  "3.9",
		Services: map[string]service{},
	}
	for _, e := range c.elem {
		var service service
		if e.build {
			service.Build = e.buildPath
		} else {
			service.Image = e.imageName
		}

		if e.restart {
			service.Restart = "unless-stopped"
		} else {
			service.Restart = "on-failure"
		}

		service.Command = &e.command

		service.Volumes = e.volumes
		service.DependsOn = e.depends
		service.Environment = e.env
		service.Ports = e.ports

		if e.dev {
			service.Environment["NIENNA_DEV"] = "true"
			if params.Enable_TTY {
				trueP := true
				service.StdINOpen = &trueP
				service.TTY = &trueP
				command := "sh"
				service.Command = &command
			}
		}

		g.Services[e.name] = service
	}

	encoder := yaml.NewEncoder(w)
	if err := encoder.Encode(g); err != nil {
		return err
	}

	return encoder.Close()
}
