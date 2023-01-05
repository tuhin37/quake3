# # STAGE1---------------Build the dedicated q3a server in this stage
FROM golang:1.18-alpine as builder

# build q3a dedicated server
ENV SERVERBIN ioq3ded
ENV BUILD_CLIENT 0

ADD ./ioq3 /ioq3
RUN \
  apk --no-cache add curl g++ gcc make && \
  cd /ioq3 && \
  make && \
  make copyfiles

# copy the build q3a dedicated server to /app
WORKDIR /app
RUN mkdir -p /app/quake3
RUN cp -r /usr/local/games/quake3/* ./quake3


# Build the go server
COPY go.mod go.sum main.go ./
COPY controllers ./controllers
RUN go mod tidy
RUN go mod vendor
RUN go build -mod=vendor
RUN rm -r controllers/ go.mod go.sum  main.go vendor/



#  STAGE2-------------Build final Image-------------
FROM alpine:3.17.0 AS imager
LABEL "Maintainer" "Tuhin Sengupta <tuhin3737@gnail.com>"

# Needed by IsRunning.sh script
RUN apk add --no-cache procps

# create a non root user
RUN adduser ioq3srv -D
WORKDIR /app

# copy the quick3 server and go server from the previous stage
COPY --chown=ioq3srv --from=builder /app /app

# copy pk3, cfg, sh files to the container from local
COPY --chown=ioq3srv ioquake3/baseq3/pak*.pk3 /app/quake3/baseq3/
COPY --chown=ioq3srv ioquake3/missionpack/pak*.pk3 /app/quake3/missionpack/
COPY --chown=ioq3srv config /app/quake3/baseq3
COPY --chown=ioq3srv shellScripts /app/shellScripts

# switch user
USER ioq3srv
EXPOSE 27960/udp
EXPOSE 5000/tcp

# Run go server
CMD ["/app/Q3AServer"]