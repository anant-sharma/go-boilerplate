# Ref: Docker Multi Stage Build
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

# Base Image
FROM golang AS builder

# Install Dependencies
RUN apt-get update
RUN apt-get install -y git make curl unzip

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/anant-sharma/go-boilerplate

# Install protoc
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
RUN unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN make install

RUN make generate

# Download all the dependencies
RUN go get -d -v ./...

# Build Binary
RUN CGO_ENABLED=0 go build -o /go-app

##############################
# STEP 2 build a small image #
##############################
FROM golang:alpine

RUN apk add --no-cache curl

WORKDIR /

# Copy our static executable.
COPY --from=builder /go-app /go-app
COPY --from=builder /go/src/github.com/anant-sharma/go-boilerplate/config/config.yml /config/config.yml
COPY --from=builder /go/src/github.com/anant-sharma/go-boilerplate/protos /protos
COPY --from=builder /go/src/github.com/anant-sharma/go-boilerplate/third_party/OpenAPI /third_party/OpenAPI

# Expose Application Port(s) separated by comma
EXPOSE 8080
EXPOSE 50051

# Health Check
HEALTHCHECK --interval=5s --timeout=3s CMD curl --fail http://localhost:8080/healthz || exit 1

# Run the go-app binary.
ENTRYPOINT ["./go-app"]
