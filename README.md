<!-- markdownlint-disable MD010 -->
# koanfenv

`koanfenv` provides callbacks which convert environment variables to [koanf][]
keys.  These callbacks are used for [`env.Provider`][provider]
.

[koanf]: https://github.com/knadh/koanf
[provider]: https://pkg.go.dev/github.com/knadh/koanf@v1.3.3/providers/env#Provider

## Usage

```go
config := struct{
	Foo1 string `koanf:"foo1" env:"APP_FOO1"`
}{}

k := koanf.New(".")
if err := k.Load(env.Provider("APP_", ".", koanfenv.ByStruct(&config)), nil); err != nil {
	log.Fatal(err)
}
```

```go
m := map[string]string{
	"APP_FOO1": "foo1",
}

k := koanf.New(".")
if err := k.Load(env.Provider("APP_", ".", koanfenv.ByMap(m)), nil); err != nil {
	log.Fatal(err)
}
```

## License

[MIT](https://dochang.mit-license.org/)
