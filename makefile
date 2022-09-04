build:
	sudo docker build -t weikelake/visitcard:vl .
run:
	sudo docker run -p 80:8080 -d --rm --name visitcard weikelake/visitcard:vl
stop:
	sudo docker stop visitcard
volume-ls:
	sudo docker volume ls