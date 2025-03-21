# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)


tidy:
	go mod tidy


build:
	GOOS=windows GOARCH=amd64 go build -o LiveStreamRecorder.exe cmd/main.go