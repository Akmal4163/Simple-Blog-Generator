package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"

	//"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func read_markdown_files_from_command_line() string {
	//read markdown files from command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please specify input markdown file\n")
		fmt.Println("Example: test.md")
	}

	command_line_arguments := os.Args[1:]

	input_file, err := os.ReadFile(command_line_arguments[0])
	if err != nil {
		fmt.Errorf("cannot read markdown file")
	}

	unsafe_file_content := convert_markdown_text_into_html_data(input_file)
	file_content := bluemonday.UGCPolicy().SanitizeBytes(unsafe_file_content)
	return string(file_content)
}

func convert_markdown_text_into_html_data(markdown_files []byte) []byte {

	//create markdown parser
	markdown_extension := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	markdown_parser := parser.NewWithExtensions(markdown_extension)
	content_data := markdown_parser.Parse(markdown_files)

	//create HTML
	html_flags := html.CommonFlags | html.HrefTargetBlank
	options := html.RendererOptions{Flags: html_flags}
	renderer := html.NewRenderer(options)

	return markdown.Render(content_data, renderer)
}

func send_data_into_server(template_file string) []string {

	markdown_file_content := read_markdown_files_from_command_line()

	data := []string{}
	data = append(data, template_file)
	data = append(data, markdown_file_content)
	return data
}

func run_content_on_server(w http.ResponseWriter, r *http.Request) {
	template_and_content := send_data_into_server("template.html")
	tmpl, err := template.ParseFiles(template_and_content[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Content template.HTML
	}{
		Content: template.HTML(template_and_content[1]),
	}

	err = tmpl.Execute(w, data)
}

func generate_static_html_files() {
	template_and_content := send_data_into_server("template.html")
	tmpl := template.Must(template.ParseFiles(template_and_content[0]))
	data := struct {
		Content template.HTML
	} {
		Content: template.HTML(template_and_content[1]),
	}

	file_output, err := os.Create("output.html")
	if err != nil {
		fmt.Errorf("cannot create file output.html")
		return
	}

	defer file_output.Close()

	tmpl.Execute(file_output, data)
}

func main() {

	static_files := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", static_files))
	http.HandleFunc("/", run_content_on_server)
	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
