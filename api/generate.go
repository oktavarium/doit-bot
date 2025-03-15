package main

//go:generate oapi-codegen -generate types -o "../internal/server/ports/httpapi/planner/openapi_types.gen.go" -package "planner" "./planner.yaml"
//go:generate oapi-codegen -generate gin-server -o "../internal/server/ports/httpapi/planner/openapi_api.gen.go" -package "planner" "./planner.yaml"

