###################
# Docker commands
up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

clean:
	sudo rm -rf db/data
