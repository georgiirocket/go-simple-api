#build stage
FROM golang:1.23.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env.production ./.env
ENV MODE=production
RUN go build -o /app/core ./cmd/core/main.go

# Run the executable
CMD ["/app/core"]

FROM alpine:latest AS run

WORKDIR /app

# Copy the application executable from the build image
COPY --from=build /app/core /app/core
COPY --from=build /app/.env /app/.env

CMD ["/app/core"]