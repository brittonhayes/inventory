// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/open-farms/inventory/ent/category"
	"github.com/open-farms/inventory/ent/location"
)

// Location is the model entity for the Location schema.
type Location struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Zone holds the value of the "zone" field.
	Zone int32 `json:"zone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LocationQuery when eager-loading is set.
	Edges             LocationEdges `json:"edges"`
	category_location *int
}

// LocationEdges holds the relations/edges for other nodes in the graph.
type LocationEdges struct {
	// Vehicle holds the value of the vehicle edge.
	Vehicle []*Vehicle `json:"vehicle,omitempty"`
	// Tool holds the value of the tool edge.
	Tool []*Tool `json:"tool,omitempty"`
	// Implement holds the value of the implement edge.
	Implement []*Implement `json:"implement,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment []*Equipment `json:"equipment,omitempty"`
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// VehicleOrErr returns the Vehicle value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) VehicleOrErr() ([]*Vehicle, error) {
	if e.loadedTypes[0] {
		return e.Vehicle, nil
	}
	return nil, &NotLoadedError{edge: "vehicle"}
}

// ToolOrErr returns the Tool value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) ToolOrErr() ([]*Tool, error) {
	if e.loadedTypes[1] {
		return e.Tool, nil
	}
	return nil, &NotLoadedError{edge: "tool"}
}

// ImplementOrErr returns the Implement value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) ImplementOrErr() ([]*Implement, error) {
	if e.loadedTypes[2] {
		return e.Implement, nil
	}
	return nil, &NotLoadedError{edge: "implement"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[3] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[4] {
		if e.Category == nil {
			// The edge category was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Location) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case location.FieldID, location.FieldZone:
			values[i] = new(sql.NullInt64)
		case location.FieldName:
			values[i] = new(sql.NullString)
		case location.FieldCreateTime, location.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case location.ForeignKeys[0]: // category_location
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Location", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Location fields.
func (l *Location) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case location.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case location.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				l.CreateTime = value.Time
			}
		case location.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				l.UpdateTime = value.Time
			}
		case location.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case location.FieldZone:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field zone", values[i])
			} else if value.Valid {
				l.Zone = int32(value.Int64)
			}
		case location.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_location", value)
			} else if value.Valid {
				l.category_location = new(int)
				*l.category_location = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryVehicle queries the "vehicle" edge of the Location entity.
func (l *Location) QueryVehicle() *VehicleQuery {
	return (&LocationClient{config: l.config}).QueryVehicle(l)
}

// QueryTool queries the "tool" edge of the Location entity.
func (l *Location) QueryTool() *ToolQuery {
	return (&LocationClient{config: l.config}).QueryTool(l)
}

// QueryImplement queries the "implement" edge of the Location entity.
func (l *Location) QueryImplement() *ImplementQuery {
	return (&LocationClient{config: l.config}).QueryImplement(l)
}

// QueryEquipment queries the "equipment" edge of the Location entity.
func (l *Location) QueryEquipment() *EquipmentQuery {
	return (&LocationClient{config: l.config}).QueryEquipment(l)
}

// QueryCategory queries the "category" edge of the Location entity.
func (l *Location) QueryCategory() *CategoryQuery {
	return (&LocationClient{config: l.config}).QueryCategory(l)
}

// Update returns a builder for updating this Location.
// Note that you need to call Location.Unwrap() before calling this method if this Location
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Location) Update() *LocationUpdateOne {
	return (&LocationClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Location entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Location) Unwrap() *Location {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Location is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Location) String() string {
	var builder strings.Builder
	builder.WriteString("Location(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(l.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(l.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(l.Name)
	builder.WriteString(", zone=")
	builder.WriteString(fmt.Sprintf("%v", l.Zone))
	builder.WriteByte(')')
	return builder.String()
}

// Locations is a parsable slice of Location.
type Locations []*Location

func (l Locations) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}