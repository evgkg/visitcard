build:
	yes | sudo docker builder prune
	sudo docker rmi weikelake/visitcard:v2
	sudo docker build -t weikelake/visitcard:v2 .
run:
	sudo docker run -p 80:8080 -d --rm --name visitcard weikelake/visitcard:v2
run-dev:
	sudo docker run -p 80:8080 -v templates:/go/visitCard/templates -d --rm --name visitcard weikelake/visitcard:v2
stop:
	sudo docker stop visitcard