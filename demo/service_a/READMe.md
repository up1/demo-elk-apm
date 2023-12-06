# Demo with APM + Go

```
go mod tidy
go fmt

# Initialize using environment variables:
export ELASTIC_APM_SERVICE_NAME=service_a
export ELASTIC_APM_SECRET_TOKEN=L7MDifwx0byJrJPlzl
export ELASTIC_APM_SERVER_URL=https://dd100dedd565427da85293a7ac944157.apm.asia-southeast1.gcp.elastic-cloud.com:443
export ELASTIC_APM_ENVIRONMENT=my-environment

go run main.go
```

Access to APIs
* http://localhost:8080/hello
* http://localhost:8080/service_a
