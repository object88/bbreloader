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

### Running

The `run` command will start an executable after every successful compilation:

``` bash
% reloader run
```

The `run` command will use the `.reloader.json` file in the current directory to exersize your code.  Alternately, you may use a different file, with the `--config` flag:

``` bash
% reloader run --config .reloader.experiment.json
```

### Testing

The `test` command is coming soon to a theater near you.

### Version

To get the version of the Reloader binary, use the `version` command:

``` bash
% reloader version
```

This command does not require a configuration file.

## Configuration file spec

The complete spec for the configuration file [is available](docs/spec.md).

## Reloader is not

... a substitute for `npm`, `gradle`, `grunt`, `gulp`, or other multi-purpose tools.  Reloader is not meant to install your dependencies, package or clean up your artifacts, or commit your code to a source code repository; your project language's native tools are going to be better suited.

## Alternatives

What other tools are out there?  What inspired this project?  If it's not obvious, this project is heavily influenced by `npm` and `yarn`.

* `watcher`
* `gradle` / `grunt` / `gulp` / `npm` / `yarn`
* Whatever that tool was in the gopher meetup
