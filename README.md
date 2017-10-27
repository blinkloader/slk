# SLK - minimalistic slack cli

### Benefits
- minimalistic, enjoyable slack interface 
- no need to swich apps
- helps you to save power if you need it
- for those (like myself) who don't like slack to be opened all the time

### Demo

![demo](https://user-images.githubusercontent.com/12980380/32070344-b4a3ba00-ba94-11e7-88d3-1c6c48d2c4dc.png)

### Quick start

Install slk:

```
$ go get -u github.com/yarikbratashchuk/slk/...
```

Setup slk:

```
$ slk setup -t=<slack-token> -c=<channel> -u=<username>
```
- `<slack-token>` - you can generate it [here](https://api.slack.com/custom-integrations/legacy-tokens) (if you are authorized)
- `<channel>` - channel name in form: _channel_ for public and private channels, _@user_ for direct messages.
- `<username>` - your name (`yarik` in demo)

Wait for messages:

```
$ slk listen   #start listening for new messages
$ slk ignore   #stop listening for new messages
```

Read/Write message:

```
$ slk read     #returns 10 last messages
$ slk write 'hi bot!'         
```

Switch to channel or private chat

```
$ slk to general   #for public channel or private group
$ slk to @max      #for direct messages
```
