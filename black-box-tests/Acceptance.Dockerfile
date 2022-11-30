FROM public.ecr.aws/docker/library/golang:1.18.2-alpine as deps

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY vendor ./vendor
COPY go.mod go.sum ./

COPY src ./src/
COPY black-box-tests ./black-box-tests/

CMD [ "go", "test", "-count=1", "--tags=acceptance", "./..." ]
