FROM golang:1.17-alpine AS build

RUN apk add --no-cache git

WORKDIR /app/social-calendar

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /out/social-calendar-api ./cmd/social-calendar-api

FROM alpine

WORKDIR /

COPY --from=build /out/social-calendar-api /social-calendar-api 
EXPOSE 8080

CMD ["/social-calendar-api"]