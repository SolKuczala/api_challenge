run:
	docker-compose up

build:
	mkdir ./db-data && chmod 777 ./db-data
	docker-compose build --force-rm --no-cache --parallel
	
mongodebug:
	docker exec -it bet_challenge_oddsdb_1 sh -c "mongo -u root -p example"

clean:
	sudo rm -rf ./db-data