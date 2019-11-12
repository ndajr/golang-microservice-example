# - - -
# Base image
# - - -
FROM golang:1.11-alpine AS base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /opt/app
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download &>/dev/null

# - - -
# Build image
# - - -
FROM base AS builder
COPY . .
ARG GIT_COMMIT
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# - - -
# Runtime image
# - - -
FROM scratch
COPY --from=builder /opt/app/main /
CMD ["/main"]