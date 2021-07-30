package main

import (
	"log"
	"os"
)

func CreateWebappElement(env *Params) ComposeElement {
	return ComposeElement{
		name:      "webapp",
		build:     env.Build,
		buildPath: "webapp/.dev",
		imageName: "raboliotlegris/nienna_webapp",
		restart:   true,
		dev:       env.Dev_mode,
		volumes:   []string{"$PWD/webapp/.dev/Caddyfile:/etc/caddy/Caddyfile"},
		depends:   []string{"loadbalancer"},
	}
}

func CreateCliffElement(env *Params) ComposeElement {
	sessionID := "TO GENERATE PROPERLY"

	return ComposeElement{
		name:      "cliff",
		build:     env.Build,
		buildPath: "cliff/.dev",
		imageName: "raboliotlegris/nienna_cliff",
		restart:   true,
		dev:       env.Dev_mode,
		env: map[string]string{
			"DB_URI":         "postgresql://pg/" + env.DB_database + "?user=" + env.DB_user + "&password=" + env.DB_password,
			"REDIS_URI":      env.Redis_uri,
			"S3_URI":         env.S3_uri,
			"S3_DISABLE_TLS": "true",
			"S3_ACCESS_KEY":  env.S3_key,
			"S3_SECRET_KEY":  env.S3_secret,
			"AMQP_URI":       "amqp://" + env.AMQP_user + ":" + env.AMQP_password + "@rabbitmq:5672",
			"SESSION_KEY":    sessionID,
		},
		volumes: []string{"$PWD/cliff:/go/src/cliff"}, // TODO: correct path
		depends: []string{"db", "redis", "s3", "amqp"},
	}
}

func CreateBackburnerElement(env *Params) ComposeElement {
	return ComposeElement{
		name:      "backburner",
		build:     env.Build,
		buildPath: "backburner/.dev",
		imageName: "raboliotlegris/nienna_backburner",
		restart:   true,
		dev:       env.Dev_mode,
		env: map[string]string{
			"S3_URI":        env.S3_uri,
			"S3_ACCESS_KEY": env.S3_key,
			"S3_SECRET_KEY": env.S3_secret,
			"AMQP_URI":      "amqp://" + env.AMQP_user + ":" + env.AMQP_password + "@rabbitmq:5672",
		},
		volumes: []string{"$PWD/backburner:/usr/src/backburner"},
		depends: []string{"s3", "amqp"},
	}
}

func CreatePulsarElement(env *Params) ComposeElement {
	return ComposeElement{
		name:      "pulsar",
		build:     env.Build,
		buildPath: "pulsar/.dev",
		imageName: "raboliotlegris/nienna_pulsar",
		restart:   true,
		dev:       env.Dev_mode,
		env: map[string]string{
			"DB_PARAMS": "Host=pg;Username=" + env.DB_user + ";Password=" + env.DB_password + ";Database=" + env.DB_database,
			"AMQP_URI":  "amqp://" + env.AMQP_user + ":" + env.AMQP_password + "@rabbitmq:5672",
		},
		volumes: []string{"$PWD/pulsar:/pulsar"},
		depends: []string{"db", "rabbitmq"},
	}
}

func CreateDBElement(env *Params) ComposeElement {
	return ComposeElement{
		name:      "db",
		build:     env.Build,
		buildPath: "db",
		imageName: "raboliotlegris/nienna_db",
		restart:   true,
		dev:       env.Dev_mode,
		env: map[string]string{
			"DB_URI": "postgresql://pg/" + env.DB_database + "?user=" + env.DB_user + "&password=" + env.DB_password,
		},
		depends: []string{"pg"},
	}
}

func CreatePGElement(env *Params) ComposeElement {
	return ComposeElement{
		name:      "pg",
		build:     false,
		imageName: "postgres:13.2@sha256:868a8ec5f24cce648991c8e41811a43acddebdf55aff0b3aad90fab359f2c518",
		restart:   true,
		dev:       false,
		env: map[string]string{
			"POSTGRES_DB":       env.DB_database,
			"POSTGRES_USER":     env.DB_user,
			"POSTGRES_PASSWORD": env.DB_password,
		},
	}
}

func CreateRedisElement(env *Params) ComposeElement {
	return ComposeElement{
		name:      "redis",
		build:     false,
		restart:   true,
		dev:       false,
		imageName: "redis:6.2.1-buster@sha256:2084204018c52ea78ef43302df2d284f46175e3d6218347b58a09e5b97c6e828",
	}
}

func CreateAMQPElement(env *Params) ComposeElement {
	c := ComposeElement{
		name:      "rabbitmq",
		build:     false,
		restart:   true,
		imageName: "rabbitmq:3.8.14-management@sha256:09bb2f34a383403b7c4a320f9037975d811efb92ecbef728d5760c2ab3fc8ca0",
		dev:       false,
		env: map[string]string{
			"RABBITMQ_DEFAULT_USER": env.AMQP_user,
			"RABBITMQ_DEFAULT_PASS": env.AMQP_password,
		},
		volumes: []string{
			"rabbitmq:/var/lib/rabbitmq",
		},
	}

	if env.Dev_mode {
		c.ports = append(c.ports, "15672:15672")
	}

	return c
}

func CreateS3Element(env *Params) ComposeElement {
	c := ComposeElement{
		name:      "s3",
		build:     false,
		restart:   true,
		dev:       false,
		imageName: "minio/minio:RELEASE.2021-07-22T05-23-32Z@sha256:a7ada72685dd94dd63627bf17ccbbff000bc04a978f230084f1156a3e8d9550f",
		env: map[string]string{
			"MINIO_ROOT_USER":     env.S3_key,
			"MINIO_ROOT_PASSWORD": env.S3_secret,
		},
		command: "server /home/shared",
		volumes: []string{
			"s3:/home/shared",
		},
	}

	if env.Dev_mode {
		c.ports = append(c.ports, "9000:9000")
	}

	return c
}

func CreateLoadbalancerElement(env *Params) ComposeElement {
	c := ComposeElement{
		name:      "loadbalancer",
		build:     false,
		restart:   true,
		dev:       false,
		imageName: "nginx:1.21.0@sha256:2f1cd90e00fe2c991e18272bb35d6a8258eeb27785d121aa4cc1ae4235167cfd",
		ports: []string{
			"80:80",
		},
		volumes: []string{
			"$PWD/.docker/services/nginx/nginx.conf:/etc/nginx/nginx.conf",
		},
	}

	if env.LB_enable_tls {
		c.ports = append(c.ports, "443:443")
	}

	return c
}

func main() {
	outputFilename := "docker-compose.yml"
	_ = os.Remove(outputFilename)

	params, err := ParseParams("env.json")
	if err != nil {
		log.Fatal("Failed to parse env.json ", err)
	}

	var g Generator
	g.Append(CreateWebappElement(params))
	g.Append(CreateCliffElement(params))
	g.Append(CreateBackburnerElement(params))
	g.Append(CreatePulsarElement(params))
	g.Append(CreateDBElement(params))
	g.Append(CreatePGElement(params))
	g.Append(CreateRedisElement(params))
	g.Append(CreateAMQPElement(params))
	g.Append(CreateS3Element(params))
	g.Append(CreateLoadbalancerElement(params))

	// opening file to write yaml
	f, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal("Failed to open file", outputFilename, "with error:", err)
	}
	defer f.Close()

	// generating yaml to the given io.Writer
	g.Generate(params, f)
}
