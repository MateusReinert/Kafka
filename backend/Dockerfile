FROM golang:1.24-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN apk add openssl && rm -rf /var/cache/apk/*

RUN mkdir certs && openssl req -x509 -newkey rsa:4096 -keyout certs/server.key -out certs/server.crt -days 365 -nodes -subj "/CN=localhost"

# Define the command to run your app
CMD ["go", "run", "main.go"]

EXPOSE 4000