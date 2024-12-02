# Comical
Comical is a lite web framework, focus on web serving and web application development.

**This is a PRACTICE project, only for personal uses, DO NOT use it in production**  
This project mainly inspried by the blog [7days-golang - Gee](https://geektutu.com/post/gee.html).

## Getting Started
install the package
```shell
go get github.com/kiritoxkiriko/comical@v0.0.2
```

serve a simple http router
```golang
package main

import (
	"log"
	"fmt"

	"github.com/kiritoxkiriko/comical"
	"github.com/kiritoxkiriko/comical/middleware"
)

func main() {
    r := comical.New()
    // use std logger
    logger := log.Default()
    
    // use build-in middleware
    r.Use(
		middleware.Recovery(logger),
		middleware.Logger(logger),
    )

    // register hello route
    r.GET("/hello", func(c *comical.Context) {
        name := c.Query("name")
        c.String(http.StatusOK, fmt.Sprintf("hello %s, u r at %s\n", name, c.Path))
    })

    // serve at localhost:8080
    if err := r.Run(":8080"); err != nil {
        logger.Fatal("failed to start")
    }
}

```

## Usage
See [examples](https://github.com/kiritoxkiriko/comical/tree/main/example)

## Features
- [x] RESTful API
- [x] Router
- [x] Group Control
- [x] Middleware
- [x] Template
- [x] Logger
- [x] Error Handling
- [x] Static File Server
- [ ] Unit Test
- [ ] Graceful Shutdown
- [x] Panic Recover

## License
MIT License
