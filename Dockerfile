FROM golang:1.15.3-alpine
WORKDIR /url
ADD . /url
RUN cd /url && go build -o main
EXPOSE 1323
ENTRYPOINT ./main