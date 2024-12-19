FROM --platform=${BUILDPLATFORM} golang:1.23-bookworm

WORKDIR /app

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

COPY . .
RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build main.go
