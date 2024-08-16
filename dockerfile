FROM golang:1.22.4

WORKDIR /app

COPY ../../Desktop/final-project .

RUN go mod download

EXPOSE 7540

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /vabulan ./cmd/TODO-LIST/main.go

CMD ["/vabulan"]