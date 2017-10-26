# SLK - minimalistic slack cli

### Benefits
- minimalistic, enjoyable slack interface 
- no need to swich apps
- helps you to save power if you need it
- for those (like myself) who don't like slack to be opened all the time

### Demo

![demo](https://user-images.githubusercontent.com/12980380/31976469-6d593b0c-b940-11e7-90ef-7a0c3fbcd392.png)

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

Read last 10 messages:    

```
$ slk read
```
