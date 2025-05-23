FROM golang:1.24.3-bullseye AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /api cmd/api/*.go 

FROM scratch AS release

WORKDIR /
COPY --from=build /api /api 
EXPOSE 8080
ENTRYPOINT ["/api"]