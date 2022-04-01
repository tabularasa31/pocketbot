# PocketBot

[![CI](https://github.com/tabularasa31/PocketBot/actions/workflows/main.yml/badge.svg)](https://github.com/tabularasa31/PocketBot/actions/workflows/main.yml)


## Overview

This telegram bot can save the links into your Pocket app (need authorization)

Stack:
- Go
- BoltDB
- Docker

## Quick start

1. Clone the repo to your computer
2. Create `.env` file in the root - it defines all enviroment variables which PocketBot needs. You can find an example at `.env.example`
3. Run `make up` to build and sturt a docker container


## Requirements

- Golang > 1.17
- Docker
- Make
- Telegram bot - create bot via BotFather
- Pocket app - create app https://getpocket.com/developer/apps/new

## TODO
- [X] Add CI\CD
- [ ] Add unit tests
