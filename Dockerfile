FROM golang:alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /main


FROM alpine
COPY --from=build /main /main

ENV GIN_MODE=release
USER 65534:65534
EXPOSE 8080/tcp

ENTRYPOINT ["/main"]
