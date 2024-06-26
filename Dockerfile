FROM golang:1.22.3-alpine as base

WORKDIR $GOPATH/src/goprom/app/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main cmd/main.go

FROM gcr.io/distroless/static-debian11

# Copy the binary and necessary files for the user
COPY --from=base /main .

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 9000 

CMD ["./main"]
