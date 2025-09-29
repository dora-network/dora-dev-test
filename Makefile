# Check for a .env file for environment variables and if one exists, export those
# for use in the Make environment.
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

.DEFAULT_GOAL := compose-up

compose-up:
	@docker compose -f build/env/test-network.yml up --build -d

compose-down:
	@docker compose -f build/env/test-network.yml down
