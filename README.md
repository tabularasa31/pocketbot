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
2. Create .env file in the root - it defines all enviroment variables which PocketBot needs. You can find an example at .env.example
3. Run `make up` to build and sturt a docker container


## Requirements

Golang > 1.13
Docker
Make
Telegram bot

## TODO
- [] Add unit tests
- [X] Add CI\CD
- [] Add supporting many projects at the same time
- [] Add Web UI for configurating
