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
	"github.com/thoohv5/template-grpc/internal/ent/phoneaccount"
	"github.com/thoohv5/template-grpc/internal/ent/predicate"
)

// PhoneAccountQuery is the builder for querying PhoneAccount entities.
type PhoneAccountQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.PhoneAccount
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (paq *PhoneAccountQuery) Where(ps ...predicate.PhoneAccount) *PhoneAccountQuery {
	paq.predicates = append(paq.predicates, ps...)
	return paq
}

// Limit adds a limit step to the query.
func (paq *PhoneAccountQuery) Limit(limit int) *PhoneAccountQuery {
	paq.limit = &limit
	return paq
}

// Offset adds an offset step to the query.
func (paq *PhoneAccountQuery) Offset(offset int) *PhoneAccountQuery {
	paq.offset = &offset
	return paq
}

// Order adds an order step to the query.
func (paq *PhoneAccountQuery) Order(o ...OrderFunc) *PhoneAccountQuery {
	paq.order = append(paq.order, o...)
	return paq
}

// First returns the first PhoneAccount entity in the query. Returns *NotFoundError when no phoneaccount was found.
func (paq *PhoneAccountQuery) First(ctx context.Context) (*PhoneAccount, error) {
	nodes, err := paq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{phoneaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (paq *PhoneAccountQuery) FirstX(ctx context.Context) *PhoneAccount {
	node, err := paq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PhoneAccount id in the query. Returns *NotFoundError when no id was found.
func (paq *PhoneAccountQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = paq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{phoneaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (paq *PhoneAccountQuery) FirstXID(ctx context.Context) int64 {
	id, err := paq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only PhoneAccount entity in the query, returns an error if not exactly one entity was returned.
func (paq *PhoneAccountQuery) Only(ctx context.Context) (*PhoneAccount, error) {
	nodes, err := paq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{phoneaccount.Label}
	default:
		return nil, &NotSingularError{phoneaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (paq *PhoneAccountQuery) OnlyX(ctx context.Context) *PhoneAccount {
	node, err := paq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only PhoneAccount id in the query, returns an error if not exactly one id was returned.
func (paq *PhoneAccountQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = paq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = &NotSingularError{phoneaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (paq *PhoneAccountQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := paq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PhoneAccounts.
func (paq *PhoneAccountQuery) All(ctx context.Context) ([]*PhoneAccount, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return paq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (paq *PhoneAccountQuery) AllX(ctx context.Context) []*PhoneAccount {
	nodes, err := paq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PhoneAccount ids.
func (paq *PhoneAccountQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := paq.Select(phoneaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (paq *PhoneAccountQuery) IDsX(ctx context.Context) []int64 {
	ids, err := paq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (paq *PhoneAccountQuery) Count(ctx context.Context) (int, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return paq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (paq *PhoneAccountQuery) CountX(ctx context.Context) int {
	count, err := paq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (paq *PhoneAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return paq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (paq *PhoneAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := paq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (paq *PhoneAccountQuery) Clone() *PhoneAccountQuery {
	return &PhoneAccountQuery{
		config:     paq.config,
		limit:      paq.limit,
		offset:     paq.offset,
		order:      append([]OrderFunc{}, paq.order...),
		unique:     append([]string{}, paq.unique...),
		predicates: append([]predicate.PhoneAccount{}, paq.predicates...),
		// clone intermediate query.
		sql:  paq.sql.Clone(),
		path: paq.path,
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
//	client.PhoneAccount.Query().
//		GroupBy(phoneaccount.FieldUserIdentity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (paq *PhoneAccountQuery) GroupBy(field string, fields ...string) *PhoneAccountGroupBy {
	group := &PhoneAccountGroupBy{config: paq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return paq.sqlQuery(), nil
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
//	client.PhoneAccount.Query().
//		Select(phoneaccount.FieldUserIdentity).
//		Scan(ctx, &v)
//
func (paq *PhoneAccountQuery) Select(field string, fields ...string) *PhoneAccountSelect {
	selector := &PhoneAccountSelect{config: paq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return paq.sqlQuery(), nil
	}
	return selector
}

func (paq *PhoneAccountQuery) prepareQuery(ctx context.Context) error {
	if paq.path != nil {
		prev, err := paq.path(ctx)
		if err != nil {
			return err
		}
		paq.sql = prev
	}
	return nil
}

func (paq *PhoneAccountQuery) sqlAll(ctx context.Context) ([]*PhoneAccount, error) {
	var (
		nodes = []*PhoneAccount{}
		_spec = paq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &PhoneAccount{config: paq.config}
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
	if err := sqlgraph.QueryNodes(ctx, paq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (paq *PhoneAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := paq.querySpec()
	return sqlgraph.CountNodes(ctx, paq.driver, _spec)
}

func (paq *PhoneAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := paq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (paq *PhoneAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   phoneaccount.Table,
			Columns: phoneaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: phoneaccount.FieldID,
			},
		},
		From:   paq.sql,
		Unique: true,
	}
	if ps := paq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := paq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := paq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := paq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, phoneaccount.ValidColumn)
			}
		}
	}
	return _spec
}

func (paq *PhoneAccountQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(paq.driver.Dialect())
	t1 := builder.Table(phoneaccount.Table)
	selector := builder.Select(t1.Columns(phoneaccount.Columns...)...).From(t1)
	if paq.sql != nil {
		selector = paq.sql
		selector.Select(selector.Columns(phoneaccount.Columns...)...)
	}
	for _, p := range paq.predicates {
		p(selector)
	}
	for _, p := range paq.order {
		p(selector, phoneaccount.ValidColumn)
	}
	if offset := paq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := paq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PhoneAccountGroupBy is the builder for group-by PhoneAccount entities.
type PhoneAccountGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pagb *PhoneAccountGroupBy) Aggregate(fns ...AggregateFunc) *PhoneAccountGroupBy {
	pagb.fns = append(pagb.fns, fns...)
	return pagb
}

// Scan applies the group-by query and scan the result into the given value.
func (pagb *PhoneAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pagb.path(ctx)
	if err != nil {
		return err
	}
	pagb.sql = query
	return pagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pagb *PhoneAccountGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: PhoneAccountGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) StringsX(ctx context.Context) []string {
	v, err := pagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pagb *PhoneAccountGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = fmt.Errorf("ent: PhoneAccountGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) StringX(ctx context.Context) string {
	v, err := pagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pagb *PhoneAccountGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: PhoneAccountGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) IntsX(ctx context.Context) []int {
	v, err := pagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pagb *PhoneAccountGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = fmt.Errorf("ent: PhoneAccountGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) IntX(ctx context.Context) int {
	v, err := pagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pagb *PhoneAccountGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: PhoneAccountGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pagb *PhoneAccountGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = fmt.Errorf("ent: PhoneAccountGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pagb *PhoneAccountGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: PhoneAccountGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pagb *PhoneAccountGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = fmt.Errorf("ent: PhoneAccountGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pagb *PhoneAccountGroupBy) BoolX(ctx context.Context) bool {
	v, err := pagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pagb *PhoneAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pagb.fields {
		if !phoneaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pagb *PhoneAccountGroupBy) sqlQuery() *sql.Selector {
	selector := pagb.sql
	columns := make([]string, 0, len(pagb.fields)+len(pagb.fns))
	columns = append(columns, pagb.fields...)
	for _, fn := range pagb.fns {
		columns = append(columns, fn(selector, phoneaccount.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(pagb.fields...)
}

// PhoneAccountSelect is the builder for select fields of PhoneAccount entities.
type PhoneAccountSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (pas *PhoneAccountSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := pas.path(ctx)
	if err != nil {
		return err
	}
	pas.sql = query
	return pas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pas *PhoneAccountSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (pas *PhoneAccountSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: PhoneAccountSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pas *PhoneAccountSelect) StringsX(ctx context.Context) []string {
	v, err := pas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (pas *PhoneAccountSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = fmt.Errorf("ent: PhoneAccountSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pas *PhoneAccountSelect) StringX(ctx context.Context) string {
	v, err := pas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (pas *PhoneAccountSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: PhoneAccountSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pas *PhoneAccountSelect) IntsX(ctx context.Context) []int {
	v, err := pas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (pas *PhoneAccountSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = fmt.Errorf("ent: PhoneAccountSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pas *PhoneAccountSelect) IntX(ctx context.Context) int {
	v, err := pas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (pas *PhoneAccountSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: PhoneAccountSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pas *PhoneAccountSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (pas *PhoneAccountSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = fmt.Errorf("ent: PhoneAccountSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pas *PhoneAccountSelect) Float64X(ctx context.Context) float64 {
	v, err := pas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (pas *PhoneAccountSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: PhoneAccountSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pas *PhoneAccountSelect) BoolsX(ctx context.Context) []bool {
	v, err := pas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (pas *PhoneAccountSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phoneaccount.Label}
	default:
		err = fmt.Errorf("ent: PhoneAccountSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pas *PhoneAccountSelect) BoolX(ctx context.Context) bool {
	v, err := pas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pas *PhoneAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pas.fields {
		if !phoneaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := pas.sqlQuery().Query()
	if err := pas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pas *PhoneAccountSelect) sqlQuery() sql.Querier {
	selector := pas.sql
	selector.Select(selector.Columns(pas.fields...)...)
	return selector
}
