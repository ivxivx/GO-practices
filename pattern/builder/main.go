package builder

func main() {
	builder := NewBuilder("required", 1)

	builder.WithOptionalString("optional")

	builder.WithOptionalInt(2)
}
