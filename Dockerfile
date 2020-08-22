# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.14 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY ./src/go.* ./
RUN go mod download

# Copy local code to the container image.
COPY ./src/ ./
COPY ./app-credential.json ./
COPY ./view/frontendgen/ ./frontendgen/

WORKDIR cmd

# Build the binary.
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
