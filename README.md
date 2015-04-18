# Golang deadapp

PWS will not log if your app dies before the log system can
bind to the output stream.

This app demonstrates this behavior. If you push it with no
configuration, the app will die and not show the log message
that should appear. By pushing with the `SLEEP` env variable
set, the app will sleep for two seconds, which seems to be
long enough for the log system to bind.

## Broken (No Logs)

```bash
cf push go-deadapp
```

## Working (Shows logs)

```bash
# Assuming you've already pushed the app
cf setenv go-deadapp SLEEP true
cf restage go-deadapp
```
