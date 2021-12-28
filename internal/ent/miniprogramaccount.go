// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/thoohv5/template-grpc/internal/ent/miniprogramaccount"
)

// MiniProgramAccount is the model entity for the MiniProgramAccount schema.
type MiniProgramAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserIdentity holds the value of the "user_identity" field.
	UserIdentity string `json:"user_identity,omitempty"`
	// OpenID holds the value of the "open_id" field.
	OpenID string `json:"open_id,omitempty"`
	// NickName holds the value of the "nick_name" field.
	NickName string `json:"nick_name,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender int32 `json:"gender,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MiniProgramAccount) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // user_identity
		&sql.NullString{}, // open_id
		&sql.NullString{}, // nick_name
		&sql.NullString{}, // avatar_url
		&sql.NullInt64{},  // gender
		&sql.NullString{}, // country
		&sql.NullString{}, // province
		&sql.NullString{}, // city
		&sql.NullString{}, // language
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // deleted_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MiniProgramAccount fields.
func (mpa *MiniProgramAccount) assignValues(values ...interface{}) error {
	if m, n := len(values), len(miniprogramaccount.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	mpa.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user_identity", values[0])
	} else if value.Valid {
		mpa.UserIdentity = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field open_id", values[1])
	} else if value.Valid {
		mpa.OpenID = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field nick_name", values[2])
	} else if value.Valid {
		mpa.NickName = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field avatar_url", values[3])
	} else if value.Valid {
		mpa.AvatarURL = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field gender", values[4])
	} else if value.Valid {
		mpa.Gender = int32(value.Int64)
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field country", values[5])
	} else if value.Valid {
		mpa.Country = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field province", values[6])
	} else if value.Valid {
		mpa.Province = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field city", values[7])
	} else if value.Valid {
		mpa.City = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field language", values[8])
	} else if value.Valid {
		mpa.Language = value.String
	}
	if value, ok := values[9].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[9])
	} else if value.Valid {
		mpa.CreatedAt = value.Time
	}
	if value, ok := values[10].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[10])
	} else if value.Valid {
		mpa.UpdatedAt = value.Time
	}
	if value, ok := values[11].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field deleted_at", values[11])
	} else if value.Valid {
		mpa.DeletedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this MiniProgramAccount.
// Note that, you need to call MiniProgramAccount.Unwrap() before calling this method, if this MiniProgramAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (mpa *MiniProgramAccount) Update() *MiniProgramAccountUpdateOne {
	return (&MiniProgramAccountClient{config: mpa.config}).UpdateOne(mpa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (mpa *MiniProgramAccount) Unwrap() *MiniProgramAccount {
	tx, ok := mpa.config.driver.(*txDriver)
	if !ok {
		panic("ent: MiniProgramAccount is not a transactional entity")
	}
	mpa.config.driver = tx.drv
	return mpa
}

// String implements the fmt.Stringer.
func (mpa *MiniProgramAccount) String() string {
	var builder strings.Builder
	builder.WriteString("MiniProgramAccount(")
	builder.WriteString(fmt.Sprintf("id=%v", mpa.ID))
	builder.WriteString(", user_identity=")
	builder.WriteString(mpa.UserIdentity)
	builder.WriteString(", open_id=")
	builder.WriteString(mpa.OpenID)
	builder.WriteString(", nick_name=")
	builder.WriteString(mpa.NickName)
	builder.WriteString(", avatar_url=")
	builder.WriteString(mpa.AvatarURL)
	builder.WriteString(", gender=")
	builder.WriteString(fmt.Sprintf("%v", mpa.Gender))
	builder.WriteString(", country=")
	builder.WriteString(mpa.Country)
	builder.WriteString(", province=")
	builder.WriteString(mpa.Province)
	builder.WriteString(", city=")
	builder.WriteString(mpa.City)
	builder.WriteString(", language=")
	builder.WriteString(mpa.Language)
	builder.WriteString(", created_at=")
	builder.WriteString(mpa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(mpa.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(mpa.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MiniProgramAccounts is a parsable slice of MiniProgramAccount.
type MiniProgramAccounts []*MiniProgramAccount

func (mpa MiniProgramAccounts) config(cfg config) {
	for _i := range mpa {
		mpa[_i].config = cfg
	}
}
