# https://habr.com/ru/companies/otus/articles/660301/
FROM golang:alpine AS builder
LABEL authors='dkhorkov'

WORKDIR /build

COPY . .

RUN go build -o server ./phonebook/server/cmd/main.go

FROM alpine AS runner

WORKDIR /app

COPY .env .

COPY --from=builder /build/server /app/server

CMD ["./server"]
