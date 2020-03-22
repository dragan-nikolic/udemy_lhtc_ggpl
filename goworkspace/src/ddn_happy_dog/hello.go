package hello

import "rsc.io/quote"

func Hello() string {
    return "Hello, world."
}

func RscHello() string {
    return quote.Hello()
}
