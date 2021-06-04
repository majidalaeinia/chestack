// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Stack is an object representing the database table.
type Stack struct {
	ID        uint64    `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *stackR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stackL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StackColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var StackWhere = struct {
	ID        whereHelperuint64
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperuint64{field: "`stacks`.`id`"},
	CreatedAt: whereHelpernull_Time{field: "`stacks`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`stacks`.`updated_at`"},
}

// StackRels is where relationship names are stored.
var StackRels = struct {
}{}

// stackR is where relationships are stored.
type stackR struct {
}

// NewStruct creates a new relationship struct
func (*stackR) NewStruct() *stackR {
	return &stackR{}
}

// stackL is where Load methods for each relationship are stored.
type stackL struct{}

var (
	stackAllColumns            = []string{"id", "created_at", "updated_at"}
	stackColumnsWithoutDefault = []string{"created_at", "updated_at"}
	stackColumnsWithDefault    = []string{"id"}
	stackPrimaryKeyColumns     = []string{"id"}
)

type (
	// StackSlice is an alias for a slice of pointers to Stack.
	// This should generally be used opposed to []Stack.
	StackSlice []*Stack
	// StackHook is the signature for custom Stack hook methods
	StackHook func(boil.Executor, *Stack) error

	stackQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stackType                 = reflect.TypeOf(&Stack{})
	stackMapping              = queries.MakeStructMapping(stackType)
	stackPrimaryKeyMapping, _ = queries.BindMapping(stackType, stackMapping, stackPrimaryKeyColumns)
	stackInsertCacheMut       sync.RWMutex
	stackInsertCache          = make(map[string]insertCache)
	stackUpdateCacheMut       sync.RWMutex
	stackUpdateCache          = make(map[string]updateCache)
	stackUpsertCacheMut       sync.RWMutex
	stackUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var stackBeforeInsertHooks []StackHook
var stackBeforeUpdateHooks []StackHook
var stackBeforeDeleteHooks []StackHook
var stackBeforeUpsertHooks []StackHook

var stackAfterInsertHooks []StackHook
var stackAfterSelectHooks []StackHook
var stackAfterUpdateHooks []StackHook
var stackAfterDeleteHooks []StackHook
var stackAfterUpsertHooks []StackHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Stack) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stackBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Stack) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stackBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Stack) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stackBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Stack) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stackBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Stack) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stackAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Stack) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stackAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Stack) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stackAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Stack) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stackAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Stack) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stackAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStackHook registers your hook function for all future operations.
func AddStackHook(hookPoint boil.HookPoint, stackHook StackHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stackBeforeInsertHooks = append(stackBeforeInsertHooks, stackHook)
	case boil.BeforeUpdateHook:
		stackBeforeUpdateHooks = append(stackBeforeUpdateHooks, stackHook)
	case boil.BeforeDeleteHook:
		stackBeforeDeleteHooks = append(stackBeforeDeleteHooks, stackHook)
	case boil.BeforeUpsertHook:
		stackBeforeUpsertHooks = append(stackBeforeUpsertHooks, stackHook)
	case boil.AfterInsertHook:
		stackAfterInsertHooks = append(stackAfterInsertHooks, stackHook)
	case boil.AfterSelectHook:
		stackAfterSelectHooks = append(stackAfterSelectHooks, stackHook)
	case boil.AfterUpdateHook:
		stackAfterUpdateHooks = append(stackAfterUpdateHooks, stackHook)
	case boil.AfterDeleteHook:
		stackAfterDeleteHooks = append(stackAfterDeleteHooks, stackHook)
	case boil.AfterUpsertHook:
		stackAfterUpsertHooks = append(stackAfterUpsertHooks, stackHook)
	}
}

// OneG returns a single stack record from the query using the global executor.
func (q stackQuery) OneG() (*Stack, error) {
	return q.One(boil.GetDB())
}

// One returns a single stack record from the query.
func (q stackQuery) One(exec boil.Executor) (*Stack, error) {
	o := &Stack{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stacks")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Stack records from the query using the global executor.
func (q stackQuery) AllG() (StackSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Stack records from the query.
func (q stackQuery) All(exec boil.Executor) (StackSlice, error) {
	var o []*Stack

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stack slice")
	}

	if len(stackAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Stack records in the query, and panics on error.
func (q stackQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Stack records in the query.
func (q stackQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stacks rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q stackQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q stackQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stacks exists")
	}

	return count > 0, nil
}

// Stacks retrieves all the records using an executor.
func Stacks(mods ...qm.QueryMod) stackQuery {
	mods = append(mods, qm.From("`stacks`"))
	return stackQuery{NewQuery(mods...)}
}

// FindStackG retrieves a single record by ID.
func FindStackG(iD uint64, selectCols ...string) (*Stack, error) {
	return FindStack(boil.GetDB(), iD, selectCols...)
}

// FindStack retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStack(exec boil.Executor, iD uint64, selectCols ...string) (*Stack, error) {
	stackObj := &Stack{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `stacks` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, stackObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stacks")
	}

	return stackObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Stack) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Stack) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stacks provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stackColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stackInsertCacheMut.RLock()
	cache, cached := stackInsertCache[key]
	stackInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stackAllColumns,
			stackColumnsWithDefault,
			stackColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stackType, stackMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stackType, stackMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `stacks` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `stacks` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `stacks` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, stackPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into stacks")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == stackMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for stacks")
	}

CacheNoHooks:
	if !cached {
		stackInsertCacheMut.Lock()
		stackInsertCache[key] = cache
		stackInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Stack record using the global executor.
// See Update for more documentation.
func (o *Stack) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Stack.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Stack) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	stackUpdateCacheMut.RLock()
	cache, cached := stackUpdateCache[key]
	stackUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stackAllColumns,
			stackPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stacks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `stacks` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, stackPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stackType, stackMapping, append(wl, stackPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update stacks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stacks")
	}

	if !cached {
		stackUpdateCacheMut.Lock()
		stackUpdateCache[key] = cache
		stackUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q stackQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q stackQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stacks")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StackSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StackSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `stacks` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stackPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stack slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stack")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Stack) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

var mySQLStackUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Stack) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stacks provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stackColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLStackUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	stackUpsertCacheMut.RLock()
	cache, cached := stackUpsertCache[key]
	stackUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stackAllColumns,
			stackColumnsWithDefault,
			stackColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stackAllColumns,
			stackPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert stacks, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`stacks`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `stacks` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(stackType, stackMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stackType, stackMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for stacks")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == stackMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(stackType, stackMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for stacks")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for stacks")
	}

CacheNoHooks:
	if !cached {
		stackUpsertCacheMut.Lock()
		stackUpsertCache[key] = cache
		stackUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Stack record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Stack) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Stack record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stack) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Stack provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stackPrimaryKeyMapping)
	sql := "DELETE FROM `stacks` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stacks")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q stackQuery) DeleteAllG() (int64, error) {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q stackQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stackQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stacks")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o StackSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StackSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(stackBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `stacks` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stackPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stack slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stacks")
	}

	if len(stackAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Stack) ReloadG() error {
	if o == nil {
		return errors.New("models: no Stack provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stack) Reload(exec boil.Executor) error {
	ret, err := FindStack(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StackSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StackSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StackSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StackSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `stacks`.* FROM `stacks` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stackPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StackSlice")
	}

	*o = slice

	return nil
}

// StackExistsG checks if the Stack row exists.
func StackExistsG(iD uint64) (bool, error) {
	return StackExists(boil.GetDB(), iD)
}

// StackExists checks if the Stack row exists.
func StackExists(exec boil.Executor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `stacks` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stacks exists")
	}

	return exists, nil
}
