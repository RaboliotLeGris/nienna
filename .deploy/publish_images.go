package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var author string
	if author = os.Getenv("NIENNA_AUTHOR"); author == "" {
		log.Fatal("version is mandatory, please fill 'NIENNA_AUTHOR'")
	}

	image_prefix := "nienna"

	var version string
	if version = os.Getenv("NIENNA_IMAGE_VERSION"); version == "" {
		log.Fatal("version is mandatory, please fill 'NIENNA_IMAGE_VERSION'")
	}

	images_to_publish := []string{"backburner", "cliff", "db", "pulsar", "webapp"}
	log.Print("Publishing Docker Images: ", images_to_publish)
	log.Print("Version to publish: ", version)

	current_dir, _ := os.Getwd()
	log.Print("Current dir: ", current_dir)

	for _, image := range images_to_publish {
		log.Print("\nPublishing image: ", image)
		if err := publish(image, author, image_prefix, version); err != nil {
			log.Fatal("Failed to publish image with error :", err)
		}
		log.Printf("Publishing image %s done\n", image)
	}
}

func publish(image, author, prefix, version string) error {
	local_version_tag := fmt.Sprintf("%s_%s:%s", prefix, image, version)             // prefix_image:version
	remote_version_tag := fmt.Sprintf("%s/%s_%s:%s", author, prefix, image, version) // author/prefix_image:version
	remote_latest_tag := fmt.Sprintf("%s/%s_%s:latest", author, prefix, image)       // author/prefix_image:latest

	log.Printf("Publishing image with local version tag: %s - remote version tag: %s - remote latest tag: %s", local_version_tag, remote_version_tag, remote_latest_tag)

	build_cmd := []string{"build", image, "-t", local_version_tag}
	tag_version_cmd := []string{"tag", local_version_tag, remote_version_tag}
	tag_latest_cmd := []string{"tag", local_version_tag, remote_latest_tag}

	log.Print("Building image")
	cmd := exec.Command("docker", build_cmd...)
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Print("Tagging to version")
	cmd = exec.Command("docker", tag_version_cmd...)
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Print("Tagging to latest version")
	cmd = exec.Command("docker", tag_latest_cmd...)
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Print("Pushing image with version")
	cmd = exec.Command("docker", "push", remote_version_tag)
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Print("Pushing image with latest version")
	cmd = exec.Command("docker", "push", remote_latest_tag)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
