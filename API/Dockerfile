FROM golang:1.16-rc-windowsservercore-1809
LABEL maintainer="william.aftring@microsoft.com"
ADD . /go/src/api
WORKDIR /go/src/api
RUN go get -v api
RUN go install -v 
ENTRYPOINT [ "API.exe" ]
