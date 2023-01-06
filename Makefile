#  build dokcer image using existing cache (Dev)
build:
	docker build -t fidays/quake3 -f Dockerfile .


#  build docker iamge in --no-cache mode (Production)
build-prod:
	docker build --no-cache -t fidays/quake3 -f Dockerfile .


# run the docker image and start go server in an interactive mode
run: 
	docker run --rm --name quake3 -p 5000:5000/tcp -p 27960:27960/udp --env RAM=128 --env PORT=27960 --env PASSWORD=password --env TOKEN=token -it fidays/quake3


# exec into the existing container
exec:
	./makeScripts/exec.sh



# stop the running container
stop:
	./makeScripts/stop-container.sh


# delete docker image
delete: 
	./makeScripts/delete-image.sh


# run go server
go:
	go run main.go


#  build go
build-go:
	./makeScripts/build-go.sh
