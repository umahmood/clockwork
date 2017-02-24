/*
Package clockwork sends SMS messages using the Clockwork SMS service.

The below example demonstrates how to send SMS messages:

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

        // print each valid number and its assigned message id
        for num, meta := range resp {
            fmt.Println(num, meta["ID"])
        }
    }
*/
package clockwork
