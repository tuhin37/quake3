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
ARG PYTHON_VERSION=3.9.9
ARG IOQUAKE3_COMMIT="unknown"
LABEL "Maintainer" "Florian Piesche <florian@yellowkeycard.net>"


# install build dependencies and needed tools
RUN apk add \
    wget \
    gcc \
    make \
    zlib-dev \
    libffi-dev \
    openssl-dev \
    musl-dev

# download and extract python sources
RUN cd /opt \
    && wget https://www.python.org/ftp/python/${PYTHON_VERSION}/Python-${PYTHON_VERSION}.tgz \                                              
    && tar xzf Python-${PYTHON_VERSION}.tgz

# build python and remove left-over sources
RUN cd /opt/Python-${PYTHON_VERSION} \ 
    && ./configure --prefix=/usr --enable-optimizations --with-ensurepip=install \
    && make install \
    && rm /opt/Python-${PYTHON_VERSION}.tgz /opt/Python-${PYTHON_VERSION} -rf

ENV IOQUAKE3_COMMIT ${IOQUAKE3_COMMIT}

RUN adduser ioq3srv -D
COPY --chown=ioq3srv --from=builder /usr/local/games/quake3 /usr/local/games/quake3/
ADD --chown=ioq3srv files/ /usr/local/games/quake3/
COPY --chown=ioq3srv config/*.cfg /usr/local/games/quake3/baseq3/
COPY --chown=ioq3srv ioquake3/baseq3/pak*.pk3 /usr/local/games/quake3/baseq3/
COPY --chown=ioq3srv ioquake3/missionpack/pak*.pk3 /usr/local/games/quake3/missionpack/


USER ioq3srv
EXPOSE 27960/udp

RUN pip3 install --upgrade pip

WORKDIR /code

COPY ./requirements.txt /code/requirements.txt

RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt

COPY ./app /code/app

# CMD ["/usr/local/games/quake3/ioq3ded.x86_64", "+exec", "server.cfg", "+exec", "levels.cfg", "+exec", "bots.cfg",  "+seta", "rconPassword", "d405" ]

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "5000"]


# Note: capture these as docker environment variables. Default ram:128, default port 27960 (UDP)
# ram
# port
# password {used in CMD line}
# API access key
