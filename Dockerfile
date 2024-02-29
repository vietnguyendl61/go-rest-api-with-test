# Start from the official golang image
FROM golang:latest as build

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy the source code from the current directory to the Working Directory inside the container
COPY . /go/src/app

# Build
RUN cd /go/src/app && CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o /user-service .

# Expose port to the outside world
EXPOSE 8001

# Command to run the executable
CMD /user-service

FROM build as test