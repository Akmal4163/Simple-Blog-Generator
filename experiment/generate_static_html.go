package experiment

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

type Route struct {
	Path     string `yaml:"path"`
	Template string `yaml:"template"`
	Content  string `yaml:"content"`
}

type RoutesConfig struct {
	Routes []Route `yaml:"routes"`
}

func read_and_parse_yaml_config(base_filepath string, yaml_config_file string) RoutesConfig {
	config_directory := filepath.Join(base_filepath, yaml_config_file)

	//read config.yaml files
	file, err := os.ReadFile(config_directory)
	if err != nil {
		log.Fatalf("failed to open config.yaml file: %s", err)
	}

	//parse config.yaml file
	var config RoutesConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("failed to parse config.yaml files: %s", err)
	}

	return config
}

func convert_markdown_to_html(md []byte) []byte {

	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	md_content := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(md_content, renderer)
}

func find_html_files(base_filepath string, html_path string) string {
	template_dir := filepath.Join(base_filepath, "layout", html_path)
	return template_dir
}

func read_and_parse_markdown_files(base_filepath string, md_path string) string {

	markdown_directory := filepath.Join(base_filepath, "content", md_path)
	markdown, err := os.ReadFile(markdown_directory)
	if err != nil {
		fmt.Println("cannot convert markdown files into HTML, ", err)
	}
	md := []byte(markdown)
	html := convert_markdown_to_html(md)

	sanitized_html := bluemonday.UGCPolicy().SanitizeBytes(html)

	return string(sanitized_html)
}

func GenerateStaticFilesFromProject(base_filepath string) {

	yaml_file_config := read_and_parse_yaml_config(base_filepath, "config.yaml")

	for i := 0; i < len(yaml_file_config.Routes); i++ {
		markdown_content := read_and_parse_markdown_files(base_filepath, yaml_file_config.Routes[i].Content)
		template_content := find_html_files(base_filepath, yaml_file_config.Routes[i].Template)
		tmpl := template.Must(template.ParseFiles(template_content))

		data := struct {
			Main_Content template.HTML
		}{
			Main_Content: template.HTML(markdown_content),
		}

		output_files := filepath.Join(base_filepath, "output", yaml_file_config.Routes[i].Template)
		file_output, err := os.Create(output_files)
		if err != nil {
			fmt.Println("cannot create HTML output files, ", err)
			break
		}

		defer file_output.Close()

		tmpl.Execute(file_output, data)

	}
}
