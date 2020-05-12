package querybuilder

// Builder contract for custom builder
type Builder interface {
	Compile() QueryBuilder
}

// Build compile Builder into QueryBuilder
func Build(b Builder) QueryBuilder {
	return b.Compile()
}
