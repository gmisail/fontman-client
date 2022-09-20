# Command structure

## General 

### `install`

First look for a `fontman.yml` file in the current working directory 
and install all listed fonts. Otherwise, look up the font in the
font registry and attempt to install.

### `uninstall`

Assuming you have the font installed, it will remove it
from the system (if you have permission.)

### `list`

Prints out each font with basic information (which styles
you have installed, etc...)

### `info <fontname>`

Reports information about a specified font. If the font
is not installed, it will simply say that the font is not
installed.

### `cache`

Regenerate the font cache for the system.

### `upload`

Assuming the client is authenticated, upload a `YAML` file
with font metadata.

### `search`

Search for a specific font in the font registry and, if 
found, print out details about it.

Optionally include a preview of how to install the 
specified command.

## Authentication

### `login`

TBD.

### `logout`

TBD.
