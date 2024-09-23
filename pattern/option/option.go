package option

type Option struct {
	requiredString string
	requiredInt    int

	optionalString string
	optionalInt    int
}

type OptionFn func(*Option)

func NewOption(options ...OptionFn) *Option {
	option := &Option{}

	for _, opt := range options {
		opt(option)
	}

	return option
}

func WithRequiredString(requiredString string) OptionFn {
	return func(option *Option) {
		option.requiredString = requiredString
	}
}

func WithRequiredInt(requiredInt int) OptionFn {
	return func(option *Option) {
		option.requiredInt = requiredInt
	}
}

func WithOptionalString(optionalString string) OptionFn {
	return func(option *Option) {
		option.optionalString = optionalString
	}
}

func WithOptionalInt(optionalInt int) OptionFn {
	return func(option *Option) {
		option.optionalInt = optionalInt
	}
}
