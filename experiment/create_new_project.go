package experiment

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateNewProjectFromCommandLine(project_name string, base_filepath string) {

	//create layout directory
	layout_path := filepath.Join(base_filepath, "layout")
	err_layout := os.MkdirAll(layout_path, 0777)
	if err_layout != nil {
		log.Fatal("ERROR !!! cannot create layout folder for your project")
	} else {
		fmt.Println("layout folder created...")
	}

	//create content directory
	content_path := filepath.Join(base_filepath, "content")
	err_content := os.MkdirAll(content_path, 0777)
	if err_content != nil {
		log.Fatal("ERROR !!! cannot create content folder for your project")
	} else {
		fmt.Println("content folder created...")
	}

	//create output directory
	output_path := filepath.Join(base_filepath, "output")
	err_output := os.MkdirAll(output_path, 0777)
	if err_output != nil {
		log.Fatal("ERROR !!! cannot create output folder for your project")
	} else {
		fmt.Println("output folder created...")
	}

	//create static files directory
	static_file_path := filepath.Join(base_filepath, "static")
	static_file_err := os.MkdirAll(static_file_path, 0777)
	if static_file_err != nil {
		log.Fatal("ERROR !!! cannot create static folder for your project")
	} else {
		fmt.Println("static folder created...")
	}

	//create config YAML files
	config_directory := filepath.Join(base_filepath, "config.yaml")
	config_content := "# this is a config file, you can set routes, HTML, and markdown files here\n" +
		"routes:\n" + " - path: / \n" + "   template: index.html\n" + "   content: index.md\n"
	config_err := os.WriteFile(config_directory, []byte(config_content), 0660)
	if config_err != nil {
		log.Fatal("ERROR !!! cannot create config.yaml files")
	} else {
		fmt.Println("config.yaml file created...")
	}

	//create index.html files
	html_starter_directory := filepath.Join(layout_path, "index.html")
	html_data := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
		<link rel="stylesheet" href="./static/style.css">
	</head>
	<body>
		<h1>you can edit this HTML for your site layout</h1>
		<div class="content">
			{{ .Main_Content }}
		</div>
	</body>
	</html>`
	html_err := os.WriteFile(html_starter_directory, []byte(html_data), 0660)
	if html_err != nil {
		log.Fatal("ERROR !!! cannot create index.html files")
	} else {
		fmt.Println("index.html file created...")
	}

	//create style.css files
	css_starter_directory := filepath.Join(base_filepath, "static", "style.css")
	css_data := `* { background-color: azure;
    font-family: sans-serif;
    color: black; 
	}`
	css_err := os.WriteFile(css_starter_directory, []byte(css_data), 0660)
	if css_err != nil {
		log.Fatal("ERROR !!! cannot create style.css files")
	} else {
		fmt.Println("style.css file created...")
	}

	//create index.md files
	markdown_starter_directory := filepath.Join(content_path, "index.md")
	markdown_err := os.WriteFile(markdown_starter_directory, []byte("# you can place content here"), 0660)
	if markdown_err != nil {
		log.Fatal("ERROR !!! cannot create index.md files")
	} else {
		fmt.Println("index.md file created...")
	}
}
