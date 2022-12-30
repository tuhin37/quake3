# Build the game in a base container
FROM alpine:3.17.0 AS builder
LABEL "Maintainer" "Florian Piesche <florian@yellowkeycard.net>"

ENV SERVERBIN ioq3ded
ENV BUILD_CLIENT 0

ADD ./ioq3 /ioq3
RUN \
  apk --no-cache add curl g++ gcc make && \
  cd /ioq3 && \
  make && \
  make copyfiles
#  ----------------------------------------------------------------------
# Copy the game files from the builder container to a new image to minimise size
FROM alpine:3.17.0 AS ioq3srv
LABEL "Maintainer" "Florian Piesche <florian@yellowkeycard.net>"


RUN adduser ioq3srv -D
COPY --chown=ioq3srv --from=builder /usr/local/games/quake3 /usr/local/games/quake3/
COPY --chown=ioq3srv config/*.cfg /usr/local/games/quake3/baseq3/
COPY --chown=ioq3srv ioquake3/baseq3/pak*.pk3 /usr/local/games/quake3/baseq3/
COPY --chown=ioq3srv ioquake3/missionpack/pak*.pk3 /usr/local/games/quake3/missionpack/


USER ioq3srv
EXPOSE 27960/udp

ENV PATH="$PATH:/home/ioq3srv/.local/bin"





CMD ["/usr/local/games/quake3/ioq3ded.x86_64", "+exec", "server.cfg", "+exec", "levels.cfg", "+exec", "bots.cfg",  "+seta", "rconPassword", "d405" ]

# CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "5000"]


# Note: capture these as docker environment variables. Default ram:128, default port 27960 (UDP)
# ram
# port
# password {used in CMD line}
# API access key
