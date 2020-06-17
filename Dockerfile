# Ref: Docker Multi Stage Build
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

# Base Image
FROM golang:alpine AS builder

# Install Dependencies
RUN apk update
RUN apk add --no-cache git
RUN apk add make

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/anant-sharma/go-boilerplate

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

RUN make install

RUN make generate

# Build Binary
RUN CGO_ENABLED=0 go build -o /go-app

##############################
# STEP 2 build a small image #
##############################
FROM golang:alpine

WORKDIR /

# Copy our static executable.
COPY --from=builder /go-app /go-app
COPY --from=builder /go/src/github.com/anant-sharma/go-boilerplate/config/config.yml /config/config.yml

# Expose Application Port(s) separated by comma
EXPOSE 8080
EXPOSE 50052

# Run the go-app binary.
ENTRYPOINT ["./go-app"]
