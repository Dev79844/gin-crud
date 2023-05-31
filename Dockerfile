FROM golang:1.20 AS build
WORKDIR /app
COPY init.sh ./
RUN sh -c "chmod +x ./init.sh"
CMD [ "./init.sh" ]
COPY go.sum go.mod ./
RUN go mod download
COPY . ./
# COPY init.sh /app/scripts/
RUN CGO_ENABLED=0 GOOS=linux go build -o /main
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build /main /main
EXPOSE 5000
USER nonroot:nonroot
ENTRYPOINT [ "/main" ]
