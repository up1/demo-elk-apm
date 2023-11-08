# Demo with APM + Go

```
go mod tidy
go fmt

# Initialize using environment variables:
export ELASTIC_APM_SERVICE_NAME=my-service-name
export ELASTIC_APM_SECRET_TOKEN=
export ELASTIC_APM_SERVER_URL=
export ELASTIC_APM_ENVIRONMENT=my-environment

go run main.go
```