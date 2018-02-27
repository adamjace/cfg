# cfganalyze

A simple config analysis tool aimed to help keep config
files in sync by scanning for missing keys.

Currently supports `json` and `env` config types.

## Usage

### Compare local files

```go
  a := cfganalyze.NewAnalyzer()

  missingKeys, err := a.AnalyzeEnv("config/.env", "config/.env.example")
  if err != nil {
    log.Println(err)
    return
  }

  for _, key := range missingKeys {
    log.Printf("Found missing key: %s", key)
  }
```

### Compare local with remote

```go
  a := cfganalyze.Connect("host-alias")

  missingKeys, err := a.AnalyzeJson("config.json", "~/home/ubuntu/app/config.json")
  if err != nil {
    log.Println(err)
    return
  }

  for _, key := range missingKeys {
    log.Printf("Found missing key: %s", key)
  }
```
