FROM registry.semaphoreci.com/golang:latest as builder

ENV APP_HOME /go/src/go-web

WORKDIR "$APP_HOME"
COPY src/ .

RUN go mod download
RUN go mod verify
RUN go build -o go-web

FROM registry.semaphoreci.com/golang:latest

ENV APP_HOME /go/src/go-web
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY src/conf/ conf/
COPY src/views/ views/
COPY --from=builder "$APP_HOME"/go-web $APP_HOME

EXPOSE 8080
CMD ["./go-web"]
