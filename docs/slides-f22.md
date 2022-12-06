# fontman

**Rensselaer Center of Open Source, Fall 2022**

---

## What is `fontman`?

`fontman` is a command line-based package manager for fonts.

Installing fonts is either done through clunky user interfaces
_or_ manually moving fonts to the correct system folder & regenerating
the font cache.

Package managers like `npm` or `pacman` make it very easy to
install executables & libraries; why is there no package
manager for fonts as well?

---

## Team

- Graham Misail (Project Lead)
- David Kim

---

## Stack

### Client

- Go
- urfave/cli (command line framework)
- FontConfig (font management utility)

### Registry

- Go
- Docker
- Fiber (web framework)
- SQLite (database)
- Railway (deployment server)

---

## Organization

For managing repositories, we went with a **polyrepo** approach, i.e. both the 
client & registry had separate repositories.

- Client: https://github.com/gmisail/fontman-client
- Registry: https://github.com/gmisail/fontman-registry

For handling "tickets" / issues, we used Github's Issue Tracker; having multiple
repositories allowed us to cleanly manage issues for both projects.

Due to being a small team, we used Discord for communication (in addition to 
having system architecture discussions during class.)

---

# Progress

---

## Client

- Boilerplate, setting up subcommands & flags
- Interface between Go <=> FontConfig
- Parser for FontConfig output (`fc-list`)
- Initializing `fontman`-specific files & directories
- Font installation
  - from file (`.ttf`, `.otf`, ...)
  - from remote (i.e. `fontman install Inconsolata`)
  - from project file (`fontman.yml`)

---

## Registry

- Fetching & storing font registry files (id, name, styles)
  - `GET /api/font/<uuid>`, `GET /api/font?name="Menl"`
- Finding fonts with a _similar_ name, i.e. "Menl" ==> "Menlo", "Liga Menlo"
- Uploading font registry files
  - `POST /api/font`

---

# Blockers

---

## Linux is open sourced!

- This allows us to have a great variety in distributions, which is awesome, however...
- It also means that things aren't all standardized.
- For example, the location where fonts are stored is different for each distro. 
- We resolved this by reading distros' documentations and directly reaching out 
to the developers, to confirm details like permissions or folder structure.

---

## MacOS is... not open sourced! 
- We encountered some issues on MacOS that were completely undocumented.
- For example, `/Library/Users/\_name_/Application Support/\_app_/` is no longer writeable since MacOS 12!

```bash
⋊> ~/L/A/fontman: pwd
/Users/meow/Library/Application Support/fontman
⋊> ~/L/A/fontman: ls -la
total 8
drwxr-xr-x   3 meow  staff    96 Nov 21 11:37 ./
drwx------+ 87 meow  staff  2784 Dec  2 13:53 ../
--wxr----x   1 meow  staff    51 Nov 21 18:08 config.yml*
```
- We resolved this by using the Linux standard `~/.config/` to store our config file instead.

---

# Demo

---

## Install from Registry

```bash
fontman install Fira
```

```bash
fontman install "IBM Plex"
```

---

## Install from File

```bash
fontman install ComicMono.ttf
```

---

## Install from Project File

```bash
fontman install
```

---

## Next Steps

- User authentication
  - Who can upload fonts?
- `upload` command
  - Add uploading from a file, instead of a REST endpoint
- Font analytics
  - Number of downloads, recently updated, etc...
- I know it's a CLI app, but...
  - Visual polish, clean up the interface
  - Rewrite the visuals to use **BubbleTea**, a TUI framework
- Web interface
  - Similar to `npmjs.com` or `aur.archlinux.org`
  - Browse & preview fonts
  - Provide a download command
    - `fontman install <UUID> <UUID> …`

---

# Conclusion

Any questions?
