.PHONY: test release clean all_tests ft ut mk clean flash startapp

test:
	docker-compose build --pull release
	docker-compose build
	docker-compose run test

release:
	docker-compose up --abort-on-container-exit migrate
	docker-compose run app python3 manage.py collectstatic --no-input
	docker-compose up ${option}

all_tests:
	docker-compose run --rm app pytest

ft:
	docker-compose run --rm app pytest -v -s -l --tb=short acceptance/${filename}

ut:
	docker-compose run --rm app pytest -v -s -l --tb=short ${dir}

flake8:
	docker-compose run --rm app pytest --flake8

mk:
	docker-compose run --rm app python3 manage.py makemigrations ${app_name}
	docker-compose run --rm app python3 manage.py migrate

clean:
	docker-compose down -v
	docker images -q -f dangling=true -f label=application=homeworker | xargs -I ARGS docker rmi -f --no-prune ARGS

flash:
	docker-compose run --rm app python3 manage.py flush --database=default --noinput
	docker-compose run --rm app python3 manage.py createsuperuser

cs:
	docker-compose run --rm app python3 manage.py collectstatic --no-input

startapp:
	docker-compose run --rm app python3 manage.py startapp ${app_name}

superuser:
	docker-compose run --rm app python3 manage.py createsuperuser

docker_clean:
	docker rms '$$(docker ps -a -q)'