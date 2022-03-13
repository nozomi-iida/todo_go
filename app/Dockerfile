FROM golang:latest
ENV APP_ROOT /$GOPATH/app
WORKDIR ${APP_ROOT}
COPY . .

RUN apt-get update -qq && \
    apt-get install git

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]