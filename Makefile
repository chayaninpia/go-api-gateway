current_dir=$(shell pwd)
project_name=$(shell basename "${current_dir}" )

build:
	docker build -t $(project_name) .

run:
	docker-compose -f deploy/docker-compose.yaml up -d

stop:
	docker-compose -f deploy/docker-compose.yaml down
	
runkaf:
	docker-compose -f deploy/kafka.yaml up -d

stopkaf:
	docker-compose -f deploy/kafka.yaml down
