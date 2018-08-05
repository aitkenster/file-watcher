run:
	docker-compose build && docker-compose run

test:
	go test ./file-aggregator/... && go test ./watcher-node/...
