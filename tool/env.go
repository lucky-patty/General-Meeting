package tool

import (
  "bufio"
  // "fmt"
  "os"
  "strings"
)

func LoadEnvFile(path string) error {
  file, err := os.Open(path)
  if err != nil {
    return err  
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := strings.TrimSpace(scanner.Text())

    // Skip line 
    if line == "" || strings.HasPrefix(line, "#") {
      continue
    }

    parts := strings.SplitN(line, "=", 2)
    if len(parts) != 2 {
      continue
    }

    key := strings.TrimSpace(parts[0])
    value := strings.TrimSpace(parts[1])

    value = strings.Trim(value, `"'`)

    os.Setenv(key, value)
  }
  
  return scanner.Err()
}
