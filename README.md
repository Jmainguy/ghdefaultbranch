# ghdefaultbranch
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/ghdefaultbranch)](https://goreportcard.com/badge/github.com/Jmainguy/ghdefaultbranch)
[![Release](https://img.shields.io/github/release/Jmainguy/ghdefaultbranch.svg?style=flat-square)](https://github.com/Jmainguy/ghdefaultbranch/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/ghdefaultbranch/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/ghdefaultbranch?branch=main)

Rename your Github branch from its current default (likely "master"), to a new default (likely "main")

Can be run against a single repo, or all repos the user owns

## Usage
By default without arguments, this program will print all repos the token has access to.

### Optional ENV variables
* ghdefaultbranchToken: Should be set to a Github API Token with access to the repos you are checking
    * Set permissions for token to repo - full control of private repositories, enable SSO if your repos require it
    * ![Github Personal Access Token Permissions](https://github.com/Jmainguy/ghdefaultbranch/blob/main/docs/permissions.png?raw=true)
    * If not set, program will prompt user to enter a token at run time.

Example configuration in ~/.bashrc
```
export ghdefaultbranchToken=e0e9eac4e84446df6f3db180d07bfb222e91234
```

### Running the progam
```
Usage of ghdefaultbranch:
  -askToken
    	Bool: Force user to enter token instead of trying the env variable ghdefaultbranchToken
  -defaultBranch string
    	Name of the desired defaultBranch, used with renameAll and repository flags (default "main")
  -renameAll
    	Bool: rename default branch in all repos or not
  -repository string
    	Repository to rename branch on, example: Jmainguy/timesheets, not used if renameAll flag is specified
```

### Examples

Set default branch for k8sCapCity to juneteenth
```
[jmainguy@jmainguy ghdefaultbranch]$ ghdefaultbranch -repository Jmainguy/k8sCapCity -defaultBranch "juneteenth"
```

Set default branch for k8sCapCity to main
```
[jmainguy@jmainguy ghdefaultbranch]$ ghdefaultbranch -repository Jmainguy/k8sCapCity
```


Set default branch to main for all repos token has access to

```
[jmainguy@jmainguy ghdefaultbranch]$ ghdefaultbranch -renameAll
```

List all repos token has access to
```
[jmainguy@jmainguy ghdefaultbranch]$ ghdefaultbranch
```
### Sample output
There is no output when changing default branch, unless an error is encountered. 

This is the example output if no arguments passed (list all repos)
```
FullName: Jmainguy/relevy-web, DefaultBranch: main, HTMLURL: https://github.com/Jmainguy/relevy-web
FullName: Jmainguy/repeatafterme, DefaultBranch: main, HTMLURL: https://github.com/Jmainguy/repeatafterme
FullName: Jmainguy/rshell, DefaultBranch: main, HTMLURL: https://github.com/Jmainguy/rshell
```

## Linux / macOS homebrew install

```/bin/bash
brew install jmainguy/tap/ghdefaultbranch
```

## Releases
We currently build releases for RPM, DEB, macOS, and Windows.

Grab Release from [The Releases Page](https://github.com/Jmainguy/ghdefaultbranch/releases)

## Build it yourself
```/bin/bash
export GO111MODULE=on
go build
```
