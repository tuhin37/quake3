#  build dokcer image using existing cache (Dev)
build:
	docker build -t drag/q3asrv -f Dockerfile .


#  build docker iamge in --no-cache mode (Production)
build-prod:
	docker build --no-cache -t drag/q3asrv -f Dockerfile .


# run the docker image and start go server in an interactive mode
run: 
	docker run --rm --name q3asrv -p 5000:5000/tcp -p 27960:27960/udp -e RAM=128 -e PORT=27960 -ePASSWORD=UYBEy6AJyHwtz2Z -it drag/q3asrv


# exec into the existing container
exec:
	./makeScripts/exec.sh


# create and exec into a new container. But it DOES NOT start the go server (for debugging)
exec-new: 
	docker run --rm -p 27960:27960/udp -p 5000:5000/tcp -it drag/q3asrv sh


# stop the running container
stop:
	./makeScripts/stop-container.sh


# delete docker image
delete: 
	./makeScripts/delete-image.sh


# run go server
run-go:
	go run main.go
