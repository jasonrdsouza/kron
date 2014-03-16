Kron (a cron improvement)
=========================


Motivation
----------
Cron is a very useful piece of software that adheres to the UNIX philosophy of
doing one thing, and doing it well. Unfortunately, for slightly more complex
workflows and job dependencies, utilizing cron introduces more problems than
solutions. In those cases, there are currently two options:

1. Create an ad-hoc, domain specific scheduling solution
2. Employ a vastly more complex scheduling solution and only utilize a fraction
   of its available capability

It seems that there is a niche to be filled by a scheduler that allows slightly
more flexibility than cron in how it gets things done, while generally maintaining
the simplicity and effectiveness of doing one thing (scheduling jobs to be run)
well.

Kron aims to be that scheduler.


Concepts
--------
### Krontab
YAML file of Kron jobs and environment. Reloaded dynamically.
Explanation of Krontab Schema:

### Kronrc
YAML file for Kron configuration. Read at startup.
Explanation of Kronrc fields:

### Dependencies
- Time
- Script


Architecture
------------
Kron Daemon goroutine reads the Krontab, and generates + keeps track of internal
state. It knows when to wake up and kick the next job off.

Kron Job goroutines get kicked off by the Kron daemon, at the appropriate time,
and are responsible for waiting for any script dependencies, and subsequently
getting into their assigned queue. Once they are in the queue, they kick off the
actual kronjob, direct its logs, and keep track of its exit code. They also
are responsible for terminating the kronjob after its timeout, or if the wait
is longer than its wait threshold.

Kron Web goroutine presents an informative view of the kron system.
