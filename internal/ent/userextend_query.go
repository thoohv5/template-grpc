// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/thoohv5/template-grpc/internal/ent/predicate"
	"github.com/thoohv5/template-grpc/internal/ent/userextend"
)

// UserExtendQuery is the builder for querying UserExtend entities.
type UserExtendQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.UserExtend
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ueq *UserExtendQuery) Where(ps ...predicate.UserExtend) *UserExtendQuery {
	ueq.predicates = append(ueq.predicates, ps...)
	return ueq
}

// Limit adds a limit step to the query.
func (ueq *UserExtendQuery) Limit(limit int) *UserExtendQuery {
	ueq.limit = &limit
	return ueq
}

// Offset adds an offset step to the query.
func (ueq *UserExtendQuery) Offset(offset int) *UserExtendQuery {
	ueq.offset = &offset
	return ueq
}

// Order adds an order step to the query.
func (ueq *UserExtendQuery) Order(o ...OrderFunc) *UserExtendQuery {
	ueq.order = append(ueq.order, o...)
	return ueq
}

// First returns the first UserExtend entity in the query. Returns *NotFoundError when no userextend was found.
func (ueq *UserExtendQuery) First(ctx context.Context) (*UserExtend, error) {
	nodes, err := ueq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userextend.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ueq *UserExtendQuery) FirstX(ctx context.Context) *UserExtend {
	node, err := ueq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserExtend id in the query. Returns *NotFoundError when no id was found.
func (ueq *UserExtendQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ueq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userextend.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ueq *UserExtendQuery) FirstXID(ctx context.Context) int64 {
	id, err := ueq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only UserExtend entity in the query, returns an error if not exactly one entity was returned.
func (ueq *UserExtendQuery) Only(ctx context.Context) (*UserExtend, error) {
	nodes, err := ueq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userextend.Label}
	default:
		return nil, &NotSingularError{userextend.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ueq *UserExtendQuery) OnlyX(ctx context.Context) *UserExtend {
	node, err := ueq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only UserExtend id in the query, returns an error if not exactly one id was returned.
func (ueq *UserExtendQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ueq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = &NotSingularError{userextend.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ueq *UserExtendQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ueq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserExtends.
func (ueq *UserExtendQuery) All(ctx context.Context) ([]*UserExtend, error) {
	if err := ueq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ueq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ueq *UserExtendQuery) AllX(ctx context.Context) []*UserExtend {
	nodes, err := ueq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserExtend ids.
func (ueq *UserExtendQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := ueq.Select(userextend.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ueq *UserExtendQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ueq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ueq *UserExtendQuery) Count(ctx context.Context) (int, error) {
	if err := ueq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ueq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ueq *UserExtendQuery) CountX(ctx context.Context) int {
	count, err := ueq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ueq *UserExtendQuery) Exist(ctx context.Context) (bool, error) {
	if err := ueq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ueq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ueq *UserExtendQuery) ExistX(ctx context.Context) bool {
	exist, err := ueq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ueq *UserExtendQuery) Clone() *UserExtendQuery {
	return &UserExtendQuery{
		config:     ueq.config,
		limit:      ueq.limit,
		offset:     ueq.offset,
		order:      append([]OrderFunc{}, ueq.order...),
		unique:     append([]string{}, ueq.unique...),
		predicates: append([]predicate.UserExtend{}, ueq.predicates...),
		// clone intermediate query.
		sql:  ueq.sql.Clone(),
		path: ueq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// PostExample:
//
//	var v []struct {
//		UserIdentity string `json:"user_identity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserExtend.Query().
//		GroupBy(userextend.FieldUserIdentity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ueq *UserExtendQuery) GroupBy(field string, fields ...string) *UserExtendGroupBy {
	group := &UserExtendGroupBy{config: ueq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ueq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ueq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// PostExample:
//
//	var v []struct {
//		UserIdentity string `json:"user_identity,omitempty"`
//	}
//
//	client.UserExtend.Query().
//		Select(userextend.FieldUserIdentity).
//		Scan(ctx, &v)
//
func (ueq *UserExtendQuery) Select(field string, fields ...string) *UserExtendSelect {
	selector := &UserExtendSelect{config: ueq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ueq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ueq.sqlQuery(), nil
	}
	return selector
}

func (ueq *UserExtendQuery) prepareQuery(ctx context.Context) error {
	if ueq.path != nil {
		prev, err := ueq.path(ctx)
		if err != nil {
			return err
		}
		ueq.sql = prev
	}
	return nil
}

func (ueq *UserExtendQuery) sqlAll(ctx context.Context) ([]*UserExtend, error) {
	var (
		nodes = []*UserExtend{}
		_spec = ueq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &UserExtend{config: ueq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ueq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ueq *UserExtendQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ueq.querySpec()
	return sqlgraph.CountNodes(ctx, ueq.driver, _spec)
}

func (ueq *UserExtendQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ueq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ueq *UserExtendQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userextend.Table,
			Columns: userextend.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userextend.FieldID,
			},
		},
		From:   ueq.sql,
		Unique: true,
	}
	if ps := ueq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ueq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ueq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ueq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, userextend.ValidColumn)
			}
		}
	}
	return _spec
}

func (ueq *UserExtendQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ueq.driver.Dialect())
	t1 := builder.Table(userextend.Table)
	selector := builder.Select(t1.Columns(userextend.Columns...)...).From(t1)
	if ueq.sql != nil {
		selector = ueq.sql
		selector.Select(selector.Columns(userextend.Columns...)...)
	}
	for _, p := range ueq.predicates {
		p(selector)
	}
	for _, p := range ueq.order {
		p(selector, userextend.ValidColumn)
	}
	if offset := ueq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ueq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserExtendGroupBy is the builder for group-by UserExtend entities.
type UserExtendGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uegb *UserExtendGroupBy) Aggregate(fns ...AggregateFunc) *UserExtendGroupBy {
	uegb.fns = append(uegb.fns, fns...)
	return uegb
}

// Scan applies the group-by query and scan the result into the given value.
func (uegb *UserExtendGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uegb.path(ctx)
	if err != nil {
		return err
	}
	uegb.sql = query
	return uegb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uegb *UserExtendGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uegb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (uegb *UserExtendGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uegb.fields) > 1 {
		return nil, errors.New("ent: UserExtendGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uegb *UserExtendGroupBy) StringsX(ctx context.Context) []string {
	v, err := uegb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (uegb *UserExtendGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uegb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = fmt.Errorf("ent: UserExtendGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uegb *UserExtendGroupBy) StringX(ctx context.Context) string {
	v, err := uegb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (uegb *UserExtendGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uegb.fields) > 1 {
		return nil, errors.New("ent: UserExtendGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uegb *UserExtendGroupBy) IntsX(ctx context.Context) []int {
	v, err := uegb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (uegb *UserExtendGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uegb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = fmt.Errorf("ent: UserExtendGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uegb *UserExtendGroupBy) IntX(ctx context.Context) int {
	v, err := uegb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (uegb *UserExtendGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uegb.fields) > 1 {
		return nil, errors.New("ent: UserExtendGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uegb *UserExtendGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uegb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (uegb *UserExtendGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uegb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = fmt.Errorf("ent: UserExtendGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uegb *UserExtendGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uegb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (uegb *UserExtendGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uegb.fields) > 1 {
		return nil, errors.New("ent: UserExtendGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uegb *UserExtendGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uegb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (uegb *UserExtendGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uegb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = fmt.Errorf("ent: UserExtendGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uegb *UserExtendGroupBy) BoolX(ctx context.Context) bool {
	v, err := uegb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uegb *UserExtendGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uegb.fields {
		if !userextend.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uegb *UserExtendGroupBy) sqlQuery() *sql.Selector {
	selector := uegb.sql
	columns := make([]string, 0, len(uegb.fields)+len(uegb.fns))
	columns = append(columns, uegb.fields...)
	for _, fn := range uegb.fns {
		columns = append(columns, fn(selector, userextend.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(uegb.fields...)
}

// UserExtendSelect is the builder for select fields of UserExtend entities.
type UserExtendSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ues *UserExtendSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ues.path(ctx)
	if err != nil {
		return err
	}
	ues.sql = query
	return ues.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ues *UserExtendSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ues.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ues *UserExtendSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ues.fields) > 1 {
		return nil, errors.New("ent: UserExtendSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ues.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ues *UserExtendSelect) StringsX(ctx context.Context) []string {
	v, err := ues.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ues *UserExtendSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ues.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = fmt.Errorf("ent: UserExtendSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ues *UserExtendSelect) StringX(ctx context.Context) string {
	v, err := ues.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ues *UserExtendSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ues.fields) > 1 {
		return nil, errors.New("ent: UserExtendSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ues.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ues *UserExtendSelect) IntsX(ctx context.Context) []int {
	v, err := ues.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ues *UserExtendSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ues.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = fmt.Errorf("ent: UserExtendSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ues *UserExtendSelect) IntX(ctx context.Context) int {
	v, err := ues.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ues *UserExtendSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ues.fields) > 1 {
		return nil, errors.New("ent: UserExtendSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ues.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ues *UserExtendSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ues.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ues *UserExtendSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ues.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = fmt.Errorf("ent: UserExtendSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ues *UserExtendSelect) Float64X(ctx context.Context) float64 {
	v, err := ues.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ues *UserExtendSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ues.fields) > 1 {
		return nil, errors.New("ent: UserExtendSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ues.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ues *UserExtendSelect) BoolsX(ctx context.Context) []bool {
	v, err := ues.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ues *UserExtendSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ues.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userextend.Label}
	default:
		err = fmt.Errorf("ent: UserExtendSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ues *UserExtendSelect) BoolX(ctx context.Context) bool {
	v, err := ues.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ues *UserExtendSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ues.fields {
		if !userextend.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := ues.sqlQuery().Query()
	if err := ues.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ues *UserExtendSelect) sqlQuery() sql.Querier {
	selector := ues.sql
	selector.Select(selector.Columns(ues.fields...)...)
	return selector
}
