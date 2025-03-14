build: build.front build.back

build.back:
	go build  -o ./bin/server ./cmd/server

build.front:
	npm --prefix ./web/client install
	npm --prefix ./web/client run build

lint:
	go run ./cmd/staticlint ./...

gen:
	go generate ./...

dc-up:
	docker compose up

dc-down:
	docker compose down

dc-rebuild:
	docker compose build

dc-up-mongo:
	docker compose up mongo
