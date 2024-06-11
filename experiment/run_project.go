package experiment

import (
	"html/template"

	"fmt"
	"net/http"
	"path/filepath"
)

func run_server(base_filepath string, project_port string) {

	yaml_config := read_and_parse_yaml_config(base_filepath, "config.yaml")

	r := http.NewServeMux()
	template_directory := filepath.Join(base_filepath, "output", "*.html")
	templates, err := template.ParseGlob(template_directory)
	if err != nil {
		fmt.Println("cannot read .html files in /output folder")
	}
	for _, route := range yaml_config.Routes {
		route := route
		r.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			err := templates.ExecuteTemplate(w, route.Template, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}

	static_files := http.Dir(filepath.Join(base_filepath, "static"))
	fs := http.StripPrefix("/static/", http.FileServer(static_files))
	r.Handle("/static/", fs)

	port := ":" + project_port
	fmt.Println("server running at port: " + port + " to stop, press CTRL+C")

	http.ListenAndServe(port, r)

}

func RunProjectFromCommandLine(project_name string, base_filepath string, project_port string) {
	GenerateStaticFilesFromProject(base_filepath)
	run_server(base_filepath, project_port)
}
