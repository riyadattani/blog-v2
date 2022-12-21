FROM public.ecr.aws/docker/library/golang:1.19 as build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /our-code

# Copy and download dependencies
COPY go.mod .
COPY go.sum .
COPY vendor ./vendor

# Copy the code into the container
COPY cmd ./cmd
COPY src/adapters ./src/adapters
COPY src/domain ./src/domain
COPY src/specifications ./src/specifications

# Build the application and copy somewhere convienient
RUN go build -mod=vendor -o main ./cmd/web/*.go

# create our new image with just the stuff we need
FROM public.ecr.aws/docker/library/alpine
USER nobody
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /our-code/main /usr/bin/main
EXPOSE 8080
CMD ["/usr/bin/main"]
