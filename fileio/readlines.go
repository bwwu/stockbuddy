package  fileio

import (
  "bufio"
  "io"
  "os"
  "strings"
)

/** Slice of lines contained in a file at the given path */
func ReadLines(path string) ([]string, error) {
  f, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  lines := make([]string, 0)
  reader := bufio.NewReader(f)
  for {
    line, err := reader.ReadString('\n')
    if err == io.EOF {
      break
    }
    if err != nil {
      return nil, err
    }
    lines = append(lines, strings.TrimSpace(line))
  }
  return lines, nil
}
