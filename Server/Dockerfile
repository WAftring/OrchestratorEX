FROM golang:1.16-rc-windowsservercore-1809
LABEL maintainer="william.aftring@outlook.com"
ADD . /go/src/srv
WORKDIR /go/src/srv
RUN go get -v srv
RUN go install -v 
ARG API_NAME
ENV API_NAME=$API_NAME
ENTRYPOINT [ "Srv.exe" ]
