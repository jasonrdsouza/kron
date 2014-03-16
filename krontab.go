package main

import (
  "log"
  "io/ioutil"
  "gopkg.in/v1/yaml"
)

type Krontab struct {
  Filepath string
  state yamlKrontab
}

// YAML environment section parsing struct
type yamlEnvironment struct {
  Key string
  Value string
}

// YAML job time dependency subsection parsing struct
type yamlTime struct {
  Month []int
  Monthday []int
  Weekday []string
  Hour []int
  Minute []int
}

// YAML job script dependency subsection parsing struct
type yamlScript struct {
  Command string
  Retry int
}

// YAML job section parsing struct
type yamlJob struct {
  Name string
  Command string
  Queue string
  Time yamlTime
  Script yamlScript
  Timeout int
  Wait int
  Multiple bool
}

type yamlKrontab struct {
  Jobs []yamlJob
  Environment []yamlEnvironment
}


// (re)parses the krontab file
func (k *Krontab) Parse() error {
  yaml_bytes, err := ioutil.ReadFile(k.Filepath)
  if err != nil {
    return err
  }
  if err := yaml.Unmarshal(yaml_bytes, &k.state); err != nil {
    log.Println("Could not parse krontab")
    return err
  }
  return nil
}
