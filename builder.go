package querybuilder

type Builder interface {
	Compile() QueryBuilder
}

func Build(b Builder) QueryBuilder {
	return b.Compile()
}
