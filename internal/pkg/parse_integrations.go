package pkg

import "github.com/d1zero/scratch/internal/models"

func ParseIntegrations(integrations []string) models.EnabledIntegrations {
	result := models.EnabledIntegrations{}

	for _, v := range integrations {
		switch v {
		case models.Postgres:
			result.Postgres = true
		case models.Kafka:
			result.Kafka = true
		case models.Grpc:
			result.Grpc = true
		case models.Http:
			result.Http = true
		}
	}

	return result
}
