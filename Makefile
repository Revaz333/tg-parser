serve:
	air

build-image:
	sudo docker build -t tg-parser .

tag:
	sudo docker tag tg-parser registry.odva.pro/tg-parser:latest

push:
	sudo docker login registry.odva.pro
	sudo docker push registry.odva.pro/tg-parser:latest