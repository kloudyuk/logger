# logger

A simple wrapper around [logrus](https://github.com/sirupsen/logrus) 

## Features

- configured via the standard LOG_LEVEL environment variable
- uses a single logger instance which can be easily included across any file in a project

## Usage

```go
package main

import "github.com/kloudyuk/logger"

var log = logger.Log()

func main() {
  log.Info("Hello world!")
}
```
