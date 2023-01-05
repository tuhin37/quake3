build:
	docker build -t drag/q3asrv -f Dockerfile .
build-prod:
	docker build --no-cache -t drag/q3asrv -f Dockerfile .

run: 
	docker run --rm --name q3asrv -p 5000:5000/tcp -p 27960:27960/udp -e RAM=128 -e PORT=27960 -ePASSWORD=UYBEy6AJyHwtz2Z -it drag/q3asrv

run-interactive: 
	docker run --rm -p 27960:27960/udp -p 5000:5000/tcp -it drag/q3asrv sh

stop:
	./makeScripts/stop-container.sh
delete: 
	./delete-image.sh

build-go:
	go build -mod=vendor 

go:
	go run main.go
