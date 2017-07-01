# Examples

## Rebuild and run tests

``` JSON
{
  "projects": [{}]
}
```

## Rebuild and restart a server when any `.go` file changes

``` JSON
{
  "projects": [{
    "target": "./bin/myservice",
    "run": {
      "retain": true,
      "rebuild-trigger-glob": "**/*.go",
    },
  }]
}
```

## Restart a JS web service with Babel

``` JSON
{
  "projects": [{
    "run": {
      "command": "babel-cli",
      "args": "./server.js",
      "retain": true,
      "restart-trigger-glob": "**/*.js"
    }
  }]
}
```

