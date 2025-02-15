FROM golang:alpine AS builder

WORKDIR /build

# Copy go mod and sum files first (faster builds with caching)
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go binary
RUN go build -o crm.quangdvn-go.com ./cmd/server

# ---- STAGE 2: Run Stage ----
# FROM scratch
FROM alpine:latest

# ✅ Install time zone data
RUN apk --no-cache add tzdata go

COPY ./configs /configs

# Copy the built binary from the build stage
COPY --from=builder /build/crm.quangdvn-go.com /

# # Install SSL certificates (needed for making HTTPS requests in Go)
# RUN apk --no-cache add ca-certificates

# Set working directory in the new container
# WORKDIR /root/
ENV TZ=Asia/Tokyo

EXPOSE 8002

# Run the compiled binary
ENTRYPOINT [ "/crm.quangdvn-go.com", "configs/local.yaml" ]