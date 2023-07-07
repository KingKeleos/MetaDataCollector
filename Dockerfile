FROM golang:latest

# syntax=docker/dockerfile:1#

FROM golang:latest
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /github.com/kingkeleos/ProjectFetcher

EXPOSE 4200

CMD [ "/github.com/kingkeleos/ProjectFetcher" ]