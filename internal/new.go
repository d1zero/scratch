package internal

import (
	"fmt"
	"github.com/d1zero/scratch/internal/pkg"
	"github.com/d1zero/scratch/templates"
	"github.com/spf13/cobra"
	"os"
)

func New(cmd *cobra.Command, args []string) {
	serviceName := args[0]

	err := os.MkdirAll(fmt.Sprintf("%s/cmd/app", serviceName), 0777)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(fmt.Sprintf("%s/internal/app", serviceName), 0777)
	if err != nil {
		panic(err)
	}

	pkg.WriteToFile(fmt.Sprintf("%s/cmd/app/main.go", serviceName), templates.MainFileTemplate, templates.MainTemplateData{
		ProjectName: serviceName,
	})

	pkg.WriteToFile(fmt.Sprintf("%s/internal/app/app.go", serviceName), templates.AppTemplate, struct{}{})

	pkg.WriteToFile(fmt.Sprintf("%s/go.mod", serviceName), templates.GoModTemplate, templates.GoModData{
		ModuleName: serviceName,
	})
}
