package main

import (
  "testing"
)

func TestAppMain(t *testing.T) {
  expect := "Zoo Application"
  actual := AppName()

  if expect != actual {
    t.Errorf("%s != %s", expect, actual)
  }
}
