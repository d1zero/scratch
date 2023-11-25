package models

const (
	Postgres = "PostgreSQL"
	Kafka    = "Kafka"
	Grpc     = "gRPC server"
	Http     = "HTTP server"
)

type EnabledIntegrations struct {
	Postgres bool
	Kafka    bool
	Grpc     bool
	Http     bool
}
