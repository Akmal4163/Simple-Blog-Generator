package experiment

import (
	"fmt"
	"os"
	"path/filepath"
)

func ProcessCommandLineArguments() {

	help_command := `USAGE:	to create a new project, type: new <project-names>
	to run project on localhost, type: run <project-names> <port>
	to generate static HTML, CSS, and assets from project, type: build <project-names>`

	if len(os.Args) < 2 {

		fmt.Println("\n" + help_command + "\n")
		os.Exit(2)

	} else if os.Args[1] == "new" {

		//defining main directory
		project_name := os.Args[2]
		base_filepath := filepath.Join(".", project_name)

		CreateNewProjectFromCommandLine(project_name, base_filepath)

	} else if os.Args[1] == "run" {

		project_name := os.Args[2]
		base_filepath := filepath.Join(".", project_name)
		project_run_port := os.Args[3]

		RunProjectFromCommandLine(project_name, base_filepath, project_run_port)

	} else if os.Args[1] == "build" {

		project_name := os.Args[2]
		base_filepath := filepath.Join(".", project_name)

		GenerateStaticFilesFromProject(base_filepath)
		fmt.Println("build sites succesfully for", project_name)

	} else {

		fmt.Println("\n" + help_command + "\n")
		os.Exit(2)

	}
}
