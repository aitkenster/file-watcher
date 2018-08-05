# Watcher Node

Store an updated list of the changes which happen in a specfied folder, and update an aggregator service with the changes

## Endpoints

`GET http://localhost:4000/directory`

Response:
```
{
    "files" [{
        "name: "file.txt"
    },{
        "name": "anotherfile.txt"
    }]
}
```

# To run:

system requirements: Golang

`make build` then run `./watcher-node -dir=./yourwatched/directory`

To test run `make test`
