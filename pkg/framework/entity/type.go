package entity

// TODO generate from type.cue
type RawMeta struct {
	Extensions []string
}

func (m RawMeta) _private() {}

// TODO
type CoreStructuredMeta struct {
}

func (m CoreStructuredMeta) _private() {}

// type CoreStructuredKind[T thema.Assignee] struct {
// 	Name    string
// 	Lineage thema.ConvergentLineage[T]
// 	// Maturity StructuredMaturity
// 	Maturity string
// }

// TODO
type CustomStructuredMeta struct {
}

func (m CustomStructuredMeta) _private() {}

// type CustomStructuredKind struct {
// 	Name    string
// 	Lineage thema.Lineage
// 	// Maturity StructuredMaturity
// 	Maturity string
// }

// TODO
type SlotImplMeta struct {
}

// SomeKindMeta is an interface type to abstract over the different
// kind metadata struct types: [RawMeta], [CoreStructuredMeta], [CustomStructuredMeta].
//
// It is the traditional interface counterpart to the generic type constraint KindMetas.
type SomeKindMeta interface {
	_private()
}

// KindMetas is a type parameter that comprises the base possible set of
// kind metadata configurations.
type KindMetas = interface {
	RawMeta | CoreStructuredMeta | CustomStructuredMeta // | SlotImplMeta
}
