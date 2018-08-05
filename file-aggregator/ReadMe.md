# File aggregator

Stores and returns a list of file names.

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

### Update information about a file

This is a JSON Patch endpoint

```
PATCH http://localhost:9999/files
```
Request body:
```
[{
    "op": "add",
    "path": "/files",
    "value": {
        "name": "filename.txt"
    }
}]
```

Available ops are "add and "remove". "/files" is the only available path

## Configuration

System requirements: Golang

To run: `make run`
To run tests: `make test`

## Improvements
- Mock server for watcher nodes
