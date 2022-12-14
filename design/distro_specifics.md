### Tentative approach to finding dir to install to:
iterate through a list of dirs, check if they exist, install there.

For running with root, pass in `global` to `GetInstallPath`, choose a path. 

## Arch Linux
### Single User
- ~/.local/share/fonts/
- ~/.fonts (marked as deprecated, should still search)


### Global
- /usr/local/share/fonts/

/usr/share/fonts/ is managed by `pacman`, should never be modified manually.

https://wiki.archlinux.org/title/fonts#Manual_installation

## Ubuntu
### Single User
- ~/.fonts

### Global
- /usr/share/fonts
- /usr/local/share/fonts

https://wiki.ubuntu.com/Fonts#Manually

## Debian
### Single User
- ~/.local/share/fonts
- ~/.fonts

### Global
- /usr/local/share/fonts

Perm set to 644.

https://wiki.debian.org/Fonts#Manually

## Fedora
### Single User
- ~/.local/share/fonts

### Global
- /usr/local/share/fonts

https://docs.fedoraproject.org/en-US/quick-docs/fonts/#unpackaged

## MacOS
### Single User
- ~/Library/Fonts

### Global
- /Library/Fonts