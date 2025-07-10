FROM --platform=$BUILDPLATFORM golang:1.24.4 AS buildstage


WORKDIR /usr/src/app

COPY go.mod go.sum ./
#go mod and go sum are files that our dependecies are written, similar to requirements.txt or package.json

RUN go mod download
#Downloading our go modules that are written in go.mod

COPY . .


RUN go build -o app ./
#building the application executable binary with name "app" by using source code from current working directory "."


FROM --platform=$BUILDPLATFORM alpine:latest
#using alpine image for security and minimal image size

#omitted the workdir , it is unnecessary in out final alpine image.

COPY --from=buildstage /usr/src/app/app /usr/local/bin/app
#copying the exec binary "app" from buildsatge

EXPOSE 9002
#Exposing port 9002, this is defined by the dev and ops engineers

CMD ["app"]
#As we copied the executable directly into usr/local/bin/app, we can simply run cmd "app" for application start