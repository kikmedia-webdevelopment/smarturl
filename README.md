# tiny URL Shortener for `smart testsolutions`

![Test Go Application](https://github.com/juliankoehn/mchurl/workflows/Test%20Go%20Application/badge.svg)

`MCHURL` ist ein URL-Shortening Service für das [smart testsolutions](https://smart-testsolutions.de/). 

# cli commands

- mchurl create -u https://www.modecentrum-hamburg.de/das-mch/service.html -c code
- - `INFO[0000] ShortURL by code has been created`
- mchurl delete -c code
- - `INFO[0000] ShortURL by code has been deleted`

# config
```yaml
Web:
  ListenAddr: :8080 // listen address for the http server
  BaseURL: "http://localhost:8080"
DB:
    Driver: "sqlite" // currently: sqlite
    URL: "database.db" // path / filename of db file
    IDLength: 4 // length of autogenerated IDs
```

# Screenshots

## Dashboard
![Dashboard](/docs/dashboard.png?raw=true)

## Links
![Link List](/docs/link-list.png?raw=true)

![Link Create](/docs/link-create.png?raw=true)

![Link Delete](/docs/link-delete.png?raw=true)

![Link Edit](/docs/link-edit.png?raw=true)
