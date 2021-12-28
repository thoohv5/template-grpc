// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/thoohv5/template-grpc/internal/ent/phoneaccount"
)

// PhoneAccount is the model entity for the PhoneAccount schema.
type PhoneAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserIdentity holds the value of the "user_identity" field.
	UserIdentity string `json:"user_identity,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PhoneAccount) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // user_identity
		&sql.NullString{}, // phone
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // deleted_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PhoneAccount fields.
func (pa *PhoneAccount) assignValues(values ...interface{}) error {
	if m, n := len(values), len(phoneaccount.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user_identity", values[0])
	} else if value.Valid {
		pa.UserIdentity = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[1])
	} else if value.Valid {
		pa.Phone = value.String
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[2])
	} else if value.Valid {
		pa.CreatedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[3])
	} else if value.Valid {
		pa.UpdatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field deleted_at", values[4])
	} else if value.Valid {
		pa.DeletedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this PhoneAccount.
// Note that, you need to call PhoneAccount.Unwrap() before calling this method, if this PhoneAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *PhoneAccount) Update() *PhoneAccountUpdateOne {
	return (&PhoneAccountClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *PhoneAccount) Unwrap() *PhoneAccount {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: PhoneAccount is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *PhoneAccount) String() string {
	var builder strings.Builder
	builder.WriteString("PhoneAccount(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", user_identity=")
	builder.WriteString(pa.UserIdentity)
	builder.WriteString(", phone=")
	builder.WriteString(pa.Phone)
	builder.WriteString(", created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(pa.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PhoneAccounts is a parsable slice of PhoneAccount.
type PhoneAccounts []*PhoneAccount

func (pa PhoneAccounts) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
