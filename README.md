# GU-CLI

GU-CLI is a collection of utilities mostly of which are HTTP APIs that may
come in handy when studying at the University of Gothenburg.

This software has **no official status** and **no affiliation** or whatsoever
with the University of Gothenburg.

## Usage

`gucli <command> <action> <flags ...>`


### SGS

Student housing service.

- [x] `gucli sgs search [-m <market>]` Search vacant objects
- [x] `gucli sgs auth` Exchange user and password for token
- [ ] `gucli sgs info` Show current user information
- [ ] `gucli sgs renew` Renew place in the queue
- [ ] `gucli sgs register -o <object>` Register interest

##### Example

Poll SGS for last minute vacants with an interval of 4 seconds
`watch -n 4 gucli sgs search -m sistam`


### Student Portal

Authentication system for external services.

- [x] `gucli sp auth` Authenticate via student portal
- [ ] `gucli sp password` Change password
- [ ] `gucli sp syllabus -q <query>`


### Ladok

Course registration.


### GUL

Platform providing course descriptions, syllabuses and schedules.


### Library

University library.
