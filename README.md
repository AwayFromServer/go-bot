# gobot
[![Go Builder][gh-actions-image]][gh-actions-url] [![Codebeat Badge][codebeat-image]][codebeat-url] [![DockerHub Pulls][dockerhub-pulls-image]][dockerhub-url]

Gobot is a simple Discord bot written in Go that listens for commands in your Discord Server and responds with some (questionably useful) information.
For example, the `!status` command checks if the bot can reach a given URL, defined by the ENV variable TARGET_URL and sends a message back to the issuer.

## Prereqs
- Have a server where you have enough permissions to add integrations and bots, I guess
- Have either Go v`1.22.3` or greater installed, or Docker if using containerized Gobot

## Getting Gobot Running

- Go to [discord.com/developers/applications](https://discord.com/developers/applications)
- Create a new app (this isn't the bot, it's just the config for its integration behavior)
- Under the `Bot` section of settings, un-check the `Public Bot` selector unless specifically desired
- Click the `Reset Token` button to generate a new `Token ID`
- Next, use the `URL Generator` to invite the bot to your server
- - `Scope` = `Bot`
- - `Permissions` += `Receive Message`
- - `Permissions` += `Send Message`
- - `Permissions` += `Send Message in Thread`
- Store the TokenID of the bot! Gobot can't work without it!
- Run gobot with `BOT_TOKEN` set to your TokenID from the previous step
- Set `TARGET_URL` to a website you'd like to poke for status with the `!status` command
- Set `BOT_PREFIX` to pick a prefix for your commands, such as `status`, as seen above

### Required inputs:
```
TARGET_URL = https://some.website:goes/here
BOT_TOKEN = abc123bottokenfordiscordaccesstoyourserver
```
[gh-actions-image]: https://github.com/AwayFromServer/gobot/actions/workflows/build.yml/badge.svg
[gh-actions-url]: https://github.com/AwayFromServer/gobot/actions/workflows/build.yml

[codebeat-image]: https://codebeat.co/badges/c5af66ea-68e5-4b2a-9826-96ddfcbfa513
[codebeat-url]: https://codebeat.co/projects/github-com-awayfromserver-gobot-main

[dockerhub-pulls-image]: https://img.shields.io/docker/pulls/awayfromserver/gobot.svg
[dockerhub-url]: https://hub.docker.com/r/awayfromserver/gobot
