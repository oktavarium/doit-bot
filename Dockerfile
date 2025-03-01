FROM node:19-alpine AS front_builder
COPY ./web/client/package.json ./web/client/package-lock.json ./
RUN npm install
COPY ./web/client/ ./
RUN npm run build

FROM golang:alpine AS back_builder
LABEL stage=gobuilder
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY go.mod go.sum /
RUN go mod download
COPY cmd/ /cmd
COPY internal/ /internal
COPY web/ /web
COPY --from=front_builder /build/ /web/client/
RUN go build -ldflags="-s -w" -o /bin/server /cmd/main

FROM alpine
WORKDIR /bin
COPY --from=back_builder /bin/server /bin/server
EXPOSE 80
CMD ["/bin/server"]
