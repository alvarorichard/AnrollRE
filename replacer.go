package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func replaceBody(s string) string {
	lines := strings.Split(s, "\n")
	newBody := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			newBody = newBody + line + "\n"
			continue
		}
		b64 := base64.StdEncoding.EncodeToString([]byte(line))
		newLine := fmt.Sprintf("http://localhost:8000/embed?url=%s", b64)
		newBody = newBody + newLine + "\n"
	}
	return newBody

}
