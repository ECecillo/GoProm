# syntax=docker/dockerfile:1

FROM golang:1.22.3 as build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -v -o bin/goprom-server ./...


# Use a Docker multi-stage build to create a lean production image.

FROM alpine:latest

WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=build /app/bin/goprom-server .

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["./goprom-server"]