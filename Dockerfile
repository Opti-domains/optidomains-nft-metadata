FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Set the entrypoint command to run the compiled binary
ENTRYPOINT ["./app"]