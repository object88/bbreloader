# Reloader

Reloader is a local continuous build and integration tool.

Reloader watches your source code for changes.  When a change is detected, it will rebuild the code and run either the tests or the project.

Reloader is written with Golang projects in mind, but can be adapted to work for other languages or toolsets by overwriting certain default configurations.

## Commands

Reloader has five commands that will handle setup and operation.

### Command line help

If Reloader is run with no arguments, the help information is written to the console.  Alternately, help may be explicitely requested with the `help` command:

``` bash
% reloader help
```

There is no `man` page, because this isn't 1987.  This command does not require a configuration file.

### Starting a new configuration

Using the `init` command, Reloader will generate a new, blank configuration file for you:

``` bash
% reloader init
```

This will place a new `.reloader.json` file in the current directory.  If such a file already exists, it will _not_ be overwritten.

Alternately, you can specify a different file to create with the `--config` flag:

``` bash
% reloader init --config .reloader.experiment.json
```

Once your configuration file is created, you may customize it with your favorite JSON editor.  The JSON structure is defined in [Configuration file structure](docs/configuration.md).

### Running and testing

Reloader has to main modes of operation: running and testing.  Each has thier own command.

#### Run

The `run` command will start an executable after every successful compilation:

``` bash
% reloader run
```

### Version

To get the version of the Reloader binary, use the `version` command:

``` bash
% reloader version
```

This command does not require a configuration file.

## Example files

Configuration file example:

``` JSON
{
  "projects": [{
    "root": ".",
    "target": "./bin/bbreloader",
    "build": {
      "pre-build-steps": [],
      "command": "",
      "args": "",
      "post-build-steps": [{
        "command": "chmod",
        "args": ["744", "./bin/bbreloader"]
      }, {
        "command": "codesign",
        "args": "-s $CERT ./bin/bbreloader",
      }]
    },
    "run": {
      "command": "",
      "args": "",
      "retain": true,
      "rebuild-trigger-glob": "**/*.go, !**/*_test.go",
      "restart-trigger-glob": "*.json",
    },
    "test": {
      "command": "",
      "args": "",
      "restart-trigger-glob": "*_test.go",
    }
  }]
}
```

## Examples

### Rebuild and run tests

Command:
`rebuilder test`

``` JSON
{
  "projects": [{}]
}
```

### Rebuild and restart a server when any `.go` file changes

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

### Restart a JS web service with Babel

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

## Reloader is not

... a substitute for `npm`, `gradle`, `grunt`, `gulp`, or other multi-purpose tools.  Reloader will not install your dependencies, clean up your artifacts, or commit your code to a source code repository.

## Alternatives

What other tools are out there?  What inspired this project?  If it's not obvious, this project is heavily influenced by `npm` and `yarn`.

* `watcher`
* `gradle` / `grunt` / `gulp` / `npm` / `yarn`
* Whatever that tool was in the gopher meetup
