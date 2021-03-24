# headergrep
A small tool for examining headers from a web request

## Features
  * json output
  * curl like syntax
  * Header retrieval from arbitrary url
  * Configurable list of expected and unexpected headers

## Usage
`headergrep http://mysite.example.com:8080/helloworld`

### Flags
  * `-k`: ignore cert issues
  * `-X`: specify http verb
  * `-c`: config

### Config
```
expected = ["Strict-Transport-Security"]
unexpected = ["Server", "Device-Memory"]
```
