
# FROM golang:1.16.12-alpine3.15
FROM golang:1.18.3-alpine3.15

LABEL maintainer="CB"
LABEL version=2.0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the code
COPY ./code ./

# Download all dependencies.
RUN go get -d -v ./...
RUN go install -v ./...

# Build the application
RUN go build -o main .

# Expose port 9000
EXPOSE 9000

# Command to run the executable
CMD ["./main"]