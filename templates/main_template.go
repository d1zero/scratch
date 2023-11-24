package templates

type (
	MainTemplateData struct {
		ProjectName string
	}
)

const MainFileTemplate = `package main

import "{{.ProjectName}}/internal/app"

func main() {
	app.Run()
}
	`
