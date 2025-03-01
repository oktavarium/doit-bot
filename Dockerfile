FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY go.mod go.sum /
RUN go mod download
COPY cmd/ /cmd
COPY internal/ /internal
RUN go build -ldflags="-s -w" -o /bin/server /cmd/main
FROM alpine
WORKDIR /bin
COPY --from=builder /bin/server /bin/server
CMD ["/bin/server"]
