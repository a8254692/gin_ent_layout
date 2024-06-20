#GOPATH:=$(shell go env GOPATH)
#VERSION=$(shell git describe --tags --always)

.PHONY: init
# init ENV
init:
	go install entgo.io/ent/cmd/ent@latest


.PHONY: sql2file
# init sql2file
sql2file:
	mysqldump --compact --skip-add-drop-table -d -h 192.168.1.22 -P 3306 -u back -p merchant > ./scripts/sql/merchant.sql;

.PHONY: sql2ent
# init sql2ent
sql2ent:
	sql2ent mysql ddl -src "./scripts/sql/*.sql" -dir "./internal/data/ent/schema"

.PHONY: ent
# init sql2ent
ent:
	ent generate ./internal/data/ent/schema

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
