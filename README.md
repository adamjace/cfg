# cfganalyze

A simple config analysis tool that will help keep config
files in sync by detecting and alerting for missing keys.

Currently supports `json` and `env` config types.

## Usage

### Analyzing local files

```go
  a := cfganalyze.NewAnalyzer()

  missingKeys, err := a.AnalyzeJson("a.json", "b.json")
  if err != nil {
    log.Println(err)
    return
  }

  for _, key := range missingKeys {
    log.Printf("Found missing key: %s", key)
  }
```

### Analyzing local and remote files

```go
  a := cfganalyze.Connect("host-alias")

  missingKeys, err := a.AnalyzeJson("config.json", "~/home/ubuntu/config.json")
  if err != nil {
    log.Println(err)
    return
  }

  for _, key := range missingKeys {
    log.Printf("Found missing key: %s", key)
  }
```
