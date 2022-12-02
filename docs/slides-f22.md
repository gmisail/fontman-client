# fontman 

**Rensselaer Center of Open Source, Fall 2022**

---

## What is `fontman`?

`fontman` is a command line-based package manager for fonts.

---

## Stack 

### Client

- Go
- urfave/cli (command line framework)
- FontConfig (font management utility)

### Registry

- Go 
- Fiber (web framework)
- SQLite (database)

---

## Organization

For managing repositories, we went with a **polyrepo** approach, i.e. both the client & registry had
separate repositories.

- Client: https://github.com/gmisail/fontman-client
- Registry: https://github.com/gmisail/fontman-registry

For handling "tickets" / issues, we used Github's Issue Tracker; having multiple repositories allowed us
to cleanly manage issues for both projects.

Due to being a small team, we used Discord for communication (in addition to having system architecture discussions during class.)

---

## Progress

### Client
- Interface between Go <=> FontConfig
- Parser for FontConfig output (`fc-list`)
- Initializing `fontman`-specific files & directories
- Font installation
	- from file (`.ttf`, `.otf`, ...)
	- from remote (i.e. `fontman install Inconsolata`)
	- from project file (`fontman.yml`)

### Registry
- Fetching & storing font registry files (id, name, styles)
	- `GET /api/font/<uuid>`, `GET /api/font?name="Menl"`
- Finding fonts with a *similar* name, i.e. "Menl" ==> "Menlo", "Liga Menlo"
- Uploading font registry files
	- `POST /api/font`

---

## Client

TBD.

---

## Registry 

TBD.

---

## Demo
```bash

```
