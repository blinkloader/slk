# SLK - minimalistic slack chat cli

### Benefits
- minimalistic, enjoyable slack chat interface 
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
$ slk setup -t=<slack-token> -c=<channel-id> -u=<channel-username>
```
- `<slack-token>` - you can generate it [here](https://api.slack.com/custom-integrations/legacy-tokens) (if you are authorized)
- `<channel-id>` - you can get it from direct chat url string (for example: https://blabla.slack.com/messages/channel-id/)
- `<channel-username>` - your name in that chat (`yarik` in demo)

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
