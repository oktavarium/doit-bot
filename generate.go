package main

//go:generate oapi-codegen -generate types -o "./internal/server/ports/httpapi/handlers/openapi_types.gen.go" -package "handlers" "./api/planner.yaml"
//go:generate oapi-codegen -generate gin-server -o "./internal/server/ports/httpapi/handlers/openapi_api.gen.go" -package "handlers" "./api/planner.yaml"

