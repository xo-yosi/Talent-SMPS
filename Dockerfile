FROM golang:alpine AS build

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev linux-headers

COPY . .

RUN --mount=type=cache,target=/go/pkg/,id=gopkgs \
--mount=type=cache,target=/root/.cache,id=gocache \
go telemetry off && go build --ldflags '-linkmode external -extldflags "-static" -s -w' -o go-server

FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app

CMD ["/app/go-server"]

COPY --from=build /app/go-server /app/go-server