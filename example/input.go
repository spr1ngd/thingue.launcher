package example

// goverter:converter
type Converter interface {
	ConvertItems(source []Input) []Output

	// goverter:ignore Irrelevant
	// goverter:map Nested.AgeInYears Age
	Convert(source Input) Output
}

type Input struct {
	Name   string
	Nested InputNested
}
type InputNested struct {
	AgeInYears int
}
type Output struct {
	Name       string
	Age        int
	Irrelevant bool
}
