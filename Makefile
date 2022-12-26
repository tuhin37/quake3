build:
	docker build -t drag/q3asrv -f Dockerfile .
build-prod:
	docker build --no-cache -t drag/q3asrv -f Dockerfile .

run: 
	docker run --rm --name q3asrv -p 27960:27960/udp -it drag/q3asrv

exec: 
	docker run --rm --name q3asrv -p 27960:27960/udp -it drag/q3asrv sh

delete: 
	./delete-image.sh