up:
	DOCKER_BUILDKIT=1 docker compose up --build $(filter-out $@,$(MAKECMDGOALS))

%:
	@:

down:
	docker compose down

db-in:
	docker compose exec db bash -c "psql \"user=postgres password=postgres_pw host=localhost port=5432 dbname=postgres\""

prune:
	docker system prune

gen:
	cd app; \
	go generate ./ent

gomod-update:
	cd app; \
	go get -u && go mod tidy

deploy:
	flyctl deploy --config ./.github/workflows/fly.staging.toml --build-target deploy --remote-only

kushi:
	flyctl proxy 5432:5433 -a topicpost-api-db


setFlyEnv:
	ifeq ($(key),)
	$(error key is not set)
	endif
	ifeq ($(value),)
	$(error value is not set)
	endif
	flyctl -c ./.github/workflows/fly.staging.toml secrets set $(key)=$(value)

unsetFlyEnv:
	ifeq ($(key),)
	$(error key is not set)
	endif
	flyctl -c ./.github/workflows/fly.staging.toml secrets unset $(key)
