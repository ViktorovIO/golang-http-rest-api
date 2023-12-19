FROM golang:1.20.3

RUN mkdir -p /app

WORKDIR /app

COPY . .

ENV GOPATH /
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Build
RUN env GOOS=linux GOARCH=amd64 go build -v ./cmd/server

EXPOSE 8080

# Run
CMD ["/app/server"]
