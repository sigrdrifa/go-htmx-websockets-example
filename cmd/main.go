package main

import (
	"fmt"

	"github.com/sigrdrifa/go-htmx-websockets-example/internal/hardware"
)

func main() {
  fmt.Println("Starting system monitor..")
  systemSection, err := hardware.GetSystemSection()
  if err != nil {
    fmt.Println(err)
  }

  diskSection, err := hardware.GetDiskSection()
  if err != nil {
    fmt.Println(err)
  }
  cpuSection, err := hardware.GetCpuSection()
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(systemSection)
  fmt.Println(diskSection)
  fmt.Println(cpuSection)
}
