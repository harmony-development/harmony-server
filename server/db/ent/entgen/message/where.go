// Code generated by entc, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	v1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Createdat applies equality check predicate on the "createdat" field. It's identical to CreatedatEQ.
func Createdat(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedat), v))
	})
}

// Editedat applies equality check predicate on the "editedat" field. It's identical to EditedatEQ.
func Editedat(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEditedat), v))
	})
}

// Metadata applies equality check predicate on the "metadata" field. It's identical to MetadataEQ.
func Metadata(v *v1.Metadata) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetadata), v))
	})
}

// Override applies equality check predicate on the "override" field. It's identical to OverrideEQ.
func Override(v *v1.Override) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOverride), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v *v1.Content) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// CreatedatEQ applies the EQ predicate on the "createdat" field.
func CreatedatEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedat), v))
	})
}

// CreatedatNEQ applies the NEQ predicate on the "createdat" field.
func CreatedatNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedat), v))
	})
}

// CreatedatIn applies the In predicate on the "createdat" field.
func CreatedatIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedat), v...))
	})
}

// CreatedatNotIn applies the NotIn predicate on the "createdat" field.
func CreatedatNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedat), v...))
	})
}

// CreatedatGT applies the GT predicate on the "createdat" field.
func CreatedatGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedat), v))
	})
}

// CreatedatGTE applies the GTE predicate on the "createdat" field.
func CreatedatGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedat), v))
	})
}

// CreatedatLT applies the LT predicate on the "createdat" field.
func CreatedatLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedat), v))
	})
}

// CreatedatLTE applies the LTE predicate on the "createdat" field.
func CreatedatLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedat), v))
	})
}

// EditedatEQ applies the EQ predicate on the "editedat" field.
func EditedatEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEditedat), v))
	})
}

// EditedatNEQ applies the NEQ predicate on the "editedat" field.
func EditedatNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEditedat), v))
	})
}

// EditedatIn applies the In predicate on the "editedat" field.
func EditedatIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEditedat), v...))
	})
}

// EditedatNotIn applies the NotIn predicate on the "editedat" field.
func EditedatNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEditedat), v...))
	})
}

// EditedatGT applies the GT predicate on the "editedat" field.
func EditedatGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEditedat), v))
	})
}

// EditedatGTE applies the GTE predicate on the "editedat" field.
func EditedatGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEditedat), v))
	})
}

// EditedatLT applies the LT predicate on the "editedat" field.
func EditedatLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEditedat), v))
	})
}

// EditedatLTE applies the LTE predicate on the "editedat" field.
func EditedatLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEditedat), v))
	})
}

// EditedatIsNil applies the IsNil predicate on the "editedat" field.
func EditedatIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEditedat)))
	})
}

// EditedatNotNil applies the NotNil predicate on the "editedat" field.
func EditedatNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEditedat)))
	})
}

// MetadataEQ applies the EQ predicate on the "metadata" field.
func MetadataEQ(v *v1.Metadata) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetadata), v))
	})
}

// MetadataNEQ applies the NEQ predicate on the "metadata" field.
func MetadataNEQ(v *v1.Metadata) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMetadata), v))
	})
}

// MetadataIn applies the In predicate on the "metadata" field.
func MetadataIn(vs ...*v1.Metadata) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMetadata), v...))
	})
}

// MetadataNotIn applies the NotIn predicate on the "metadata" field.
func MetadataNotIn(vs ...*v1.Metadata) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMetadata), v...))
	})
}

// MetadataGT applies the GT predicate on the "metadata" field.
func MetadataGT(v *v1.Metadata) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMetadata), v))
	})
}

// MetadataGTE applies the GTE predicate on the "metadata" field.
func MetadataGTE(v *v1.Metadata) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMetadata), v))
	})
}

// MetadataLT applies the LT predicate on the "metadata" field.
func MetadataLT(v *v1.Metadata) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMetadata), v))
	})
}

// MetadataLTE applies the LTE predicate on the "metadata" field.
func MetadataLTE(v *v1.Metadata) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMetadata), v))
	})
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMetadata)))
	})
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMetadata)))
	})
}

// OverrideEQ applies the EQ predicate on the "override" field.
func OverrideEQ(v *v1.Override) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOverride), v))
	})
}

// OverrideNEQ applies the NEQ predicate on the "override" field.
func OverrideNEQ(v *v1.Override) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOverride), v))
	})
}

// OverrideIn applies the In predicate on the "override" field.
func OverrideIn(vs ...*v1.Override) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOverride), v...))
	})
}

// OverrideNotIn applies the NotIn predicate on the "override" field.
func OverrideNotIn(vs ...*v1.Override) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOverride), v...))
	})
}

// OverrideGT applies the GT predicate on the "override" field.
func OverrideGT(v *v1.Override) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOverride), v))
	})
}

// OverrideGTE applies the GTE predicate on the "override" field.
func OverrideGTE(v *v1.Override) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOverride), v))
	})
}

// OverrideLT applies the LT predicate on the "override" field.
func OverrideLT(v *v1.Override) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOverride), v))
	})
}

// OverrideLTE applies the LTE predicate on the "override" field.
func OverrideLTE(v *v1.Override) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOverride), v))
	})
}

// OverrideIsNil applies the IsNil predicate on the "override" field.
func OverrideIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOverride)))
	})
}

// OverrideNotNil applies the NotNil predicate on the "override" field.
func OverrideNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOverride)))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v *v1.Content) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v *v1.Content) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...*v1.Content) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...*v1.Content) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v *v1.Content) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v *v1.Content) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v *v1.Content) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v *v1.Content) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChannel applies the HasEdge predicate on the "channel" edge.
func HasChannel() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChannelTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChannelWith applies the HasEdge predicate on the "channel" edge with a given conditions (other predicates).
func HasChannelWith(preds ...predicate.Channel) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChannelInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReplies applies the HasEdge predicate on the "replies" edge.
func HasReplies() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepliesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepliesWith applies the HasEdge predicate on the "replies" edge with a given conditions (other predicates).
func HasRepliesWith(preds ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		p(s.Not())
	})
}
