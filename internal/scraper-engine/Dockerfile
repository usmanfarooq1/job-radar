FROM golang:1.24 as builder

WORKDIR /app

RUN go install github.com/air-verse/air@v1.60.0

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

# COPY ./ ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o /scraper-engine

EXPOSE 50051
CMD ["air"]


# FROM golang:1.24 AS builder
# WORKDIR /app

# COPY go.mod go.sum ./

# RUN go mod tidy && go mod download

# COPY ./ ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o /scraper-engine

# RUN /scraper-engine pinstall

# FROM alpine:latest

# WORKDIR /app

# RUN apk add --no-cache ca-certificates tzdata


# COPY --from=builder /root/ /root/ 
# COPY --from=builder /scraper-engine /app/scraper-engine

# EXPOSE 50051

# CMD [ "/app/scraper-engine", "server" ]