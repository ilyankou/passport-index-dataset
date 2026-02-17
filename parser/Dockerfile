FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /parser .

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=build /parser /usr/local/bin/parser
WORKDIR /data
ENTRYPOINT ["parser"]
