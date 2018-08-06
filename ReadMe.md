# File watcher

Watch and retrieve an aggregated, alphabetized list of files from the folders you specify

<img src="https://uc8fdb16afd46bab6da15120e62a.previews.dropboxusercontent.com/p/thumb/AAKUpXG2NPsGoekNSM0soGRw4K0G6ALhnlGHvETpwOJZtXmtTroCWdeQwJqBihEiCD17Khl-7zLwVHCcpM8ntQH-7ea5a8FVpDE78pxHrxSo7KiNFDThbOFvGWSxtmYVvRV4gKM_PX0Wsu4YU5yKu7_ktHCF4THFk6_7iNHB4tKfTdv4q9bXOCw_CNBBtH9KnRgeQbACwfEZW4efxYsIeR9m5VNdgQYqEjcf5nRaeBPVDA/p.png?size=800x600&size_mode=5" width="300px">

## Endpoint

### Retrieve a list of files

```
GET http://localhost:7807/files
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

## Running it

System requirements: Docker

1. `git clone git@github.com:aitkenster/file-watcher.git`
2. Run `make build-and-run` ( this will create a `file-watcher/example` directory with watched sub directories.
3. `make stop` to stop

To use your own folders:
Change the `WATCHED_FOLDERS` env var in the `build-and-run` command in the Makefile to use your own folders as comma seperated values. Paths can be relative (from the file-watcher folder), or absolute. One watcher node will be generated per folder.

## Tests
Run `make test` to run all tests

## Improvements
- When a watcher node restarts after failure, update all the files watched by that node in the aggregated list
- Check for changes within nested folders
- Integration and API tests
- Pagination in the API
