// Code generated by entc, DO NOT EDIT.

package userinfo

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/thoohv5/template-grpc/internal/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserIdentity applies equality check predicate on the "user_identity" field. It's identical to UserIdentityEQ.
func UserIdentity(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserIdentity), v))
	})
}

// Channel applies equality check predicate on the "channel" field. It's identical to ChannelEQ.
func Channel(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannel), v))
	})
}

// Form applies equality check predicate on the "form" field. It's identical to FormEQ.
func Form(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForm), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// UserIdentityEQ applies the EQ predicate on the "user_identity" field.
func UserIdentityEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityNEQ applies the NEQ predicate on the "user_identity" field.
func UserIdentityNEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityIn applies the In predicate on the "user_identity" field.
func UserIdentityIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserIdentity), v...))
	})
}

// UserIdentityNotIn applies the NotIn predicate on the "user_identity" field.
func UserIdentityNotIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserIdentity), v...))
	})
}

// UserIdentityGT applies the GT predicate on the "user_identity" field.
func UserIdentityGT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityGTE applies the GTE predicate on the "user_identity" field.
func UserIdentityGTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityLT applies the LT predicate on the "user_identity" field.
func UserIdentityLT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityLTE applies the LTE predicate on the "user_identity" field.
func UserIdentityLTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityContains applies the Contains predicate on the "user_identity" field.
func UserIdentityContains(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityHasPrefix applies the HasPrefix predicate on the "user_identity" field.
func UserIdentityHasPrefix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityHasSuffix applies the HasSuffix predicate on the "user_identity" field.
func UserIdentityHasSuffix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityIsNil applies the IsNil predicate on the "user_identity" field.
func UserIdentityIsNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserIdentity)))
	})
}

// UserIdentityNotNil applies the NotNil predicate on the "user_identity" field.
func UserIdentityNotNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserIdentity)))
	})
}

// UserIdentityEqualFold applies the EqualFold predicate on the "user_identity" field.
func UserIdentityEqualFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserIdentity), v))
	})
}

// UserIdentityContainsFold applies the ContainsFold predicate on the "user_identity" field.
func UserIdentityContainsFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserIdentity), v))
	})
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannel), v))
	})
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChannel), v))
	})
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...int32) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChannel), v...))
	})
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...int32) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChannel), v...))
	})
}

// ChannelGT applies the GT predicate on the "channel" field.
func ChannelGT(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChannel), v))
	})
}

// ChannelGTE applies the GTE predicate on the "channel" field.
func ChannelGTE(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChannel), v))
	})
}

// ChannelLT applies the LT predicate on the "channel" field.
func ChannelLT(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChannel), v))
	})
}

// ChannelLTE applies the LTE predicate on the "channel" field.
func ChannelLTE(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChannel), v))
	})
}

// FormEQ applies the EQ predicate on the "form" field.
func FormEQ(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForm), v))
	})
}

// FormNEQ applies the NEQ predicate on the "form" field.
func FormNEQ(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldForm), v))
	})
}

// FormIn applies the In predicate on the "form" field.
func FormIn(vs ...int32) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldForm), v...))
	})
}

// FormNotIn applies the NotIn predicate on the "form" field.
func FormNotIn(vs ...int32) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldForm), v...))
	})
}

// FormGT applies the GT predicate on the "form" field.
func FormGT(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldForm), v))
	})
}

// FormGTE applies the GTE predicate on the "form" field.
func FormGTE(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldForm), v))
	})
}

// FormLT applies the LT predicate on the "form" field.
func FormLT(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldForm), v))
	})
}

// FormLTE applies the LTE predicate on the "form" field.
func FormLTE(v int32) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldForm), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.UserInfo) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.UserInfo) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
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
func Not(p predicate.UserInfo) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}
