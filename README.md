# Mongolog
<p align="left">
<img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/itzngga/mongolog?style=flat-square">
<img alt="GitHub forks" src="https://img.shields.io/github/forks/itzngga/mongolog?style=flat-square">
<img alt="GitHub watchers" src="https://img.shields.io/github/watchers/itzngga/mongolog?style=flat-square">
<img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/itzngga/mongolog?style=flat-square">
</p>

# Description
a default mongodb driver log implementation

# Installation
```bash 
go get github.com/itzngga/mongolog
```

# Usage
```go
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
clientOptions.SetMonitor(mongolog.New()) // <- the logger
```
# Example
to this [example](https://github.com/itzngga/mongolog/tree/main/example)

# License
[GNU](https://github.com/itzngga/mongolog/blob/master/LICENSE)

# Contribute
Pull Request are pleased to