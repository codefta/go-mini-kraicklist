run:
	docker-compose up -d

update:
	docker-compose up -d --no-deps --build app

update-db:
	docker-compose up -d --build db

restart-image:
	docker restart mini_kraicklist_app mini_kraicklist_db