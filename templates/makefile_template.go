package templates

import "github.com/d1zero/scratch/internal/models"

func BuildMakefileTemplate(flags models.EnabledIntegrations) string {
	result := `test:
	go test -v ./...

lint:
	golangci-lint run --config=.golangci.yml`

	if flags.Postgres {
		result += `

migrate-new:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations \
	-database ${DB} \
	-verbose up

migrate-down:
	migrate -path db/migrations \
	-database ${DB} \
	-verbose down

migrate-version:
	migrate -path db/migrations \
	-database ${DB} \
	-verbose version`
	}

	if flags.Grpc {
		result += `

protoc:
	export PATH="${PATH}:$(go env GOPATH)/bin" & protoc --go_out=. --go-grpc_out=. api/*.proto`
	}

	return result
}
