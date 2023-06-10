install-postgres-docker:
	docker pull postgres

remove-postgres:
	docker rm -f biyaheroes-db

run-postgres: install-postgres-docker remove-postgres
	docker run --name biyaheroes-db -e POSTGRES_PASSWORD=123456 -e POSTGRES_USER=biyaheroes \
	-e POSTGRES_DB=biyaheroes \
	-p 5432:5432 -d postgres