package main

import (
  "log"
)


func main() {
  log.Println("Starting Kron")
  /*
    - parse rcfile
    - setup datastructures
      - queues
      - log locations (ensure writable, etc)
      - web ui related stuff
    - loop
      - watch for changes to krontab
        - reparse and adjust internal state when necessary
        - log an error if the parsing was not successful
      - wait until the next job needs to be kicked off
        - wake up and kick off the job
        - job handles checking when it can start
          - queueing
          - waiting for script dependency
      - listen for signals (sigterm, etc) and shutdown gracefully
      - Web UI
        - update webui state as necessary
        - respond to actions performed from the web ui
  */
  log.Println("Done")
}
