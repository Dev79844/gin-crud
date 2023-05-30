FROM golang:1.20 AS build
WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build /main /main
EXPOSE 5000
USER nonroot:nonroot
ENV MYSQL_USER=root
ENV MYSQL_PASS=root
ENV MYSQL_DB=recordings
ENTRYPOINT [ "/main" ]
