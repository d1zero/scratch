package models

const (
	Postgres = "PostgreSQL"
	Redis    = "Redis TODO"
	Kafka    = "Kafka TODO"
	Grpc     = "gRPC server TO_TEST"
	Http     = "HTTP server TODO"
)

type EnabledIntegrations struct {
	Postgres bool
	Kafka    bool
	Grpc     bool
	Http     bool
	Redis    bool
}
