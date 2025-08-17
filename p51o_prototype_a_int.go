package main

import (
	"encoding/json"
	"fmt"
)

// Stage represents a single stage in the pipeline
type Stage struct {
	Name  string `json:"name"`
	Type  string `json:"type"` // e.g. "build", "deploy", "test"
	Config map[string]string `json:"config"`
}

// Pipeline represents a DevOps pipeline
type Pipeline struct {
	Name    string   `json:"name"`
	Stages  []Stage  `json:"stages"`
	Triggers []Trigger `json:"triggers"`
}

// Trigger represents a trigger for the pipeline
type Trigger struct {
	Type  string `json:"type"` // e.g. "git-push", "schedule"
	Config map[string]string `json:"config"`
}

// Generator represents the pipeline generator
type Generator struct {
	Templates map[string]Pipeline `json:"templates"`
}

func (g *Generator) AddTemplate(name string, pipeline Pipeline) {
	g.Templates[name] = pipeline
}

func (g *Generator) GeneratePipeline(name string) (Pipeline, error) {
	if tmpl, ok := g.Templates[name]; ok {
		return tmpl, nil
	}
	return Pipeline{}, fmt.Errorf("template not found: %s", name)
}

func main() {
	g := &Generator{
		Templates: map[string]Pipeline{},
	}

	g.AddTemplate("my-pipeline", Pipeline{
		Name: "my-pipeline",
		Stages: []Stage{
			{
				Name: "build",
				Type: "build",
				Config: map[string]string{
					"image": "golang:alpine",
				},
			},
			{
				Name: "deploy",
				Type: "deploy",
				Config: map[string]string{
					"environment": "prod",
				},
			},
		},
		Triggers: []Trigger{
			{
				Type: "git-push",
				Config: map[string]string{
					"repo": "https://github.com/my-org/my-repo",
				},
			},
		},
	},
	})

	pipeline, err := g.GeneratePipeline("my-pipeline")
	if err != nil {
		fmt.Println(err)
		return
	}

	json, err := json.MarshalIndent(pipeline, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(json))
}