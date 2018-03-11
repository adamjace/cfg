# cfg

A simple config analysis tool aimed to help keep configuration
files in sync by scanning for missing keys.

This package currently supports `json` and `env` config types.

## Usage

### Compare local files

#### Scan

```go
  c := cfg.Config{
    WorkingPath: "config/.env",
    MasterPath:  "config/.env.example",
  }

  keys, _ := cfg.ScanEnv(c)

  for _, k := range keys {
    log.Printf("Uh oh! Found a missing key: %s", k)
  }
```

#### Print

```go
  c := cfg.Config{
    WorkingPath: "config/.env",
    MasterPath:  "config/.env.example",
  }

  cfg.PrintEnv(c)

  // (!) found missing keys in config/.env: [foo bar]

```

### Compare local with remote

#### Scan

```go
  c := cfg.Config{
    WorkingPath: "config.json",
    MasterPath:  "/home/ubuntu/app/config.jsob",
    HostAlias:   "host-alias",
  }

  keys, _ := cfg.ScanEnv(c)

  for _, k := range keys {
    log.Printf("Uh oh! Found a missing key: %s", k)
  }
```

#### Print

```go

  c := cfg.Config{
    WorkingPath: "config.json",
    MasterPath:  "/home/ubuntu/app/config.json",
    HostAlias:   "host-alias",
  }

  cfg.PrintJson();

  // (!) found missing keys in config.json: [foo bar]
```
