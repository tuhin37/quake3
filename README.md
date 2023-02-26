# Quake3 Server

## 1. Docker Image

build

```shell
make build
```

or 

```shell
make build-prod
```

### 2. Run docker image

```shell
make run
```

or

```shell
docker run --rm --name quake3 -p 5000:5000/tcp -p 27960:27960/udp --env RAM=128 --env PORT=27960 --env PASSWORD=password --env TOKEN=70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y -it fidays/quake3
```

or

```shell
docker-compose -f docker-compose/docker-compose.quake3.yaml up
```



This uses the following default parameters

     name: `quake3`

    api-port: `5000/TCP`

    game-port: `27960/TCP`

    ram-requested: `128MB`

    console-password: `password`

    bearer-token: `70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y`

___

## 3. APIs

### 3.1 /status

request

```sh
curl --location --request GET 'localhost:5000/status' \
--header 'Authorization: Bearer 70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y'
```

response

```json
{
    "autoexec": {
        "com_hunkmegs": "128",
        "dedicated": "1",
        "net_port": "27960",
        "vm_cgame": "2",
        "vm_game": "2",
        "vm_ui": "2"
    },
    "bots": {
        "bot_enable": "1",
        "bot_minplayers": "5",
        "bot_nochat": "1",
        "g_spskill": "4"
    },
    "map": "q3dm17",
    "server": {
        "capturelimit": "8",
        "cl_maxpackets": "40",
        "cl_packetdup": "1",
        "fraglimit": "0",
        "g_forcerespawn": "0",
        "g_friendlyFire": "1",
        "g_gametype": "CTF",
        "g_inactivity": "120",
        "g_log": "server.log",
        "g_motd": "Welcome",
        "g_quadfactor": "4",
        "g_teamAutoJoin": "0",
        "g_teamForceBalance": "0",
        "g_weaponrespawn": "2",
        "logfile": "3",
        "rate": "12400",
        "rconpassword": "password",
        "snaps": "40",
        "sv_hostname": "NIGGIS",
        "sv_maxclients": "16",
        "sv_pure": "1",
        "timelimit": "30"
    },
    "status": "stopped"
}
```

### 3.2 /start

request

```shell
curl --location --request GET 'localhost:5000/start'
```

response

```json
{
    "message": "server started"
}
```

or

```json
{
    "message": "server already running"
}
```

### 3.3 /stop

response

```json
{
    "message": "server stopped"
}
```

or

```json
{
    "message": "server was not running"
}
```

### 3.4 /update

```sh
curl --location --request PUT 'http://localhost:5000/update' \
--header 'Authorization: Bearer 70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "drag",
    "bots": {
        "bot_enable": 0,
        "bot_minplayers": 0,
        "bot_nochat": 0,
        "g_spskill": 2
    },
    "autoexec": {
        "com_hunkmegs": 512,
        "dedicated": 1,
        "net_port": 27960,
        "vm_cgame": 2,
        "vm_game": 2,
        "vm_ui": 2
    },
    "map": "q3dm17",
    "server": {
        "capturelimit": 8,
        "cl_maxpackets": 40,
        "cl_packetdup": 1,
        "fraglimit": 0,
        "g_forcerespawn": 0,
        "g_friendlyFire": 1,
        "g_gametype": "TD",
        "g_inactivity": 120,
        "g_log": "server.log",
        "g_motd": "joey",
        "g_quadfactor": 3,
        "g_teamAutoJoin": 0,
        "g_teamForceBalance": 0,
        "g_weaponrespawn": 2,
        "logfile": 3,
        "rate": 12401,
        "rconpassword": "secret123",
        "snaps": 41,
        "sv_hostname": "yoyo",
        "sv_maxclients": 15,
        "sv_pure": 1,
        "timelimit": 33
    },
    "restart": "true" 
}'
```

response

```json
{
    "autoexec": {
        "vm_ui": 2,
        "vm_game": 2,
        "vm_cgame": 2,
        "net_port": 27960,
        "dedicated": 1,
        "com_hunkmegs": 512
    },
    "bots": {
        "bot_enable": 0,
        "bot_minplayers": 0,
        "bot_nochat": 0,
        "g_spskill": 2
    },
    "map": "q3dm17",
    "server": {
        "capturelimit": 8,
        "cl_maxpackets": 40,
        "cl_packetdup": 1,
        "fraglimit": 0,
        "g_forcerespawn": 0,
        "g_friendlyFire": 1,
        "g_gametype": "TD",
        "g_inactivity": 120,
        "g_log": "server.log",
        "g_motd": "joey",
        "g_quadfactor": 3,
        "g_teamAutoJoin": 0,
        "g_teamForceBalance": 0,
        "g_weaponrespawn": 2,
        "logfile": 3,
        "rate": 12401,
        "rconpassword": "secret123",
        "snaps": 41,
        "sv_hostname": "yoyo",
        "sv_maxclients": 15,
        "sv_pure": 1,
        "timelimit": 33
    },
    "restart": "true"
}
```

example for longest yeard the request will be

```sh
curl --location --request PUT 'http://localhost:5000/update' \
--header 'Authorization: Bearer 70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "drag",
    "bots": {
        "bot_enable": 0,
        "bot_minplayers": 0,
        "bot_nochat": 0,
        "g_spskill": 2
    },
    "autoexec": {
        "com_hunkmegs": 512,
        "dedicated": 1,
        "net_port": 27960,
        "vm_cgame": 2,
        "vm_game": 2,
        "vm_ui": 2
    },
    "map": "q3dm17",
    "server": {
        "capturelimit": 8,
        "cl_maxpackets": 40,
        "cl_packetdup": 1,
        "fraglimit": 20,
        "g_forcerespawn": 0,
        "g_friendlyFire": 1,
        "g_gametype": "FFA",
        "g_inactivity": 120,
        "g_log": "server.log",
        "g_motd": "joey",
        "g_quadfactor": 3,
        "g_teamAutoJoin": 0,
        "g_teamForceBalance": 0,
        "g_weaponrespawn": 2,
        "logfile": 3,
        "rate": 12401,
        "rconpassword": "secret123",
        "snaps": 41,
        "sv_hostname": "yoyo",
        "sv_maxclients": 4,
        "sv_pure": 0,
        "timelimit": 30
    },
    "restart": "true" 
}'
```

response

```json
{
    "autoexec": {
        "vm_ui": 2,
        "vm_game": 2,
        "vm_cgame": 2,
        "net_port": 27960,
        "dedicated": 1,
        "com_hunkmegs": 512
    },
    "bots": {
        "bot_enable": 0,
        "bot_minplayers": 0,
        "bot_nochat": 0,
        "g_spskill": 2
    },
    "map": "q3dm17",
    "server": {
        "capturelimit": 8,
        "cl_maxpackets": 40,
        "cl_packetdup": 1,
        "fraglimit": 20,
        "g_forcerespawn": 0,
        "g_friendlyFire": 1,
        "g_gametype": "FFA",
        "g_inactivity": 120,
        "g_log": "server.log",
        "g_motd": "joey",
        "g_quadfactor": 3,
        "g_teamAutoJoin": 0,
        "g_teamForceBalance": 0,
        "g_weaponrespawn": 2,
        "logfile": 3,
        "rate": 12401,
        "rconpassword": "secret123",
        "snaps": 41,
        "sv_hostname": "yoyo",
        "sv_maxclients": 4,
        "sv_pure": 0,
        "timelimit": 30
    },
    "restart": "true"
}
```

### 3.5 /restore

This method restores the default configs

request

```sh
curl --location --request PUT 'http://localhost:5000/restore' \
--header 'Authorization: Bearer 70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y' \
--data-raw ''
```

response

```json
{
    "message": "config restored"
}
```

___

## 4. Make commands

### 4.1 exec

If a dontainer is running then this command will find that container and exec into that. This is used for debugging perpouses. Otherwise, if no quake3 container was running to begin with, then this command will run a new container in interactive mode.

```shell
make exec
```

### 4.2 stop

This command will stop a running quake3 container. If no containers were running then this will not output any error

```sh
make stop
```

### 4.3 push

This command will push the locally built docker image to dockerhub under fidats repository

```sh
make push
```

### 4.4 delete

This command finds out the image ID for quake3 and deletes that image from local filesystem. 

```sh
make delete
```

### 4.5 go

This commands run the go server locally, Here the go server acts as a wrapper on quake3 server. However, the quake3 server is built during the docker build process. Which means the quake3 server is not available in the localmachine. For this reason the q3a server is emulated by a bash script that counts number, in the local machine. 

This server can accept all the api(s). Note that the server runs on port 5000

```shell
make go
```

response

```shell
go run main.go
server starts...
.env file not found
SH | configs restored
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /start                    --> github.com/tuhin37/quake3/controllers.StartServer (3 handlers)
[GIN-debug] GET    /status                   --> github.com/tuhin37/quake3/controllers.GetStatus (3 handlers)
[GIN-debug] PUT    /update                   --> github.com/tuhin37/quake3/controllers.UpdateGame (3 handlers)
[GIN-debug] PUT    /restore                  --> github.com/tuhin37/quake3/controllers.RestoreDefault (3 handlers)
[GIN-debug] GET    /stop                     --> github.com/tuhin37/quake3/controllers.StopServer (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :5000
```

### 4.6 build-go

This command compiles the go code into a single binary. The final binary is generated in to root of the project directory and it will be named `quake3`. 

build

```shell
make build-go
```

run

```sh
./quake3
```

response

```shell
server starts...
.env file not found
SH | configs restored
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /start                    --> github.com/tuhin37/quake3/controllers.StartServer (3 handlers)
[GIN-debug] GET    /status                   --> github.com/tuhin37/quake3/controllers.GetStatus (3 handlers)
[GIN-debug] PUT    /update                   --> github.com/tuhin37/quake3/controllers.UpdateGame (3 handlers)
[GIN-debug] PUT    /restore                  --> github.com/tuhin37/quake3/controllers.RestoreDefault (3 handlers)
[GIN-debug] GET    /stop                     --> github.com/tuhin37/quake3/controllers.StopServer (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :5000
```

___

## 6. Deployment

### 6.1 docker compose

```yaml
version: "3"
services:
  quake3:
    image: fidays/quake3:latest
    container_name: quake3
    restart: always
    ports:
      - 5000:5000/tcp
      - 27960:27960/udp
    environment:
      - RAM=128
      - PORT=27960
      - PASSWORD=password
      - TOKEN=70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y
```

```shell
docker-compose -f docker-compose.quake3.yaml up
```

### 6.2 Kubernates manifest

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: quake3
  labels:
    name: quake3
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: quake3
  namespace: quake3
spec:
  selector:
    matchLabels:
      app: quake3
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: quake3
    spec:
      containers:
      - name: quake3
        image: fidays/quake3:latest
        env:
          - name: RAM
            value: "128"
          - name: PORT
            value: "27960"
          - name: PASSWORD
            value: "password"
          - name: TOKEN
            value: "70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y"
        ports:
        - containerPort: 5000
          name: api
          protocol: TCP
        - containerPort: 27960
          name: game
          protocol: UDP
      nodeSelector:
        kubernetes.io/hostname: w3-vm
---
apiVersion: v1
kind: Service
metadata:
  name: quake3-lb
  namespace: quake3
spec:
  ports:
  - name: api
    port: 80
    protocol: TCP
    targetPort: 5000
  - name: game
    port: 27960
    protocol: UDP
    targetPort: 27960
  selector:
    app: quake3
  type: LoadBalancer
```

```sh
kubectl apply -f quake3.yaml
```

or

```sh
kubectl apply -f https://raw.githubusercontent.com/tuhin37/quake3/drag/quake3.yaml?token=GHSAT0AAAAAABZ67GQRNEGMDXE4JQVNCZTCY63RJTQ
```
