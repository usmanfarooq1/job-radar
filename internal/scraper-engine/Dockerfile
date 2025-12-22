FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /scraper-engine

EXPOSE 50051

CMD [ "/scraper-engine" ]


# FROM golang:1.24 AS builder

# WORKDIR /app

# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# COPY go.mod go.sum ./

# RUN go mod tidy && go mod download

# COPY ./ ./

# RUN go build -trimpath -ldflags="-s -w" -o scraper-engine .

# FROM alpine:latest

# WORKDIR /app

# RUN apk add --no-cache ca-certificates tzdata


# COPY --from=builder /root/.cache /root/.cache 
# COPY --from=builder /app/scraper-engine /app/scraper-engine

# EXPOSE 50051

# CMD [ "/app/scraper-engine" ]