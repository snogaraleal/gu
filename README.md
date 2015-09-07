# GU-CLI

Collection of command line utilities that may come in handy when studying
at the University of Gothenburg.

This software has **no official status** and **no affiliation** or whatsoever
with the University of Gothenburg.


## Installation

* Configure your [golang workspace](https://golang.org/doc/code.html)
* Install with `go get github.com/snogaraleal/gu/gucli`


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

- [x] `gucli sp syllabus -q <query>` Search for syllabuses
- [x] `gucli sp auth` Authenticate via student portal
- [ ] `gucli sp password` Change password


### Ladok

Course registration.

- [ ] `gucli ladok auth` Start session with existing student portal account
- [ ] `gucli ladok info` Show current user information
- [ ] `gucli ladok course` Register for a course
- [ ] `gucli ladok exam` Sign-up for examination
- [ ] `gucli ladok results` Show current user results


### GUL

Platform providing course descriptions, syllabuses and schedules.


### Library

University library.


## License

This project is under the terms of the MIT license. See `LICENSE.txt`.
