// Code generated by ent, DO NOT EDIT.

package sourcenamespace

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldLTE(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldEQ(FieldNamespace, v))
}

// SourceID applies equality check predicate on the "source_id" field. It's identical to SourceIDEQ.
func SourceID(v int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldEQ(FieldSourceID, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldContainsFold(FieldNamespace, v))
}

// SourceIDEQ applies the EQ predicate on the "source_id" field.
func SourceIDEQ(v int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldEQ(FieldSourceID, v))
}

// SourceIDNEQ applies the NEQ predicate on the "source_id" field.
func SourceIDNEQ(v int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldNEQ(FieldSourceID, v))
}

// SourceIDIn applies the In predicate on the "source_id" field.
func SourceIDIn(vs ...int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldIn(FieldSourceID, vs...))
}

// SourceIDNotIn applies the NotIn predicate on the "source_id" field.
func SourceIDNotIn(vs ...int) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.FieldNotIn(FieldSourceID, vs...))
}

// HasSourceType applies the HasEdge predicate on the "source_type" edge.
func HasSourceType() predicate.SourceNamespace {
	return predicate.SourceNamespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SourceTypeTable, SourceTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSourceTypeWith applies the HasEdge predicate on the "source_type" edge with a given conditions (other predicates).
func HasSourceTypeWith(preds ...predicate.SourceType) predicate.SourceNamespace {
	return predicate.SourceNamespace(func(s *sql.Selector) {
		step := newSourceTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNames applies the HasEdge predicate on the "names" edge.
func HasNames() predicate.SourceNamespace {
	return predicate.SourceNamespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, NamesTable, NamesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamesWith applies the HasEdge predicate on the "names" edge with a given conditions (other predicates).
func HasNamesWith(preds ...predicate.SourceName) predicate.SourceNamespace {
	return predicate.SourceNamespace(func(s *sql.Selector) {
		step := newNamesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SourceNamespace) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SourceNamespace) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SourceNamespace) predicate.SourceNamespace {
	return predicate.SourceNamespace(sql.NotPredicates(p))
}
