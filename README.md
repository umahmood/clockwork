# Clockwork

Clockwork is a Go library which allows you to send SMS messages using the [Clockwork SMS service](https://www.clockworksms.com/).

# Installation

Requires Go version 1.7+.

> $ go get github.com/umahmood/clockwork
>
> $ cd $GOPATH/src/github.com/umahmood/clockwork
>
> $ go test ./...

# Usage

Send an SMS message:
```
package main

import (
    "fmt"

    "github.com/umahmood/clockwork"
)

func main() {
    cw := clockwork.New("API-KEY")

    msg := clockwork.SMS{
        To:      clockwork.Numbers{"44123456789", "44987654321"},
        From:    "Gopher",
        Content: "Gophers rule!",
    }

    resp, err := cw.Send(msg)
    if err != nil {
        // ...
    }

    // print each valid number an SMS was sent to and its assigned message id
    for num, meta := range resp {
        fmt.Println(num, meta["ID"])
    }
}
```

# Documentation

> http://godoc.org/github.com/umahmood/clockwork

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
