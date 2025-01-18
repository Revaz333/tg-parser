serve:
	air

build-image:
	sudo docker build -t tg-parser .
	sudo docker tag tg-parser registry.odva.pro/tg-parser:latest
	sudo docker login registry.odva.pro
	sudo docker push registry.odva.pro/tg-parser:latest