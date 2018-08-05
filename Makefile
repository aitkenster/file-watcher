run:
	docker-compose build && docker-compose up

test:
	go test ./file-aggregator/... && go test ./watcher-node/...
