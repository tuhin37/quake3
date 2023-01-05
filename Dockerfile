# # STAGE1---------------Build the game in a base container
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
COPY ioquake3/baseq3/pak*.pk3 ./quake3/baseq3/
COPY ioquake3/missionpack/pak*.pk3 ./quake3/missionpack/

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
COPY main.go .
COPY controllers ./controllers
COPY shellScripts ./shellScripts
COPY config ./config
COPY .env ./
RUN go mod tidy
RUN go mod vendor 
RUN go build -mod=vendor 
RUN rm -r controllers/ go.mod go.sum  main.go vendor/


#  STAGE3--------------final
FROM alpine:3.17.0 AS imager
LABEL "Maintainer" "Tuhin Sengupta <tuhin3737@gnail.com>"


RUN apk add --no-cache procps

RUN adduser ioq3srv -D
# copy the quick3 server from the previous stage
COPY --chown=ioq3srv --from=builder /app /app
RUN cp /app/config/* /app/quake3/baseq3
RUN rm -r /app/config
WORKDIR /app

USER ioq3srv
EXPOSE 27960/udp
EXPOSE 5000/tcp

CMD ["/app/Q3AServer"]
# /app/quake3/ioq3ded.x86_64 +exec /app/config/server.cfg +exec /app/config/levels.cfg +exec /app/config/bots.cfg +sta rconPassword d405

# "/usr/local/games/quake3/ioq3ded.x86_64", "+exec", "server.cfg", "+exec", "levels.cfg", "+exec", "bots.cfg",  "+seta", "rconPassword", "d405" 