.env:
	cp .env.example .env
	cp .env.example .env.local

tuna:
# @TODO create dcoker container with tuna that can easily execute ./sh/init-tunnel script
	echo "creating docker container with tuna"

build:
	make .env
	make tuna
