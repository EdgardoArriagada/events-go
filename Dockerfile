FROM golang:1.22-alpine

# Install SQLite development headers and other necessary packages
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set CGO_ENABLED=1 for Go build
ENV CGO_ENABLED=1

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download  

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
