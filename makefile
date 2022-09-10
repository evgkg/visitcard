build:
	sudo docker build -t weikelake/visitcard:v2 .
run:
	sudo docker run -p 80:8080 -d --rm --name visitcard weikelake/visitcard:v2
run-dev:
	sudo docker run -p 80:8080 -v certs:/go/visitCard/certs -d --rm --name visitcard weikelake/visitcard:v2
stop:
	sudo docker stop visitcard
push:
	sudo docker push weikelake/visitcard:v2
