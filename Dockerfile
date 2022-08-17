FROM golang:1.18-alpine

# Creat New Directory inside the container
RUN mkdir /app

ADD . /app

# Set the Current Working Directory inside the container

WORKDIR /app

# Build the Go app
RUN go build -o movie .

# This container exposes port 8080 to the outside world
EXPOSE 4567

# Run the binary`

ENTRYPOINT ["./movie","-container"]
