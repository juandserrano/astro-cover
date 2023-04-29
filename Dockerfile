FROM golang:1.18-alpine
 
WORKDIR /app
 
COPY . ./
RUN rm -rf bin/
RUN rm -rf Dockerfile
RUN rm -rf LICENSE
RUN rm -rf Makefile

RUN go mod download

RUN go build -o /astro-cover
 
CMD ["/astro-cover"]
