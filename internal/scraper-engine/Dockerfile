FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /scraper-engine

EXPOSE 50051

CMD [ "/scraper-engine" ]