FROM golang:1.20.2-alpine
MAINTAINER ChungTM
ENV GIN_MODE=release
ARG APP=RolePermissionApplication
WORKDIR /app

COPY go.sum go.mod ./
# Download all dependencies
RUN go mod download
# Copy the source from the current directory to the workspace
COPY . .
RUN mkdir build
RUN go build -o build/${APP} src/main.go
RUN mv build/${APP} ${APP}
#RUN export PATH=$PATH:/usr/local/go/bin
EXPOSE 8100
CMD ["./RolePermissionApplication"]
