# hermes

Never stare at console output again! `hermes` allows you to take a break from staring at your
long-running command line process and will notify you when it completes or fails.

`hermes` currently supports the following notification types:

- [x] Desktop notification (`desktop`)
- [x] Twilio SMS message (`sms`)
- [x] Multi-channel notifications (`multi`)
- [ ] Twilio voice call
- [ ] Slack message
- [ ] Email
- [ ] ...more to come...

### Installation

```bash
go install github.com/jerluc/hermes@latest
```

### Configuring

To configure your Hermes notifications, create a `.hermes.yml` file in your user's home directory,
or in the current directory. Configuration files "closest" to the current working directory will
take precedence. A basic configuration for getting desktop notifications would look like:

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
