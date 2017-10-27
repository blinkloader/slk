<p align="left">
  <img style="float: right;" src="https://user-images.githubusercontent.com/12980380/32083822-3122bcc8-bace-11e7-91cb-6a853e128ff9.png" alt="slk logo"/>
</p>


# SLK - minimalistic slack cli
[![Go Report Card](https://goreportcard.com/badge/github.com/yarikbratashchuk/slk)](https://goreportcard.com/report/yarikbratashchuk/slk)

If you like minimalism, like text interfaces and type fast - then you should try SLK. Its a tiny tool for writing and reading slack messages. 

## Benefits

- minimalistic, enjoyable slack interface 
- no need to swich apps
- helps you to save power if you need it
- for those (like myself) who don't like slack to be opened all the time

## Demo

<p align="center">
  <img style="float: right;" src="https://user-images.githubusercontent.com/12980380/32087152-a538ec7c-bae2-11e7-8eef-9158d5c5228a.gif" alt="slk demo"/>
</p>

## Quick start

Install slk:

```
$ go get -u github.com/yarikbratashchuk/slk/...
```

Setup:

```
$ slk setup -t=<slack-token> -c=<channel> -u=<username>
```
- `<slack-token>` - you can generate it [here](https://api.slack.com/custom-integrations/legacy-tokens)
- `<channel>    ` - _channel_ for public and private channels, _@user_ for direct messages.
- `<username>   ` - your name (_yarik_ in demo)

Commands:

```
$ slk listen       #start listening for new messages
$ slk ignore       #stop listening for new messages

$ slk read         #read 10 last messages
$ slk write 'hey!' #write message to channel

$ slk to channel   #switch to public channel or private group
$ slk to @user     #switch to direct messages
$ slk on           #name of the current channel
```

## Proposals / Contributions

Would you like to improve the tool, or have any ideas how to make it better? Feel free to open an [issue](https://github.com/yarikbratashchuk/slk/issues).
