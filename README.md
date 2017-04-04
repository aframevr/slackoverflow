# SlackOverflow

> Web hook that posts tagged Stack Overflow questions to Slack, updated using reaction emojis.  

[![GitHub license][license-image]][license-url]
[![Build Status][travis-ci-image]][travis-ci-url]

## Install

```bash
git clone https://github.com/mkungla/slackoverflow.git $GOPATH/src/github.com/aframevr/slackoverflow
cd $GOPATH/src/github.com/aframevr/slackoverflow
make dependencies
make install
```

## Configuration

You can simply execute following command to start interactive configuration  
All slackoverflow configuartion files including SQLite database are stored at
`$HOME/.slackoverflow/`

```bash
slackoverflow init
```

You can check or manually edit generated configuration file.

```bash
slackoverflow config
cat ~/.slackoverflow/slackoverflow.yaml
```

## Run it once

Best way to test it is to use `run` command.

```bash
slackoverflow run
```

## Available commands

> See slackoverflow --help for more info  
> See slackoverflow <command> --help for more info about specific command

```
Usage:
  slackoverflow [OPTIONS] <command>

Application Options:
  -v, --verbose  Be more verbose, This enable loglevel Info
  -d, --debug    Be even more verbose, This enables loglevel Debug
Help Options:
  -h, --help     Show help message
```

**`config`**
> Display SlackOverflow configuration.

**`credits`**
> List of SlackOverflow contributors.

**`reconfigure`**
> Interactive configuration of stackoverflow (aliases: init)

**`restart`**
> Restart SlackOverflow daemon.

**`run`**
> Run SlackOverflow once.

**`slack --help`**
> Slack related commands see slackoverflow slack --help for more info.

**`slack channels`**
> This method returns a list of all Slack channels in the team.

**`slack questions`**
> Post new or update tracked Stack Exchange questions on Slack channel.

| flags | |
| --- | --- |
| `--post-new` | Post new questions origin from configured Stack Exchange Site |
| `--update` | Update information about questions already posted to slack |
| `--all` | Get new questions and update information about existing questions |

**`stackexchange --help`**
> Stack Exchange related commands see slackoverflow stackexchange --help for more info.

**`stackexchange questions`**
> Work with stackexchange questions based on the config

| flags | |
| --- | --- |
| `--get` | Get new questions from configured Stack Exchange Site |
| `--update` | Update information about existing questions |
| `--sync` | Get new questions and update information about existing questions |

**`stackexchange watch`**
> Watch new questions from Stack Exchange site  
> (updated every minute nothing stored to db or posted to slack)

**`start`**
> Start SlackOverflow daemon.

**`status`**
> Get SlackOverflow daemon status.

**`stop`**
> Stop SlackOverflow daemon.

**`validate`**
> Validate stackoverflow configuration

<!-- ASSETS and LINKS -->
<!-- travis-ci -->
[travis-ci-image]: https://travis-ci.org/aframevr/slackoverflow.svg?branch=master
[travis-ci-url]: https://travis-ci.org/aframevr/slackoverflow

<!-- License -->
[license-image]: https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square
[license-url]: https://raw.githubusercontent.com/aframevr/slackoverflow/master/LICENSE
