doru-up:
	docker-compose -f cmd/dorud/docker-compose-dev.yml --env-file=.env up --build

doru-stop:
	docker-compose -f cmd/dorud/docker-compose-dev.yml stop

doru-clean:
	docker-compose -f cmd/dorud/docker-compose-dev.yml down -v --remove-orphans
