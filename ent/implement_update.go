// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-farms/inventory/ent/category"
	"github.com/open-farms/inventory/ent/implement"
	"github.com/open-farms/inventory/ent/location"
	"github.com/open-farms/inventory/ent/predicate"
)

// ImplementUpdate is the builder for updating Implement entities.
type ImplementUpdate struct {
	config
	hooks    []Hook
	mutation *ImplementMutation
}

// Where appends a list predicates to the ImplementUpdate builder.
func (iu *ImplementUpdate) Where(ps ...predicate.Implement) *ImplementUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetCreateTime sets the "create_time" field.
func (iu *ImplementUpdate) SetCreateTime(t time.Time) *ImplementUpdate {
	iu.mutation.SetCreateTime(t)
	return iu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (iu *ImplementUpdate) SetNillableCreateTime(t *time.Time) *ImplementUpdate {
	if t != nil {
		iu.SetCreateTime(*t)
	}
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *ImplementUpdate) SetUpdateTime(t time.Time) *ImplementUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetName sets the "name" field.
func (iu *ImplementUpdate) SetName(s string) *ImplementUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (iu *ImplementUpdate) SetLocationID(id int) *ImplementUpdate {
	iu.mutation.SetLocationID(id)
	return iu
}

// SetLocation sets the "location" edge to the Location entity.
func (iu *ImplementUpdate) SetLocation(l *Location) *ImplementUpdate {
	return iu.SetLocationID(l.ID)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (iu *ImplementUpdate) SetCategoryID(id int) *ImplementUpdate {
	iu.mutation.SetCategoryID(id)
	return iu
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (iu *ImplementUpdate) SetNillableCategoryID(id *int) *ImplementUpdate {
	if id != nil {
		iu = iu.SetCategoryID(*id)
	}
	return iu
}

// SetCategory sets the "category" edge to the Category entity.
func (iu *ImplementUpdate) SetCategory(c *Category) *ImplementUpdate {
	return iu.SetCategoryID(c.ID)
}

// Mutation returns the ImplementMutation object of the builder.
func (iu *ImplementUpdate) Mutation() *ImplementMutation {
	return iu.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (iu *ImplementUpdate) ClearLocation() *ImplementUpdate {
	iu.mutation.ClearLocation()
	return iu
}

// ClearCategory clears the "category" edge to the Category entity.
func (iu *ImplementUpdate) ClearCategory() *ImplementUpdate {
	iu.mutation.ClearCategory()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImplementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImplementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImplementUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImplementUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImplementUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ImplementUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := implement.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImplementUpdate) check() error {
	if _, ok := iu.mutation.LocationID(); iu.mutation.LocationCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"location\"")
	}
	return nil
}

func (iu *ImplementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   implement.Table,
			Columns: implement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: implement.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: implement.FieldCreateTime,
		})
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: implement.FieldUpdateTime,
		})
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: implement.FieldName,
		})
	}
	if iu.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   implement.LocationTable,
			Columns: []string{implement.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   implement.LocationTable,
			Columns: []string{implement.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   implement.CategoryTable,
			Columns: []string{implement.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   implement.CategoryTable,
			Columns: []string{implement.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{implement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ImplementUpdateOne is the builder for updating a single Implement entity.
type ImplementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImplementMutation
}

// SetCreateTime sets the "create_time" field.
func (iuo *ImplementUpdateOne) SetCreateTime(t time.Time) *ImplementUpdateOne {
	iuo.mutation.SetCreateTime(t)
	return iuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (iuo *ImplementUpdateOne) SetNillableCreateTime(t *time.Time) *ImplementUpdateOne {
	if t != nil {
		iuo.SetCreateTime(*t)
	}
	return iuo
}

// SetUpdateTime sets the "update_time" field.
func (iuo *ImplementUpdateOne) SetUpdateTime(t time.Time) *ImplementUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetName sets the "name" field.
func (iuo *ImplementUpdateOne) SetName(s string) *ImplementUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (iuo *ImplementUpdateOne) SetLocationID(id int) *ImplementUpdateOne {
	iuo.mutation.SetLocationID(id)
	return iuo
}

// SetLocation sets the "location" edge to the Location entity.
func (iuo *ImplementUpdateOne) SetLocation(l *Location) *ImplementUpdateOne {
	return iuo.SetLocationID(l.ID)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (iuo *ImplementUpdateOne) SetCategoryID(id int) *ImplementUpdateOne {
	iuo.mutation.SetCategoryID(id)
	return iuo
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (iuo *ImplementUpdateOne) SetNillableCategoryID(id *int) *ImplementUpdateOne {
	if id != nil {
		iuo = iuo.SetCategoryID(*id)
	}
	return iuo
}

// SetCategory sets the "category" edge to the Category entity.
func (iuo *ImplementUpdateOne) SetCategory(c *Category) *ImplementUpdateOne {
	return iuo.SetCategoryID(c.ID)
}

// Mutation returns the ImplementMutation object of the builder.
func (iuo *ImplementUpdateOne) Mutation() *ImplementMutation {
	return iuo.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (iuo *ImplementUpdateOne) ClearLocation() *ImplementUpdateOne {
	iuo.mutation.ClearLocation()
	return iuo
}

// ClearCategory clears the "category" edge to the Category entity.
func (iuo *ImplementUpdateOne) ClearCategory() *ImplementUpdateOne {
	iuo.mutation.ClearCategory()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImplementUpdateOne) Select(field string, fields ...string) *ImplementUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Implement entity.
func (iuo *ImplementUpdateOne) Save(ctx context.Context) (*Implement, error) {
	var (
		err  error
		node *Implement
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImplementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImplementUpdateOne) SaveX(ctx context.Context) *Implement {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImplementUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImplementUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ImplementUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := implement.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImplementUpdateOne) check() error {
	if _, ok := iuo.mutation.LocationID(); iuo.mutation.LocationCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"location\"")
	}
	return nil
}

func (iuo *ImplementUpdateOne) sqlSave(ctx context.Context) (_node *Implement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   implement.Table,
			Columns: implement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: implement.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Implement.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, implement.FieldID)
		for _, f := range fields {
			if !implement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != implement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: implement.FieldCreateTime,
		})
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: implement.FieldUpdateTime,
		})
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: implement.FieldName,
		})
	}
	if iuo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   implement.LocationTable,
			Columns: []string{implement.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   implement.LocationTable,
			Columns: []string{implement.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   implement.CategoryTable,
			Columns: []string{implement.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   implement.CategoryTable,
			Columns: []string{implement.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Implement{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{implement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
