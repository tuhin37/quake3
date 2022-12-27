build:
	docker build -t drag/q3asrv -f Dockerfile .
build-prod:
	docker build --no-cache -t drag/q3asrv -f Dockerfile .

run: 
	docker run --rm --name q3asrv -p 5000:5000/tcp -p 27960:27960/udp -e RAM=128 -e PORT=27960 -ePASSWORD=UYBEy6AJyHwtz2Z -e API_KEY=70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y -it drag/q3asrv

exec: 
	docker run --rm --name q3asrv -p 27960:27960/udp -it drag/q3asrv sh

delete: 
	./delete-image.sh