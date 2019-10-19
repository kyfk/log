# Log

[![GoDoc](https://godoc.org/github.com/kyfk/log?status.svg)](https://godoc.org/github.com/kyfk/log)
[![Build Status](https://cloud.drone.io/api/badges/kyfk/log/status.svg)](https://cloud.drone.io/kyfk/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/kyfk/log)](https://goreportcard.com/report/github.com/kyfk/log)
[![codecov](https://codecov.io/gh/kyfk/log/branch/master/graph/badge.svg)](https://codecov.io/gh/kyfk/log)
[![codebeat badge](https://codebeat.co/badges/d45a5e1a-6745-4945-8201-7d9f256fb817)](https://codebeat.co/projects/github-com-kyfk-log-master)

This is the simplest logging for Go.

## Level Management

Use the level package and [MinLevel](https://godoc.org/github.com/kyfk/log#MinLevel)/[SetMinLevel](https://godoc.org/github.com/kyfk/log#SetMinLevel).
```go
// Set minimum level Warn to default logger.
log.SetMinLevel(level.Warn)

log.Debug("debug") // Output nothing
log.Warn("warn") // Output `warn`
```

## Output Format Customization

You can customize the output format easily.
Let's use [Format](https://godoc.org/github.com/kyfk/log#Format)/[SetFormat](https://godoc.org/github.com/kyfk/log#SetFormat) to set the output format into the logger.

```go
SetFormat(format.JSON)
logger.Info("info")
// Output:
// {"level":"INFO","message":"info","time":"2019-10-22T16:50:17.637733482+09:00"}

SetFormat(format.JSONPretty)
logger.Info("info")
// Output:
// {
//   "level": "INFO",
//   "message": "info",
//   "time": "2019-10-22T16:49:00.253014475+09:00"
// }
```

This repository supports only 2 formats that are plain JSON and pretty JSON.

however, you can make a new format that is along `func(map[string]interface{}) string`.
After creating it, just needed to use Format/SetFormat to set it into the logger.

## Common Output Field (Metadata)

If you use some querying service for searching specific logs like BigQuery, CloudWatch Logs Insight, Elasticsearch and other more, [Metadata](https://godoc.org/github.com/kyfk/log#Metadata)/[SetMetadata](https://godoc.org/github.com/kyfk/log#SetMetadata) can be used to set additional pieces of information to be able to search conveniently.
For instance, Service ID, HTTP Request ID, the id of user signed in, EC2 instance-id and other more.

```go
logger := log.New(
    log.Metadata(map[string]interface{}{
        "uesr_id":    "86f32b8b-ec0d-479f-aed1-1070aa54cecf",
        "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
    }),
    log.FlattenMetadata(true),
    log.Format(format.JSONPretty),
)

logger.Info("info")
// Output:
// {
//   "level": "INFO",
//   "message": "info",
//   "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
//   "time": "2019-10-22T17:02:02.748123389+09:00",
//   "uesr_id": "86f32b8b-ec0d-479f-aed1-1070aa54cecf"
// }
```

## Example

```go
package main

import (
    "errors"

    "github.com/kyfk/log"
    "github.com/kyfk/log/format"
    "github.com/kyfk/log/level"
)

func main() {
    logger := log.New(
        log.MinLevel(level.Warn),
        log.Format(format.JSONPretty),
        log.Metadata(map[string]interface{}{
            "service":    "book",
            "uesr_id":    "86f32b8b-ec0d-479f-aed1-1070aa54cecf",
            "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
            "path":       "/operator/hello",
            // more other metadatas
        }),
        log.FlattenMetadata(true),
    )

    logger.Debug("debug")
    logger.Info("info")
    logger.Warn("warn")
    logger.Error(errors.New("error"))
}

// Output:
//
// {
//   "level": "WARN",
//   "message": "warn",
//   "path": "/operator/hello",
//   "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
//   "service": "book",
//   "time": "2019-10-22T16:28:04.180385072+09:00",
//   "trace": [
//     "main.main /home/koya/go/src/github.com/kyfk/log/example/main.go:26",
//     "runtime.main /home/koya/go/src/github.com/golang/go/src/runtime/proc.go:203",
//     "runtime.goexit /home/koya/go/src/github.com/golang/go/src/runtime/asm_amd64.s:1357"
//   ],
//   "uesr_id": "86f32b8b-ec0d-479f-aed1-1070aa54cecf"
// }
// {
//   "error": "*errors.fundamental: error",
//   "level": "ERROR",
//   "path": "/operator/hello",
//   "request_id": "943ad105-7543-11e6-a9ac-65e093327849",
//   "service": "book",
//   "time": "2019-10-22T16:28:04.180589439+09:00",
//   "trace": [
//     "main.main /home/koya/go/src/github.com/kyfk/log/example/main.go:27",
//     "runtime.main /home/koya/go/src/github.com/golang/go/src/runtime/proc.go:203",
//     "runtime.goexit /home/koya/go/src/github.com/golang/go/src/runtime/asm_amd64.s:1357"
//   ],
//   "uesr_id": "86f32b8b-ec0d-479f-aed1-1070aa54cecf"
// }
```
