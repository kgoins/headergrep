# headergrep
A small tool for examining headers from a web request

## Features
  * curl like syntax
  * Header retrieval from arbitrary url
  * Configurable list of expected and unexpected headers

## Usage
`headergrep http://mysite.example.com:8080/helloworld`

### Flags
  * `-k`: ignore cert issues
  * `-H`: add header to outbound request
  * `-X`: specify http verb

### Config
```
[expected]
Strict-Transport-Security

[unexpected]
Server
Device-Memory
```