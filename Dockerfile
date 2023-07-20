# Builder
FROM golang:1.20 as build

WORKDIR /go/src/app

# Speed up build by caching dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o bin/azassigner .


# Final image
FROM gcr.io/distroless/static-debian11@sha256:312a533b1f5584141a7d212ddcc1d079259a84ef68a1a5b0f522017093e3afda

COPY --from=build /go/src/app/bin/azassigner /azassigner
ENTRYPOINT ["/azassigner"]
