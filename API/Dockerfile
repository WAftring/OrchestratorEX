FROM golang:1.16-rc-windowsservercore-1809
LABEL maintainer="william.aftring@microsoft.com"
ADD . /go/src/gosrv
WORKDIR /go/src/gosrv
RUN go get -v gosrv
RUN go install -v 
ENTRYPOINT [ "API.exe" ]
