// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/emote"
	"github.com/harmony-development/legato/server/db/ent/entgen/file"
	"github.com/harmony-development/legato/server/db/ent/entgen/filehash"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Contenttype holds the value of the "contenttype" field.
	Contenttype string `json:"contenttype,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges          FileEdges `json:"edges"`
	emote_file     *int
	file_hash_file *int
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Filehash holds the value of the filehash edge.
	Filehash *FileHash `json:"filehash,omitempty"`
	// Emote holds the value of the emote edge.
	Emote *Emote `json:"emote,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FilehashOrErr returns the Filehash value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) FilehashOrErr() (*FileHash, error) {
	if e.loadedTypes[0] {
		if e.Filehash == nil {
			// The edge filehash was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: filehash.Label}
		}
		return e.Filehash, nil
	}
	return nil, &NotLoadedError{edge: "filehash"}
}

// EmoteOrErr returns the Emote value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) EmoteOrErr() (*Emote, error) {
	if e.loadedTypes[1] {
		if e.Emote == nil {
			// The edge emote was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: emote.Label}
		}
		return e.Emote, nil
	}
	return nil, &NotLoadedError{edge: "emote"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldSize:
			values[i] = &sql.NullInt64{}
		case file.FieldID, file.FieldName, file.FieldContenttype:
			values[i] = &sql.NullString{}
		case file.ForeignKeys[0]: // emote_file
			values[i] = &sql.NullInt64{}
		case file.ForeignKeys[1]: // file_hash_file
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type File", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case file.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case file.FieldContenttype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contenttype", values[i])
			} else if value.Valid {
				f.Contenttype = value.String
			}
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = int(value.Int64)
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field emote_file", value)
			} else if value.Valid {
				f.emote_file = new(int)
				*f.emote_file = int(value.Int64)
			}
		case file.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field file_hash_file", value)
			} else if value.Valid {
				f.file_hash_file = new(int)
				*f.file_hash_file = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryFilehash queries the "filehash" edge of the File entity.
func (f *File) QueryFilehash() *FileHashQuery {
	return (&FileClient{config: f.config}).QueryFilehash(f)
}

// QueryEmote queries the "emote" edge of the File entity.
func (f *File) QueryEmote() *EmoteQuery {
	return (&FileClient{config: f.config}).QueryEmote(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("entgen: File is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", name=")
	builder.WriteString(f.Name)
	builder.WriteString(", contenttype=")
	builder.WriteString(f.Contenttype)
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
