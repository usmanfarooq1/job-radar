FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /scraper-engine

EXPOSE 50051

CMD [ "/scraper-engine" ]