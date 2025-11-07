FROM golang:1.25.3-alpine AS base

WORKDIR /app

RUN apk add --no-cache git ca-certificates curl tar

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM base AS dev

# Install Air prebuilt
ARG TARGETARCH
RUN AIR_ARCH=amd64 \
	&& if [ "${TARGETARCH}" = "arm64" ] || [ "${TARGETARCH}" = "aarch64" ]; then AIR_ARCH=arm64; fi \
	&& echo "Installing air for arch=${AIR_ARCH}" \
	&& curl -sL -o /tmp/air.tar.gz "https://github.com/cosmtrek/air/releases/latest/download/air_linux_${AIR_ARCH}.tar.gz" \
	&& tar -C /usr/local/bin -xzf /tmp/air.tar.gz \
	&& rm /tmp/air.tar.gz \
	&& chmod +x /usr/local/bin/air

ENV PATH="/usr/local/bin:${PATH}"

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

## Production image: build the binary and run it
FROM base AS prod

RUN go build -o server ./cmd/main.go

EXPOSE 8080
CMD ["./server"]