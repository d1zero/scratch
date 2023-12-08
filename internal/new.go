package internal

import (
	"fmt"
	"github.com/d1zero/scratch/internal/models"
	"github.com/d1zero/scratch/internal/pkg"
	"github.com/d1zero/scratch/templates"
	"github.com/spf13/cobra"
	MultipleChoice "github.com/thewolfnl/go-multiplechoice"
	"os"
)

func New(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("error: new service name must not be empty")
		return
	}

	integrations := MultipleChoice.MultiSelection("Choose integration to enable in project: ", []string{
		models.Postgres,
		models.Redis,
		models.Kafka,
		models.Grpc,
		models.Http,
	})

	serviceName := args[0]

	enabledIntegrations := pkg.ParseIntegrations(integrations)

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

	pkg.WriteToFile(fmt.Sprintf("%s/internal/app/app.go", serviceName), templates.BuildAppTemplate(enabledIntegrations), templates.AppTemplateData{
		ProjectName: serviceName,
	})

	pkg.WriteToFile(fmt.Sprintf("%s/internal/app/config.go", serviceName), templates.BuildConfigTemplate(enabledIntegrations), templates.GoModData{
		ModuleName: serviceName,
	})

	pkg.WriteToFile(fmt.Sprintf("%s/go.mod", serviceName), templates.GoModTemplate, templates.GoModData{
		ModuleName: serviceName,
	})

	pkg.WriteToFile(fmt.Sprintf("%s/.gitignore", serviceName), templates.GitIgnoreTemplate, struct{}{})
	pkg.WriteToFile(fmt.Sprintf("%s/Makefile", serviceName), templates.BuildMakefileTemplate(enabledIntegrations), struct{}{})

	if enabledIntegrations.Postgres {
		err = os.MkdirAll(fmt.Sprintf("%s/db/migrations", serviceName), 0777)
		if err != nil {
			panic(err)
		}

		pkg.WriteToFile(fmt.Sprintf("%s/db/migrations/000001_initial.up.sql", serviceName), "", struct{}{})
		pkg.WriteToFile(fmt.Sprintf("%s/db/migrations/000001_initial.down.sql", serviceName), "", struct{}{})
	}

	if enabledIntegrations.Grpc {
		err = os.MkdirAll(fmt.Sprintf("%s/internal/controller/grpc/v1", serviceName), 0777)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(fmt.Sprintf("%s/internal/entity", serviceName), 0777)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(fmt.Sprintf("%s/api", serviceName), 0777)
		if err != nil {
			panic(err)
		}

		pkg.WriteToFile(fmt.Sprintf("%s/internal/controller/grpc/v1/error.go", serviceName), templates.GrpcV1ErrorTemplate, struct{}{})
		pkg.WriteToFile(fmt.Sprintf("%s/internal/entity/error.go", serviceName), templates.InternalErrorTemplate, struct{}{})
	}

	pkg.ReformatFile(fmt.Sprintf("%s/internal/app/config.go", serviceName))
	pkg.ReformatFile(fmt.Sprintf("%s/internal/app/app.go", serviceName))

	pkg.DownloadModules()
}
