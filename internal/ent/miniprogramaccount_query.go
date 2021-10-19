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
	"github.com/thoohv5/template/internal/ent/miniprogramaccount"
	"github.com/thoohv5/template/internal/ent/predicate"
)

// MiniProgramAccountQuery is the builder for querying MiniProgramAccount entities.
type MiniProgramAccountQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.MiniProgramAccount
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (mpaq *MiniProgramAccountQuery) Where(ps ...predicate.MiniProgramAccount) *MiniProgramAccountQuery {
	mpaq.predicates = append(mpaq.predicates, ps...)
	return mpaq
}

// Limit adds a limit step to the query.
func (mpaq *MiniProgramAccountQuery) Limit(limit int) *MiniProgramAccountQuery {
	mpaq.limit = &limit
	return mpaq
}

// Offset adds an offset step to the query.
func (mpaq *MiniProgramAccountQuery) Offset(offset int) *MiniProgramAccountQuery {
	mpaq.offset = &offset
	return mpaq
}

// Order adds an order step to the query.
func (mpaq *MiniProgramAccountQuery) Order(o ...OrderFunc) *MiniProgramAccountQuery {
	mpaq.order = append(mpaq.order, o...)
	return mpaq
}

// First returns the first MiniProgramAccount entity in the query. Returns *NotFoundError when no miniprogramaccount was found.
func (mpaq *MiniProgramAccountQuery) First(ctx context.Context) (*MiniProgramAccount, error) {
	nodes, err := mpaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{miniprogramaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpaq *MiniProgramAccountQuery) FirstX(ctx context.Context) *MiniProgramAccount {
	node, err := mpaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MiniProgramAccount id in the query. Returns *NotFoundError when no id was found.
func (mpaq *MiniProgramAccountQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{miniprogramaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mpaq *MiniProgramAccountQuery) FirstXID(ctx context.Context) int64 {
	id, err := mpaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only MiniProgramAccount entity in the query, returns an error if not exactly one entity was returned.
func (mpaq *MiniProgramAccountQuery) Only(ctx context.Context) (*MiniProgramAccount, error) {
	nodes, err := mpaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{miniprogramaccount.Label}
	default:
		return nil, &NotSingularError{miniprogramaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpaq *MiniProgramAccountQuery) OnlyX(ctx context.Context) *MiniProgramAccount {
	node, err := mpaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only MiniProgramAccount id in the query, returns an error if not exactly one id was returned.
func (mpaq *MiniProgramAccountQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = &NotSingularError{miniprogramaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpaq *MiniProgramAccountQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mpaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MiniProgramAccounts.
func (mpaq *MiniProgramAccountQuery) All(ctx context.Context) ([]*MiniProgramAccount, error) {
	if err := mpaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mpaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mpaq *MiniProgramAccountQuery) AllX(ctx context.Context) []*MiniProgramAccount {
	nodes, err := mpaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MiniProgramAccount ids.
func (mpaq *MiniProgramAccountQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := mpaq.Select(miniprogramaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpaq *MiniProgramAccountQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mpaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpaq *MiniProgramAccountQuery) Count(ctx context.Context) (int, error) {
	if err := mpaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mpaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mpaq *MiniProgramAccountQuery) CountX(ctx context.Context) int {
	count, err := mpaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpaq *MiniProgramAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := mpaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mpaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mpaq *MiniProgramAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := mpaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpaq *MiniProgramAccountQuery) Clone() *MiniProgramAccountQuery {
	return &MiniProgramAccountQuery{
		config:     mpaq.config,
		limit:      mpaq.limit,
		offset:     mpaq.offset,
		order:      append([]OrderFunc{}, mpaq.order...),
		unique:     append([]string{}, mpaq.unique...),
		predicates: append([]predicate.MiniProgramAccount{}, mpaq.predicates...),
		// clone intermediate query.
		sql:  mpaq.sql.Clone(),
		path: mpaq.path,
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
//	client.MiniProgramAccount.Query().
//		GroupBy(miniprogramaccount.FieldUserIdentity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mpaq *MiniProgramAccountQuery) GroupBy(field string, fields ...string) *MiniProgramAccountGroupBy {
	group := &MiniProgramAccountGroupBy{config: mpaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mpaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mpaq.sqlQuery(), nil
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
//	client.MiniProgramAccount.Query().
//		Select(miniprogramaccount.FieldUserIdentity).
//		Scan(ctx, &v)
//
func (mpaq *MiniProgramAccountQuery) Select(field string, fields ...string) *MiniProgramAccountSelect {
	selector := &MiniProgramAccountSelect{config: mpaq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mpaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mpaq.sqlQuery(), nil
	}
	return selector
}

func (mpaq *MiniProgramAccountQuery) prepareQuery(ctx context.Context) error {
	if mpaq.path != nil {
		prev, err := mpaq.path(ctx)
		if err != nil {
			return err
		}
		mpaq.sql = prev
	}
	return nil
}

func (mpaq *MiniProgramAccountQuery) sqlAll(ctx context.Context) ([]*MiniProgramAccount, error) {
	var (
		nodes = []*MiniProgramAccount{}
		_spec = mpaq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &MiniProgramAccount{config: mpaq.config}
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
	if err := sqlgraph.QueryNodes(ctx, mpaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mpaq *MiniProgramAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpaq.querySpec()
	return sqlgraph.CountNodes(ctx, mpaq.driver, _spec)
}

func (mpaq *MiniProgramAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mpaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mpaq *MiniProgramAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   miniprogramaccount.Table,
			Columns: miniprogramaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: miniprogramaccount.FieldID,
			},
		},
		From:   mpaq.sql,
		Unique: true,
	}
	if ps := mpaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mpaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mpaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mpaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, miniprogramaccount.ValidColumn)
			}
		}
	}
	return _spec
}

func (mpaq *MiniProgramAccountQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mpaq.driver.Dialect())
	t1 := builder.Table(miniprogramaccount.Table)
	selector := builder.Select(t1.Columns(miniprogramaccount.Columns...)...).From(t1)
	if mpaq.sql != nil {
		selector = mpaq.sql
		selector.Select(selector.Columns(miniprogramaccount.Columns...)...)
	}
	for _, p := range mpaq.predicates {
		p(selector)
	}
	for _, p := range mpaq.order {
		p(selector, miniprogramaccount.ValidColumn)
	}
	if offset := mpaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mpaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MiniProgramAccountGroupBy is the builder for group-by MiniProgramAccount entities.
type MiniProgramAccountGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpagb *MiniProgramAccountGroupBy) Aggregate(fns ...AggregateFunc) *MiniProgramAccountGroupBy {
	mpagb.fns = append(mpagb.fns, fns...)
	return mpagb
}

// Scan applies the group-by query and scan the result into the given value.
func (mpagb *MiniProgramAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mpagb.path(ctx)
	if err != nil {
		return err
	}
	mpagb.sql = query
	return mpagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mpagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mpagb *MiniProgramAccountGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mpagb.fields) > 1 {
		return nil, errors.New("ent: MiniProgramAccountGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mpagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) StringsX(ctx context.Context) []string {
	v, err := mpagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (mpagb *MiniProgramAccountGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mpagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = fmt.Errorf("ent: MiniProgramAccountGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) StringX(ctx context.Context) string {
	v, err := mpagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mpagb *MiniProgramAccountGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mpagb.fields) > 1 {
		return nil, errors.New("ent: MiniProgramAccountGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mpagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) IntsX(ctx context.Context) []int {
	v, err := mpagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (mpagb *MiniProgramAccountGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mpagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = fmt.Errorf("ent: MiniProgramAccountGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) IntX(ctx context.Context) int {
	v, err := mpagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mpagb *MiniProgramAccountGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mpagb.fields) > 1 {
		return nil, errors.New("ent: MiniProgramAccountGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mpagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mpagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (mpagb *MiniProgramAccountGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mpagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = fmt.Errorf("ent: MiniProgramAccountGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mpagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mpagb *MiniProgramAccountGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mpagb.fields) > 1 {
		return nil, errors.New("ent: MiniProgramAccountGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mpagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mpagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (mpagb *MiniProgramAccountGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mpagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = fmt.Errorf("ent: MiniProgramAccountGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mpagb *MiniProgramAccountGroupBy) BoolX(ctx context.Context) bool {
	v, err := mpagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mpagb *MiniProgramAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mpagb.fields {
		if !miniprogramaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mpagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mpagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mpagb *MiniProgramAccountGroupBy) sqlQuery() *sql.Selector {
	selector := mpagb.sql
	columns := make([]string, 0, len(mpagb.fields)+len(mpagb.fns))
	columns = append(columns, mpagb.fields...)
	for _, fn := range mpagb.fns {
		columns = append(columns, fn(selector, miniprogramaccount.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(mpagb.fields...)
}

// MiniProgramAccountSelect is the builder for select fields of MiniProgramAccount entities.
type MiniProgramAccountSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (mpas *MiniProgramAccountSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := mpas.path(ctx)
	if err != nil {
		return err
	}
	mpas.sql = query
	return mpas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mpas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (mpas *MiniProgramAccountSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mpas.fields) > 1 {
		return nil, errors.New("ent: MiniProgramAccountSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mpas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) StringsX(ctx context.Context) []string {
	v, err := mpas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (mpas *MiniProgramAccountSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mpas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = fmt.Errorf("ent: MiniProgramAccountSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) StringX(ctx context.Context) string {
	v, err := mpas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (mpas *MiniProgramAccountSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mpas.fields) > 1 {
		return nil, errors.New("ent: MiniProgramAccountSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mpas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) IntsX(ctx context.Context) []int {
	v, err := mpas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (mpas *MiniProgramAccountSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mpas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = fmt.Errorf("ent: MiniProgramAccountSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) IntX(ctx context.Context) int {
	v, err := mpas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (mpas *MiniProgramAccountSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mpas.fields) > 1 {
		return nil, errors.New("ent: MiniProgramAccountSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mpas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mpas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (mpas *MiniProgramAccountSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mpas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = fmt.Errorf("ent: MiniProgramAccountSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) Float64X(ctx context.Context) float64 {
	v, err := mpas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (mpas *MiniProgramAccountSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mpas.fields) > 1 {
		return nil, errors.New("ent: MiniProgramAccountSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mpas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) BoolsX(ctx context.Context) []bool {
	v, err := mpas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (mpas *MiniProgramAccountSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mpas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{miniprogramaccount.Label}
	default:
		err = fmt.Errorf("ent: MiniProgramAccountSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mpas *MiniProgramAccountSelect) BoolX(ctx context.Context) bool {
	v, err := mpas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mpas *MiniProgramAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mpas.fields {
		if !miniprogramaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := mpas.sqlQuery().Query()
	if err := mpas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mpas *MiniProgramAccountSelect) sqlQuery() sql.Querier {
	selector := mpas.sql
	selector.Select(selector.Columns(mpas.fields...)...)
	return selector
}
