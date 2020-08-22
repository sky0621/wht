# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.14 as builder

WORKDIR /app

ADD ./view/frontendgen/ ./frontendgen/
COPY ./app-credential.json ./
COPY ./src/ ./
RUN go mod download

WORKDIR cmd

# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
#FROM gcr.io/distroless/base
FROM alpine:latest

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/cmd/server /server
COPY --from=builder /app/frontendgen/ /frontend/
COPY --from=builder /app/app-credential.json /app-credential.json
ENV GOOGLE_APPLICATION_CREDENTIALS /app-credential.json

# Run the web application on container startup.
CMD ["/server"]
