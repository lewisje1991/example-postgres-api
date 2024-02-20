FROM golang:1.22.0-bullseye as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api cmd/api/*.go 
RUN ls -l

FROM scratch AS release

WORKDIR /
COPY --from=build /api /api 
EXPOSE 8080
ENTRYPOINT ["/api"]