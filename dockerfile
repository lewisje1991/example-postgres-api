FROM golang:1.23.0-bullseye as build

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