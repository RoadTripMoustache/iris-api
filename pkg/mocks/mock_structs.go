// Package mocks contains all the resources to mock data and services to execute easily unit tests.
package mocks

type FakeRequiredValue struct {
	Toto string `required:"true"`
}

type FakeWithSubObject struct {
	Toto FakeRequiredValue `required:"true"`
}

type FakeMin struct {
	Toto int `min:"0"`
}

type FakeMax struct {
	Toto int `max:"10"`
}

type FakeMinLengthString struct {
	Toto string `minLength:"2"`
}

type FakeMinLengthSlice struct {
	Toto []string `minLength:"2"`
}

type FakeMaxLengthString struct {
	Toto string `maxLength:"2"`
}

type FakeMaxLengthSlice struct {
	Toto []string `maxLength:"2"`
}

type FakePattern struct {
	Toto string `pattern:"[A-Z]"`
}

type FakeDatePattern struct {
	Toto string `datePattern:"2006-01-02T15:04:05Z"`
}

type FakeEnumString struct {
	Toto string `enum:"toto;tata"`
}

type FakeEnumSlice struct {
	Toto []string `enum:"toto;tata"`
}

type FakeMultiple struct {
	Tata []string `minLength:"2"`
	Toto []string `enum:"toto;tata"`
}
