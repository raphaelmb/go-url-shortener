FROM golang:1.23-alpine3.20 AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

FROM scratch

WORKDIR /app

COPY --from=build /app/api .

EXPOSE 8080

CMD [ "./api" ]

