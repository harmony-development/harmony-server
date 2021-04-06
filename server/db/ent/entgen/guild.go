// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
)

// Guild is the model entity for the Guild schema.
type Guild struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner uint64 `json:"owner,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Picture holds the value of the "picture" field.
	Picture string `json:"picture,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata []byte `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GuildQuery when eager-loading is set.
	Edges GuildEdges `json:"edges"`
}

// GuildEdges holds the relations/edges for other nodes in the graph.
type GuildEdges struct {
	// Invite holds the value of the invite edge.
	Invite []*Invite `json:"invite,omitempty"`
	// Bans holds the value of the bans edge.
	Bans []*User `json:"bans,omitempty"`
	// Channel holds the value of the channel edge.
	Channel []*Channel `json:"channel,omitempty"`
	// Role holds the value of the role edge.
	Role []*Role `json:"role,omitempty"`
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// InviteOrErr returns the Invite value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) InviteOrErr() ([]*Invite, error) {
	if e.loadedTypes[0] {
		return e.Invite, nil
	}
	return nil, &NotLoadedError{edge: "invite"}
}

// BansOrErr returns the Bans value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) BansOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Bans, nil
	}
	return nil, &NotLoadedError{edge: "bans"}
}

// ChannelOrErr returns the Channel value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) ChannelOrErr() ([]*Channel, error) {
	if e.loadedTypes[2] {
		return e.Channel, nil
	}
	return nil, &NotLoadedError{edge: "channel"}
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) RoleOrErr() ([]*Role, error) {
	if e.loadedTypes[3] {
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[4] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Guild) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case guild.FieldMetadata:
			values[i] = &[]byte{}
		case guild.FieldID, guild.FieldOwner:
			values[i] = &sql.NullInt64{}
		case guild.FieldName, guild.FieldPicture:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Guild", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Guild fields.
func (gu *Guild) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case guild.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gu.ID = uint64(value.Int64)
		case guild.FieldOwner:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				gu.Owner = uint64(value.Int64)
			}
		case guild.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gu.Name = value.String
			}
		case guild.FieldPicture:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field picture", values[i])
			} else if value.Valid {
				gu.Picture = value.String
			}
		case guild.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil {
				gu.Metadata = *value
			}
		}
	}
	return nil
}

// QueryInvite queries the "invite" edge of the Guild entity.
func (gu *Guild) QueryInvite() *InviteQuery {
	return (&GuildClient{config: gu.config}).QueryInvite(gu)
}

// QueryBans queries the "bans" edge of the Guild entity.
func (gu *Guild) QueryBans() *UserQuery {
	return (&GuildClient{config: gu.config}).QueryBans(gu)
}

// QueryChannel queries the "channel" edge of the Guild entity.
func (gu *Guild) QueryChannel() *ChannelQuery {
	return (&GuildClient{config: gu.config}).QueryChannel(gu)
}

// QueryRole queries the "role" edge of the Guild entity.
func (gu *Guild) QueryRole() *RoleQuery {
	return (&GuildClient{config: gu.config}).QueryRole(gu)
}

// QueryUser queries the "user" edge of the Guild entity.
func (gu *Guild) QueryUser() *UserQuery {
	return (&GuildClient{config: gu.config}).QueryUser(gu)
}

// Update returns a builder for updating this Guild.
// Note that you need to call Guild.Unwrap() before calling this method if this Guild
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *Guild) Update() *GuildUpdateOne {
	return (&GuildClient{config: gu.config}).UpdateOne(gu)
}

// Unwrap unwraps the Guild entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gu *Guild) Unwrap() *Guild {
	tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("entgen: Guild is not a transactional entity")
	}
	gu.config.driver = tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *Guild) String() string {
	var builder strings.Builder
	builder.WriteString("Guild(")
	builder.WriteString(fmt.Sprintf("id=%v", gu.ID))
	builder.WriteString(", owner=")
	builder.WriteString(fmt.Sprintf("%v", gu.Owner))
	builder.WriteString(", name=")
	builder.WriteString(gu.Name)
	builder.WriteString(", picture=")
	builder.WriteString(gu.Picture)
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", gu.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// Guilds is a parsable slice of Guild.
type Guilds []*Guild

func (gu Guilds) config(cfg config) {
	for _i := range gu {
		gu[_i].config = cfg
	}
}
