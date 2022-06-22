// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package car

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	})
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	})
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumber), v))
	})
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumber), v...))
	})
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumber), v...))
	})
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumber), v))
	})
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumber), v))
	})
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumber), v))
	})
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumber), v))
	})
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNumber), v))
	})
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNumber), v))
	})
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNumber), v))
	})
}

// NumberIsNil applies the IsNil predicate on the "number" field.
func NumberIsNil() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNumber)))
	})
}

// NumberNotNil applies the NotNil predicate on the "number" field.
func NumberNotNil() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNumber)))
	})
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNumber), v))
	})
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNumber), v))
	})
}

// HasRentals applies the HasEdge predicate on the "rentals" edge.
func HasRentals() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RentalsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RentalsTable, RentalsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRentalsWith applies the HasEdge predicate on the "rentals" edge with a given conditions (other predicates).
func HasRentalsWith(preds ...predicate.Rental) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RentalsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RentalsTable, RentalsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
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
func Not(p predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		p(s.Not())
	})
}
