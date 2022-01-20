// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/YEVER/pkg/database/ent/wuser"
)

// Wuser is the model entity for the Wuser schema.
type Wuser struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// UID holds the value of the "uid" field.
	UID string `json:"uid,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wuser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case wuser.FieldID, wuser.FieldPhone, wuser.FieldUID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Wuser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wuser fields.
func (w *Wuser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wuser.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				w.ID = value.String
			}
		case wuser.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				w.Phone = value.String
			}
		case wuser.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				w.UID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Wuser.
// Note that you need to call Wuser.Unwrap() before calling this method if this Wuser
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wuser) Update() *WuserUpdateOne {
	return (&WuserClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Wuser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Wuser) Unwrap() *Wuser {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Wuser is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wuser) String() string {
	var builder strings.Builder
	builder.WriteString("Wuser(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", phone=")
	builder.WriteString(w.Phone)
	builder.WriteString(", uid=")
	builder.WriteString(w.UID)
	builder.WriteByte(')')
	return builder.String()
}

// Wusers is a parsable slice of Wuser.
type Wusers []*Wuser

func (w Wusers) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
