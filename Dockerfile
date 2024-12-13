#build stage
FROM golang:1.23.1-alpine as development

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Run the executable
CMD ["./main"]