# build backend
FROM golang:1.18-alpine AS backend
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
RUN apk --no-cache --no-progress add --virtual \
  build-deps \
  build-base \
  git

WORKDIR /paopao-ce
COPY . .
ENV GOPROXY=https://goproxy.cn
RUN  make build TAGS='embed jsoniter'

FROM alpine:3.16
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app/paopao-ce
COPY --from=backend /paopao-ce/release/paopao-ce .
COPY configs ./configs

VOLUME ["/app/paopao-ce/configs"]
EXPOSE 8008
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep paopao-ce || exit 1
ENTRYPOINT ["/app/paopao-ce/paopao-ce"]
