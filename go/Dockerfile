# Build: docker build -t go/webservice:1.0 .
# Run: docker run -p 8080:8080 go/webservice:1.0

FROM golang:latest AS build
COPY . /src/
WORKDIR /src/
RUN CGO_ENABLED=0 go build -o webservice

#FROM alpine:latest
FROM scratch
WORKDIR /app
COPY --from=build /src/webservice /app/
EXPOSE 8080
ENTRYPOINT ["./webservice"]