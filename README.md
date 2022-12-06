# fontman-client

Fontman is a font management utility for Linux.

## Building

Fontman uses [Task](https://taskfile.dev/) as a build utility / task runner. Once `task` is installed,
you can build `fontman` by simply running: `task build`. This will produce a binary in the `/bin` directory.

For testing, you can use the `task run` command; this will run a linter, run tests, build and then run the executable.

***Note:*** you can run the linter and test steps separately by running `task lint` and `task test`.

## Installation

### From source

To build `fontman`, simply run `task build`. This will produce a binary in `/bin`. Move this binary to wherever binaries are installed on your system. For
instance, on macOS you will run:

```
task build
mv bin/fontman /usr/local/bin  
fontman
```

### From package managers

TBD.

## Usage

### Install from Remote

```
fontman install "Jetbrains Mono"
```

### Install from `fontman.yml`

Run `fontman install` in the same directory as your `fontman.yml` file. This will download & install all fonts
to your system if necessary.

### Install from Local File (`.ttf`, `.otf`)

If you have a font downloaded locally, you can easily install it through `fontman` by passing in the filename
as an argument to `install`.

#### Example
```
fontman install Arial.ttf
```
