package main

import (
  "testing"
)

const (
  SAMPLE_KRONRC_FILEPATH = "samples/kronrc.yaml"
)

func Test_ParseRCFile(t *testing.T) {
  kronrc, err := ParseRCfile(SAMPLE_KRONRC_FILEPATH)
  if err != nil {
    t.Log("Could not parse rcfile")
    t.Fatal(err)
  }
  t.Log(kronrc)
}
