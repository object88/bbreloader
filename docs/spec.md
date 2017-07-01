# Complete configuration file spec

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
        "args": "-s", "$CERT", "./bin/bbreloader",
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

## Top level

### `projects`: [[Project](#Project)]

An array of all the projects this configuration is aware of.

## [Project](#project)

### [`root`](#project_root): `string`

* default value: `.`
* required: false

Contains the absolute or relative path to the base directory for this project.

### `target`: `string`

* default value: _no default_
* required: false

Contains the absolute or relative path to the build result, if necessary.  In the case of a library, this should be left blank.  If this is an executable, this value _may_ be left blank; omitting the entry will put the executable in the [`root`](#project_root) location.

### `build`: [`Build`](#build)

### `run`: [`Run`](#run)

### `test`: [`Test`](#test)

## [`Build`](#build)

### `command`: `string`

* default value: `go build`
* required: false

Contains the command used to start the build.

### `args`: `[string]`

* default value: _no default_
* required: false

Contains any additional arguments meant to be passed to the build command.

#### Example

Setting the following parameters...

`"args": ["-tags", "'a b c'"]`

... results in the following command

`go build -tags 'a b c'`

### [`pre-build-steps`](#build_pre): [Step](#step)

* default value: _no default_
* required: false

Contains any steps necessary before the primary compilation.  This may include, for exmaple, any code gen operations, format, linting, or other tasks.

### [`post-build-steps`](#build_post): [Step](#step)

* default value: _no default_
* required: false

Contains any steps necessary after the primary compilation.  This may include, for example, code signing, or other tasks.

## [`Step`](#step)

Describes a [pre-](#build_pre) or [post-](#build_post) build instruction, used from the [build](#build) operation.  Each step is a spawned process, run synchronously.  If the process returns a failure result, the entire build will fail.

### `command`: `string`

* default value: _no default_
* required: true

Contains the command issued to complete this step.

### `args`: `[string]`

* default value: _no default_
* required: false
