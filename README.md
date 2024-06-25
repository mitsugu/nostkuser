nostkuser
========
Get the username of the specified hex npub from the nostk contact list.

### Develop Environment
* Ubuntu 23.04 and later
* Go Language 1.22.4 and later

### Requirements
* [nostk](https://github.com/mitsugu/nostk)

### Setup
#### Install tools
1. Install [git](https://www.git-scm.com/)
2. Install [golang](https://go.dev/)

#### Install nostkuser:
```
go install github.com/mitsugu/nostkuser@<tag name>
```

### Usage
#### Display help documanets
``` bash
nostkuser help

nostkuser -h

nostkuser --help

nostkuser
```

```
nostkuser <hex npub>
```

