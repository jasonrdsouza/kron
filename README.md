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


Usage
-----


Todo
----
- Implement QueueMap
  - responsible for managing queue creation/ deletion
  - remove the concept of a name from the queue struct
  
