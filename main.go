package main

import (
	"flag"
	"fmt"
	"os"

	"yolo/pkg/service"
)

func main() {
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	projectName := initCommand.String("name", "default", "Name of the project")
	projectPort := initCommand.Int("port", 8080, "Port of the project")

	if len(os.Args) < 2 {
		fmt.Println("expected 'init' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initCommand.Parse(os.Args[2:])
	default:
		fmt.Println("expected 'init' subcommand")
		os.Exit(1)
	}

	if *projectName == "" {
		fmt.Println("name is required")
		os.Exit(1)
	}

	if *projectPort <= 0 {
		fmt.Println("port is required")
		os.Exit(1)
	}

	rander := service.NewRander(*projectName, *projectPort)
	if err := rander.InitDir(); err != nil {
		fmt.Printf("generate project dir failed: %v\n", err)
		os.Exit(1)
	}

	if err := rander.InitPkg(); err != nil {
		fmt.Printf("generate template failed: %v\n", err)
		os.Exit(1)
	}

	if err := rander.RunGoMod(); err != nil {
		fmt.Printf("run go mod failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("project init success")
}
