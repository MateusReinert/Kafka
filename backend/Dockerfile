FROM golang:1.24-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

# Define the command to run your app
CMD ["go", "run", "main.go"]

EXPOSE 4000