# go-logger

Logger for Golang

## Install

```shell
go get github.com/i14t/go-logger
```

## Example

```go
package main

import (
  "github.com/i14t/go-logger"
)

func main() {
  opts := logger.LoggerOptions{
    Level:      "debug",
    Name:       "DEMO",
    HideCaller: false,
  }
  log := logger.NewLogger(opts)
  log.Debug("I am a debug log")
  log.Info("I am a info log")
  log.Warn("I am a warn log")
  log.Error("I am a error log")
}
```

Output

```shell
2024-05-01 23:38:56.256 +07:00 [go] 🟪 DEBUG  [DEMO] [example/main.go:14 main.main] I am a debug log
2024-05-01 23:38:56.256 +07:00 [go] ⬜️ INFO   [DEMO] [example/main.go:15 main.main] I am a info log
2024-05-01 23:38:56.256 +07:00 [go] 🟧 WARN   [DEMO] [example/main.go:16 main.main] I am a warn log
2024-05-01 23:38:56.256 +07:00 [go] 🟥 ERROR  [DEMO] [example/main.go:17 main.main] I am a error log
```

## Reference

- [uber-go/zap][1]
- [dollarsignteam/go-logger][2]

## License

Licensed under the MIT License - see the [LICENSE][3] file for details.

[1]: https://github.com/uber-go/zap
[2]: https://github.com/dollarsignteam/go-logger
[3]: https://github.com/i14t/go-logger/blob/main/LICENSE
