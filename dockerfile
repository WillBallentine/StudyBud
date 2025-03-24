FROM golang:1.23-alpine as build-env
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /studybud studybud/src/cmd/server

FROM alpine:3.14

RUN apk update \
	&& apk upgrade\
	&& apk add --no-cache tzdata curl

#RUN apk --no-cache add bash
ENV TZ America/New_York

WORKDIR /app
COPY --from=build-env /studybud .
COPY --from=build-env /app/src/cmd/server /app/
COPY --from=build-env /app/src/pkg /app/pkg/
COPY --from=build-env /app/src/web /app/web/

EXPOSE 80
EXPOSE 8080
CMD [ "./studybud" ]
