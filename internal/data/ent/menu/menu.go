// Code generated by ent, DO NOT EDIT.

package menu

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMenuName holds the string denoting the menu_name field in the database.
	FieldMenuName = "menu_name"
	// FieldLeaf holds the string denoting the leaf field in the database.
	FieldLeaf = "leaf"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldComponent holds the string denoting the component field in the database.
	FieldComponent = "component"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldMenuDesc holds the string denoting the menu_desc field in the database.
	FieldMenuDesc = "menu_desc"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the menu in the database.
	Table = "menus"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "menus"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "menus"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldMenuName,
	FieldLeaf,
	FieldPath,
	FieldComponent,
	FieldParentID,
	FieldMenuDesc,
	FieldSort,
	FieldLevel,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// MenuNameValidator is a validator for the "menu_name" field. It is called by the builders before save.
	MenuNameValidator func(string) error
	// DefaultLeaf holds the default value on creation for the "leaf" field.
	DefaultLeaf bool
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// ComponentValidator is a validator for the "component" field. It is called by the builders before save.
	ComponentValidator func(string) error
)

// OrderOption defines the ordering options for the Menu queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMenuName orders the results by the menu_name field.
func ByMenuName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuName, opts...).ToFunc()
}

// ByLeaf orders the results by the leaf field.
func ByLeaf(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLeaf, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByComponent orders the results by the component field.
func ByComponent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComponent, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByMenuDesc orders the results by the menu_desc field.
func ByMenuDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuDesc, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
