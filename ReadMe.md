# File watcher

Watch and retrieve an aggregated, alphabetized list of files from the folders you specify

## Endpoint

### Retrieve a list of files

```
GET http://localhost:9999/files
```

Response body:
```
{
    "items" [{
        "name": "afile.txt"
    },{
        "name": "filename.txt"
    }]
}
```

## Configuration

System requirements: Docker

1. `git clone git@github.com:aitkenster/file-watcher.git`
2. Change the .env file to point to the folders you want to watch. Paths can be relative (from the file-watcher folder), or absolute
2. Run `make run`

To add extra file watchers, increase the number of watcher nodes in the docker-compose file. Remember to update the port and WATCHER_ADDRESSES in the file_aggregator service.

## Tests
Run `make tests` to run all tests

## Improvements
- When a watcher node restarts after failure, update all the files watched by that node in the aggregated list
- Check for changes within nested folders
- Integration and API tests
- Pagination in the API
- Generate watcher services with a docker compose template
