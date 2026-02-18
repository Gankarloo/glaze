# glaze

Small helpers that extend lipgloss.

## Usage

```go
package main

import (
    "fmt"

    "github.com/charmbracelet/lipgloss"
    "github.com/gankarloo/glaze"
)

func main() {
    style := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("62"))

    boxed := glaze.BorderWithTitle("hello", "title", style, lipgloss.Center)
    fmt.Println(boxed)
}
```
