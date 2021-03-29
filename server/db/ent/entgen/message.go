// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Author holds the value of the "author" field.
	Author uint64 `json:"author,omitempty"`
	// Createdat holds the value of the "createdat" field.
	Createdat time.Time `json:"createdat,omitempty"`
	// Editedat holds the value of the "editedat" field.
	Editedat time.Time `json:"editedat,omitempty"`
	// Replyto holds the value of the "replyto" field.
	Replyto uint64 `json:"replyto,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges           MessageEdges `json:"edges"`
	channel_message *uint64
	user_message    *uint64
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Override holds the value of the override edge.
	Override []*Override `json:"override,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// OverrideOrErr returns the Override value or an error if the edge
// was not loaded in eager-loading.
func (e MessageEdges) OverrideOrErr() ([]*Override, error) {
	if e.loadedTypes[1] {
		return e.Override, nil
	}
	return nil, &NotLoadedError{edge: "override"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldID, message.FieldAuthor, message.FieldReplyto:
			values[i] = &sql.NullInt64{}
		case message.FieldCreatedat, message.FieldEditedat:
			values[i] = &sql.NullTime{}
		case message.ForeignKeys[0]: // channel_message
			values[i] = &sql.NullInt64{}
		case message.ForeignKeys[1]: // user_message
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Message", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = uint64(value.Int64)
		case message.FieldAuthor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				m.Author = uint64(value.Int64)
			}
		case message.FieldCreatedat:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdat", values[i])
			} else if value.Valid {
				m.Createdat = value.Time
			}
		case message.FieldEditedat:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field editedat", values[i])
			} else if value.Valid {
				m.Editedat = value.Time
			}
		case message.FieldReplyto:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field replyto", values[i])
			} else if value.Valid {
				m.Replyto = uint64(value.Int64)
			}
		case message.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field channel_message", value)
			} else if value.Valid {
				m.channel_message = new(uint64)
				*m.channel_message = uint64(value.Int64)
			}
		case message.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_message", value)
			} else if value.Valid {
				m.user_message = new(uint64)
				*m.user_message = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Message entity.
func (m *Message) QueryUser() *UserQuery {
	return (&MessageClient{config: m.config}).QueryUser(m)
}

// QueryOverride queries the "override" edge of the Message entity.
func (m *Message) QueryOverride() *OverrideQuery {
	return (&MessageClient{config: m.config}).QueryOverride(m)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return (&MessageClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("entgen: Message is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", author=")
	builder.WriteString(fmt.Sprintf("%v", m.Author))
	builder.WriteString(", createdat=")
	builder.WriteString(m.Createdat.Format(time.ANSIC))
	builder.WriteString(", editedat=")
	builder.WriteString(m.Editedat.Format(time.ANSIC))
	builder.WriteString(", replyto=")
	builder.WriteString(fmt.Sprintf("%v", m.Replyto))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message

func (m Messages) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}