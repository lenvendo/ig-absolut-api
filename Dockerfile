FROM registry.dev.kubedev.ru/docker/ca-certificates:latest as certs
FROM registry.dev.kubedev.ru/cached/golang:1.16-alpine AS build

#указать имя сервиса
ENV SERVICE_NAME=api
ENV APP=./cmd/app
ENV BIN=/bin/api
ENV PATH_ROJECT=${GOPATH}/src/golang/api
ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOFLAGS=-mod=vendor
ARG VERSION
ENV VERSION ${VERSION:-0.1.0}
ARG BUILD_TIME
ENV BUILD_TIME ${BUILD_TIME:-unknown}
ARG COMMIT
ENV COMMIT ${COMMIT:-unknown}

RUN apk add --update --no-cache curl


WORKDIR ${PATH_ROJECT}
COPY . ${PATH_ROJECT}

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags  -a -o ${BIN} ${APP}

FROM registry.dev.kubedev.ru/cached/alpine:3.12 as geodb

FROM registry.dev.kubedev.ru/cached/alpine:3.12 as production
RUN apk add --update --no-cache ca-certificates

WORKDIR /migrations
COPY ./migrations /migrations

COPY --from=build /bin/${SERVICE_NAME} /bin/${SERVICE_NAME}
COPY --from=geodb /var /var

##В Энтропии указать название сервиса
ENTRYPOINT ["/bin/api"]