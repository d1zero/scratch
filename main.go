package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/d1zero/scratch/templates"
)

func main() {
	taskPtr := flag.String("service-name", "", "New service name")
	flag.Parse()

	if *taskPtr == "" {
		fmt.Println("Available flags:")
		flag.PrintDefaults()
		return
	}

	err := os.MkdirAll(fmt.Sprintf("%s/cmd/app", *taskPtr), 0777)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(fmt.Sprintf("%s/internal/app", *taskPtr), 0777)
	if err != nil {
		panic(err)
	}

	data := templates.MainTemplateData{
		ProjectName: *taskPtr,
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	tmpl := template.Must(template.New("cmd/app/main.go").Parse(templates.MainFileTemplate))
	err = tmpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
	writer.Flush()

	if err := os.WriteFile("cmd/app/main.go", b.Bytes(), 0666); err != nil {
		panic(err)
	}

	var goMod bytes.Buffer
	writer = bufio.NewWriter(&goMod)
	goModData := templates.GoModData{
		ModuleName: *taskPtr,
	}
	tmpl = template.Must(template.New("go.mod").Parse(templates.GoModTemplate))
	err = tmpl.Execute(writer, goModData)
	if err != nil {
		panic(err)
	}
	writer.Flush()

	if err := os.WriteFile("go.mod", goMod.Bytes(), 0666); err != nil {
		panic(err)
	}

	if err := os.WriteFile("internal/app/app.go", []byte(templates.AppTemplate), 0666); err != nil {
		panic(err)
	}
}
