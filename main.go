package main

import "fmt"
import "slices"
import "maps"
import "cmp"
import "log/slog"

func main() {
  _ = maps.Clone(map[string]int{})
  _ = cmp.Compare(1, 2)
  slog.Info("hello", "count", 3)
  fmt.Printf("%v\n", slices.Max([]int{1, 2, 3}))
}
