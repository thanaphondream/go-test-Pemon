FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/server .

FROM alpine:3.22

WORKDIR /app
COPY --from=build /app/server ./server
COPY --from=build /app/templates ./templates

EXPOSE 3000
CMD ["./server"]