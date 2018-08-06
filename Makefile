make build-and-run:
	make build-config && make run

build-config:
	WATCHED_FOLDERS=./example/folder1,./example/folder2,./example/folder3,./example/folder4 go run generator/main.go

run:
	docker-compose -f .docker-compose-gen.yml build && docker-compose -f .docker-compose-gen.yml up

stop:
	docker-compose -f .docker-compose-gen.yml stop

test:
	go test ./file-aggregator/... && go test ./watcher-node/...
