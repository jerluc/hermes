# hermes

Never stare at console output again!

### Installation

```bash
go get -d github.com/0xAX/notificator
go install github.com/jerluc/hermes
```

### Configuring

To configure your Hermes notifications, create a `.hermes.yml` file in
your user's home directory, or in the current directory. A basic
configuration for getting desktop notifications would look like:

```yaml
notifier:
  type: desktop
```

For more example configurations, check out the [examples](examples)
directory.

### Usage

```bash
hermes <COMMAND> [ARGS ... ]
```

For example:
```bash
# A successful command
hermes echo 'Hello, world!'

# A failing command
hermes cat /path/does/not/exist
```
