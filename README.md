# beeconfig
beeconfig is used to load config file under conf for beego.

It supports multi adapter and keep the code module's config files independent and looks clear.

For example, your app conf structure is just like this:
```
conf/prod/database.ini
conf/dev/database.ini
```

You want to load the specific file by runmode, so you just need this for solving this problem.

```go

cf, err := beeconfig.Load("database.ini")

...

// or other config
// cf2, err := beeconfig.Load("module_config.json")

...

```

It will load conf/dev/database.ini if you run in dev mode, or load conf/prod/database.ini in prod mode.

If can not find the file in above dir, it will load conf/database.ini.

- add func ParseDIYToMap and ParseDIYToMaps to convert json object to golang map
