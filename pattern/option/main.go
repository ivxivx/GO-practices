package option

func main() {
	_ = NewOption(
		WithRequiredString("required"),
		WithRequiredInt(1),
		WithOptionalString("optional"),
		WithOptionalInt(2),
	)
}
