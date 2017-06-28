# Reloader

Reloader is a local continuous build and integration tool.

Reloader watches your source code for changes.  When a change is detected, it will rebuild the code and runs either the tests or the project.

Reloader is written with Golang projects in mind, but can be adapted to work for other languages or toolsets by overwriting certain default configurations.

Configuration file example:

``` JSON
{
  "projects": [{
    "root": ".",
    "target": "./bin/bbreloader",
    "build": {
      "pre-build-steps": [],
      "args": "",
      "post-build-steps": [{
        "command": "chmod",
        "args": ["744", "./bin/bbreloader"]
      }, {
        "command": "codesign",
        "args": "-s $CERT ./bin/bbreloader",
      }]
    },
    "test": {
      "args": "",
      "restart-trigger-glob": "*_test.go",
    },
    "run": {
      "args": "",
      "retain": true,
      "rebuild-trigger-glob": "**/*.go, !**/*_test.go",
      "restart-trigger-glob": "*.json",
    },
  }]
}
```

## `projects`
An array of [projects](#Project)

## Project
* `root`
* `target` (optional)
* `build`
* `test`
* `run`

# Examples
## Rebuild and run tests
Command:
`rebuilder test`

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

# Reloader is not...
A substitute for `npm`, `grunt`, `gulp`, or other multi-purpose tools.  Reloader will not install your dependencies, clean up your artifacts, or commit your code to a source code repository.

# What other tools are there?
* `watcher`
* `gradle` / `grunt` / `gulp` / `npm`
* Whatever that tool was in the gopher meetup
