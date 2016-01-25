# beeconfig
beeconfig is used to load config file under conf for beego

For example, your app conf structure just like this:
```
conf/prod/database.ini
conf/dev/database.ini
```

You want to load the specific file by runmode, so you just need this for solving this problem.

```go

cf, err := beeconfig.Load("database", "ini")

...

```

