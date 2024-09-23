package builder

type Builder struct {
	requiredString string
	requiredInt    int

	optionalString string
	optionalInt    int
}

func NewBuilder(requiredString string, requiredInt int) *Builder {
	return &Builder{
		requiredString: requiredString,
		requiredInt:    requiredInt,
	}
}

func (b *Builder) WithOptionalString(optionalString string) *Builder {
	b.optionalString = optionalString
	return b
}

func (b *Builder) WithOptionalInt(optionalInt int) *Builder {
	b.optionalInt = optionalInt
	return b
}
