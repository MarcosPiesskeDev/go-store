FROM golang

WORKDIR /app/src/go-store-back

ENV GOPATH=/app

COPY . /app/src/go-store-back

RUN go build cmd/server/main.go

ENTRYPOINT ["./main"]

EXPOSE ${API_PORT}
