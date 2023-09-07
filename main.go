package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Project struct {
	URL string `json:"url"`
}

func main() {
	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		projectName := r.URL.Query().Get("name")

		projectURLs, err := GetProjectURLs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var filteredProjects []Project
		for _, projectURL := range projectURLs {
			if projectName == "" || strings.Contains(projectURL, projectName) {
				filteredProjects = append(filteredProjects, Project{URL: projectURL})
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string][]Project{"projects": filteredProjects})
	})

	http.ListenAndServe(":8080", nil)
}
