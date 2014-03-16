package main

import (
  "log"
  "io/ioutil"
  "gopkg.in/v1/yaml"
)


// YAML queue section parsing struct
type yamlQueue struct {
  Name string
  Slots int
}

type Kronrc struct {
  Queues []yamlQueue
  Logfile string
  JobLogs string
}

// (re)parses the krontab file
func ParseRCfile(filepath string) (*Kronrc, error) {
  yaml_bytes, err := ioutil.ReadFile(filepath)
  if err != nil {
    return nil, err
  }
  var kronrc Kronrc
  if err := yaml.Unmarshal(yaml_bytes, &kronrc); err != nil {
    log.Println("Could not parse kronrc file")
    return nil, err
  }
  return &kronrc, nil
}
