# Command structure

## General 

### `install`

First look for a `fontman.yml` file in the current working directory 
and install all listed fonts. Otherwise, look up the font in the
font registry and attempt to install.

Install directory is `~/.fontman/` by default.

Subdirectories `~/.fontman/ttf/<fontname>/`, `~/.fontman/otf/<fontname>/`

#### Flags

- `-s, --style`
- `-e, --exclude`
- `-g, --global`

### `uninstall`

Assuming you have the font installed, it will remove it
from the system (if you have permission.)

- `-s, --style`
- `-e, --exclude`
- `-g, --global`

If a font is installed both locally and globally, it will prompt the user for which one to delete.

### `list`

Prints out each font with basic information (which styles
you have installed, etc...)

- `-s, --style`
- `-e, --exclude`
- `-g, --global` 

Recursively search through single user directories first, then system-wide. 

### `info <fontname>`

Reports information about a specified font. If the font
is not installed, it will simply say that the font is not
installed.

### `cache`

Regenerate the font cache for the system.

- `-f, --force` 

### `upload`

Assuming the client is authenticated, upload a `YAML` file
with font metadata.

### `search <fontname>`

Search for a specific font in the font registry and, if 
found, print out details about it.

Optionally include a preview of how to install the 
specified command.

## Authentication

### `authenticate`

