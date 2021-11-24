FROM golang:1.17 AS build


WORKDIR /go/src/mockerize

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o mockerize .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/mockerize/mockerize .

ENV PORT 8080

EXPOSE 8080


CMD ["./mockerize"]