# cfg

A simple config analysis tool aimed to help keep configuration
files in sync by scanning for missing keys.

This package currently supports `json` and `env` config types.

## Usage

### Compare local files

```go
  c := cfg.Config{
    WorkingPath: "config/.env",
    MasterPath:  "config/.env.example",
  }

  keys, err := cfg.ScanEnv(c)
  if err != nil {
    log.Println(err)
    return
  }

  for _, k := range keys {
    log.Printf("Uh oh! Found a missing key: %s", k)
  }
```

### Compare local with remote

```go

  c := cfg.Config{
    WorkingPath: "config/.env",
    MasterPath:  "config/.env.example",
    HostAlias:   "host-alias",
  }

  if err := cfg.PrintJson(); err != nil {
    log.Println(err)
    return
  }

  // prints missing keys
```
