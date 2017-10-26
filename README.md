# SLK - minimalistic slack cli

### Benefits
- minimalistic, enjoyable slack interface 
- no need to swich apps
- helps you to save power if you need it
- for those (like myself) who don't like slack to be opened all the time

### Demo

![demo](https://user-images.githubusercontent.com/12980380/32066777-96ccc284-ba89-11e7-8823-78fea6d48eb1.png)

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
$ slk listen
```

Write message:

```
$ slk write -m "how are you?"
```

Switch to channel or private chat

```
$ slk to general // for public channel or private group
$ slk to @max // for direct messages
```

Read last 10 messages:    

```
$ slk read
```
