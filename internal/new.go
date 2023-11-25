package internal

import (
	"fmt"
	"github.com/d1zero/scratch/internal/models"
	"github.com/d1zero/scratch/internal/pkg"
	"github.com/d1zero/scratch/templates"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

func New(cmd *cobra.Command, args []string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	serviceName := args[0]

	postgres, err := cmd.Flags().GetBool("postgres")
	if err != nil {
		logger.Error("error while getting postgres param: %s", err)
	}

	allFlags := models.AllFlags{
		Postgres: postgres,
	}

	err = os.MkdirAll(fmt.Sprintf("%s/cmd/app", serviceName), 0777)
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

	pkg.WriteToFile(fmt.Sprintf("%s/internal/app/app.go", serviceName), templates.BuildAppTemplate(allFlags), struct{}{})

	pkg.WriteToFile(fmt.Sprintf("%s/internal/app/config.go", serviceName), templates.BuildConfigTemplate(allFlags), templates.GoModData{
		ModuleName: serviceName,
	})

	pkg.WriteToFile(fmt.Sprintf("%s/go.mod", serviceName), templates.GoModTemplate, templates.GoModData{
		ModuleName: serviceName,
	})

	pkg.WriteToFile(fmt.Sprintf("%s/.gitignore", serviceName), templates.GitIgnoreTemplate, struct{}{})
	pkg.WriteToFile(fmt.Sprintf("%s/Makefile", serviceName), templates.BuildMakefileTemplate(allFlags), struct{}{})

	if allFlags.Postgres {
		err = os.MkdirAll(fmt.Sprintf("%s/db/migrations", serviceName), 0777)
		if err != nil {
			panic(err)
		}

		pkg.WriteToFile(fmt.Sprintf("%s/db/migrations/000001_initial.up.sql", serviceName), "", struct{}{})
		pkg.WriteToFile(fmt.Sprintf("%s/db/migrations/000001_initial.down.sql", serviceName), "", struct{}{})
	}
}
