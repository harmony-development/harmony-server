// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/permission"
	"github.com/harmony-development/legato/server/db/ent/entgen/role"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
	config
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionQuery when eager-loading is set.
	Edges           PermissionEdges `json:"edges"`
	role_permission *uint64
}

// PermissionEdges holds the relations/edges for other nodes in the graph.
type PermissionEdges struct {
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[0] {
		if e.Role == nil {
			// The edge role was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			values[i] = &sql.NullInt64{}
		case permission.ForeignKeys[0]: // role_permission
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Permission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = uint64(value.Int64)
		case permission.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field role_permission", value)
			} else if value.Valid {
				pe.role_permission = new(uint64)
				*pe.role_permission = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryRole queries the "role" edge of the Permission entity.
func (pe *Permission) QueryRole() *RoleQuery {
	return (&PermissionClient{config: pe.config}).QueryRole(pe)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return (&PermissionClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("entgen: Permission is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission

func (pe Permissions) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
