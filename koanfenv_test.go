package koanfenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type config struct {
	Foo1 string `koanf:"foo1Koanf" env:"foo1Env"`
	Foo2 string `koanf:"" env:"foo2Env"`
	Foo3 string `env:"foo3Env"`
	Foo4 string `koanf:"-" env:"foo4Env"`
	Foo5 string `koanf:"foo5Koanf" env:"-"`
	Foo6 string `koanf:"foo6Koanf" env:""`
	Foo7 string `koanf:"foo7Koanf"`
	Foo8 string
}

func TestEnv2Koanf(t *testing.T) {
	conf := config{}
	fn := ByStruct(&conf)

	cases := []struct {
		envTag   string
		koanfKey string
	}{
		{
			envTag:   "foo1Env",
			koanfKey: "foo1Koanf",
		},
		{
			envTag:   "Foo1Env",
			koanfKey: "",
		},
		{
			envTag:   "FOO1Env",
			koanfKey: "",
		},
		{
			envTag:   "foo2Env",
			koanfKey: "Foo2",
		},
		{
			envTag:   "foo3Env",
			koanfKey: "Foo3",
		},
		{
			envTag:   "foo4Env",
			koanfKey: "",
		},
		{
			envTag:   "Foo5",
			koanfKey: "",
		},
		{
			envTag:   "Foo6",
			koanfKey: "",
		},
		{
			envTag:   "Foo7",
			koanfKey: "",
		},
		{
			envTag:   "Foo8",
			koanfKey: "",
		},
	}

	for _, c := range cases {
		k := fn(c.envTag)
		assert.Equal(t, k, c.koanfKey)
	}
}
