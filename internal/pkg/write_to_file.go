package pkg

import (
	"bufio"
	"bytes"
	"os"
	"text/template"
)

func WriteToFile[T any](filename, templateValue string, data T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	tmpl := template.Must(template.New(filename).Parse(templateValue))
	err := tmpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
	writer.Flush()

	if err := os.WriteFile(filename, b.Bytes(), 0666); err != nil {
		panic(err)
	}
}
