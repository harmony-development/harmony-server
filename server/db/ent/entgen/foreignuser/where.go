// Code generated by entc, DO NOT EDIT.

package foreignuser

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Foreignid applies equality check predicate on the "foreignid" field. It's identical to ForeignidEQ.
func Foreignid(v uint64) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForeignid), v))
	})
}

// Host applies equality check predicate on the "host" field. It's identical to HostEQ.
func Host(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHost), v))
	})
}

// ForeignidEQ applies the EQ predicate on the "foreignid" field.
func ForeignidEQ(v uint64) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForeignid), v))
	})
}

// ForeignidNEQ applies the NEQ predicate on the "foreignid" field.
func ForeignidNEQ(v uint64) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldForeignid), v))
	})
}

// ForeignidIn applies the In predicate on the "foreignid" field.
func ForeignidIn(vs ...uint64) predicate.ForeignUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ForeignUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldForeignid), v...))
	})
}

// ForeignidNotIn applies the NotIn predicate on the "foreignid" field.
func ForeignidNotIn(vs ...uint64) predicate.ForeignUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ForeignUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldForeignid), v...))
	})
}

// ForeignidGT applies the GT predicate on the "foreignid" field.
func ForeignidGT(v uint64) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldForeignid), v))
	})
}

// ForeignidGTE applies the GTE predicate on the "foreignid" field.
func ForeignidGTE(v uint64) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldForeignid), v))
	})
}

// ForeignidLT applies the LT predicate on the "foreignid" field.
func ForeignidLT(v uint64) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldForeignid), v))
	})
}

// ForeignidLTE applies the LTE predicate on the "foreignid" field.
func ForeignidLTE(v uint64) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldForeignid), v))
	})
}

// HostEQ applies the EQ predicate on the "host" field.
func HostEQ(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHost), v))
	})
}

// HostNEQ applies the NEQ predicate on the "host" field.
func HostNEQ(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHost), v))
	})
}

// HostIn applies the In predicate on the "host" field.
func HostIn(vs ...string) predicate.ForeignUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ForeignUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHost), v...))
	})
}

// HostNotIn applies the NotIn predicate on the "host" field.
func HostNotIn(vs ...string) predicate.ForeignUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ForeignUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHost), v...))
	})
}

// HostGT applies the GT predicate on the "host" field.
func HostGT(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHost), v))
	})
}

// HostGTE applies the GTE predicate on the "host" field.
func HostGTE(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHost), v))
	})
}

// HostLT applies the LT predicate on the "host" field.
func HostLT(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHost), v))
	})
}

// HostLTE applies the LTE predicate on the "host" field.
func HostLTE(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHost), v))
	})
}

// HostContains applies the Contains predicate on the "host" field.
func HostContains(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHost), v))
	})
}

// HostHasPrefix applies the HasPrefix predicate on the "host" field.
func HostHasPrefix(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHost), v))
	})
}

// HostHasSuffix applies the HasSuffix predicate on the "host" field.
func HostHasSuffix(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHost), v))
	})
}

// HostEqualFold applies the EqualFold predicate on the "host" field.
func HostEqualFold(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHost), v))
	})
}

// HostContainsFold applies the ContainsFold predicate on the "host" field.
func HostContainsFold(v string) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHost), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ForeignUser) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ForeignUser) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
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
func Not(p predicate.ForeignUser) predicate.ForeignUser {
	return predicate.ForeignUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
