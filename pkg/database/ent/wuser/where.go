// Code generated by entc, DO NOT EDIT.

package wuser

import (
	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/YEVER/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Wuser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wuser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Wuser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wuser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.Wuser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wuser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUID), v...))
	})
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...string) predicate.Wuser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wuser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUID), v...))
	})
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUID), v))
	})
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUID), v))
	})
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUID), v))
	})
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUID), v))
	})
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUID), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Wuser) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Wuser) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
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
func Not(p predicate.Wuser) predicate.Wuser {
	return predicate.Wuser(func(s *sql.Selector) {
		p(s.Not())
	})
}
