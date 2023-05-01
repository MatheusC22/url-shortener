FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /server

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /server /server

EXPOSE 3001

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]