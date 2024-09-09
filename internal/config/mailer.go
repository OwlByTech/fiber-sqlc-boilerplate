package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ConvertAllMailingTemplates() error {
	path := "./internal/service/mailing/templates"
	items, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error read path templates %v", err)
	}

	for _, item := range items {
		filename := item.Name()
		split := strings.Split(filename, ".")
		if len(split) != 2 {
			continue
		}

		if split[1] != "mjml" {
			continue
		}

		inputPath := fmt.Sprintf("%s/%s", path, filename)
		outputPath := fmt.Sprintf("%s/%s.html", path, split[0])
		err := ConvertMJMLToHTML(inputPath, outputPath)

		if err != nil {
			return err
		}
	}

	return nil
}

func ConvertMJMLToHTML(inputPath string, outputPath string) error {
	cmd := exec.Command("mjml", "-r", inputPath, "-o", outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting MJML to HTML: %w", err)
	}
	return nil
}
