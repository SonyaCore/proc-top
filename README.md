# ProcTop

![go]
![goversion]
![version]
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

### Simple cli system monitor written in Go

![Sample](contents/tittle.png)

## Description

ProcTop is a cli monitoring tool for watching system information (ram , cpu , disk , avg and ..)

It's using a time interval for refreshing detail and its helpful for benchmarking and seeing the status of the system

## Requirements

- go1.19 or above is required.

## Build

```
git clone https://github.com/SonyaCore/proc-top.git
cd proc-top
go build .
```

## Usage

`proc-top` arguments

> To view full information use -full in proc-top binary

```
  -cpu
    	Show Cpu info (default true)
  -disk
    	Show disk usage
  -full
    	Show all information
  -interval int
    	refresh screen per second (default 1)
  -kernel
    	Show kernel info & uptime (default true)
  -load
    	Show load average
  -memory
    	Show memory usage
  -sensors
    	Show sensors
  -swap
    	Show swap usage
  -version
    	Show version & exit
```

## License

Licensed under the [GPL-3][license] license.

[contributors-shield]: https://img.shields.io/github/contributors/SonyaCore/proc-top?style=flat
[contributors-url]: https://github.com/SonyaCore/proc-top/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/SonyaCore/proc-top?style=flat
[forks-url]: https://github.com/SonyaCore/proc-top/network/members
[stars-shield]: https://img.shields.io/github/stars/SonyaCore/proc-top?style=flat
[stars-url]: https://github.com/SonyaCore/proc-top/stargazers
[issues-shield]: https://img.shields.io/github/issues/SonyaCore/proc-top?style=flat
[issues-url]: https://github.com/SonyaCore/proc-top/issues
[version]: https://img.shields.io/badge/Version-0.4-blue
[goversion]: https://img.shields.io/github/go-mod/go-version/SonyaCore/proc-top/master
[go]: https://img.shields.io/badge/Go-cyan?logo=go
[license]: LICENSE
