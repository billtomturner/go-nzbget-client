# go-nzbget-client

## Description

This project is a Go client SDK for the [NZBGet](https://nzbget.net/) usenet
downloader.

## Usage

The `nzbget_test.go` file includes more comprehensive details, and structs are
documented within `nzbget.go`.  Usage is pretty straight forward:

```go
package main

import "github.com/billtomturner/go-nzbget-client"

client, err := nzbget.New("http://localhost:6789", "username", "password")

// Get configuration
config, err := client.Config()

//  Get server transfer volumes
volumes, err := client.ServerVolumes()

// Get active file groups
volumes, err := client.FileGroups()

// Get server status
status, err := client.Status()

// Get server file group history
history, err := client.History()
```
