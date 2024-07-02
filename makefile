# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

cobra:
	go run viper-example/cobra/main.go
cobra-help:
	go run viper-example/cobra/main.go --help
cobra-echo-help:
	go run viper-example/cobra/main.go echo --help

cobra-echo:
	go run viper-example/cobra/main.go echo "Hello, World"



viper-help:
	go run viper-example/viper/main.go --help

