package hello

import "testing"

// func TestHello(t *testing.T) {
//     want := "Hello, world."
//     if got := Hello(); got != want {
//         t.Errorf("Hello() = %q, want %q", got, want)
//     }
// }

func TestRscHello(t *testing.T) {
    want := "Hello, world."
    if got := RscHello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
