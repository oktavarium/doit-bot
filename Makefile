build: build.front build.back

build.back:
	go build  -o ./bin/server ./cmd/server

build.front:
	npm --prefix ./web/client install
	npm --prefix ./web/client run build

lint:
	go run ./cmd/staticlint ./...
