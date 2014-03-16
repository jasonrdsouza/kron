package main

import (
  "testing"
)

const (
  SAMPLE_KRONTAB_FILEPATH = "samples/krontab.yaml"
)

func Test_ParseYAML(t *testing.T) {
  k := Krontab{
        SAMPLE_KRONTAB_FILEPATH,
        *new(yamlKrontab),
       }
  if err := k.Parse(); err != nil {
    t.Fatal(err)
  }
  t.Log(k.state.Jobs[0].Time)
}
