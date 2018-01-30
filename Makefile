up:
	docker-compose up
build:
	docker-compose build --pull circuit

destroy:
	docker-compose stop
	docker-compose rm -f

.PHONY: up build destroy # let's go to reserve rules names
