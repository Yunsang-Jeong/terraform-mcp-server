ARG VERSION=dev
ARG PORT=8080

FROM golang:1.24-alpine AS build

RUN apk add --no-cache git ca-certificates

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=0
RUN go build -ldflags "-X terraform-mcp-server/version.Version=${VERSION:-dev}" -o /out/terraform-mcp-server ./main.go

FROM alpine:3.20 AS runtime

RUN addgroup -S app && adduser -S app -G app
USER app

WORKDIR /app
COPY --from=build /out/terraform-mcp-server /app/terraform-mcp-server

EXPOSE $PORT
ENTRYPOINT ["/app/terraform-mcp-server"]