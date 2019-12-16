// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package article

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"time"
)

// arModel is a active record design model for table gf_article operations.
type arModel struct {
	M *gdb.Model
}

var (
	// Table is the table name of gf_article.
	Table = "gf_article"
	// Model is the model object of gf_article.
	Model = &arModel{g.DB("default").Table(Table).Safe()}
	// Primary is the primary key name of table gf_article.
	Primary = gdb.GetPrimaryKey(new(Entity))
)

// FindOne retrieves and returns a single Entity by a primary key or where conditions by
// Model.Where. The optional parameter <where> is like follows:
// 123, []int{1, 2, 3}, "john", []string{"john", "smith"}
// g.Map{"id": g.Slice{1,2,3}}, g.Map{"id": 1, "name": "john"}, etc.
//
// Note that it differs with Mode.One is that any single where condition parameter will
// be treated as the value of the primary key in FindOne, but a string condition in
// Model.One. That is, if primary key is "id" and given <where> parameter as "123", the
// FindOne function treats it as "id=123", but Model.One treats it as "123", just a string.
func FindOne(where ...interface{}) (*Entity, error) {
	return Model.One(gdb.GetPrimaryKeyCondition(Primary, where...)...)
}

// FindAll retrieves and returns multiple Entity by a primary key or where conditions by
// Model.Where. Also see FindOne.
func FindAll(where ...interface{}) ([]*Entity, error) {
	return Model.All(gdb.GetPrimaryKeyCondition(Primary, where...)...)
}

// FindValue retrieves and returns single field value by a primary key or where conditions by
// Model.Where. Also see FindOne.
func FindValue(fieldsAndWhere ...interface{}) (gdb.Value, error) {
	if len(fieldsAndWhere) == 2 {
		return Model.Value(append(fieldsAndWhere[0:1], gdb.GetPrimaryKeyCondition(Primary, fieldsAndWhere[1:]...)...)...)
	}
	return Model.Value(fieldsAndWhere...)
}

// Insert is alias of Model.Insert.
func Insert(data ...interface{}) (result sql.Result, err error) {
	return Model.Insert(data...)
}

// Replace is alias of Model.Replace.
func Replace(data ...interface{}) (result sql.Result, err error) {
	return Model.Replace(data...)
}

// Save is alias of Model.Save.
func Save(data ...interface{}) (result sql.Result, err error) {
	return Model.Save(data...)
}

// Update is alias of Model.Update.
func Update(dataAndWhere ...interface{}) (result sql.Result, err error) {
	return Model.Update(dataAndWhere...)
}

// Delete is alias of Model.Delete.
func Delete(where ...interface{}) (result sql.Result, err error) {
	return Model.Delete(where...)
}

// TX sets the transaction for current operation.
func (m *arModel) TX(tx *gdb.TX) *arModel {
	return &arModel{m.M.TX(tx)}
}

// Master marks the following operation on master node.
func (m *arModel) Master() *arModel {
	return &arModel{m.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (m *arModel) Slave() *arModel {
	return &arModel{m.M.Slave()}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
func (m *arModel) LeftJoin(joinTable string, on string) *arModel {
	return &arModel{m.M.LeftJoin(joinTable, on)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
func (m *arModel) RightJoin(joinTable string, on string) *arModel {
	return &arModel{m.M.RightJoin(joinTable, on)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
func (m *arModel) InnerJoin(joinTable string, on string) *arModel {
	return &arModel{m.M.InnerJoin(joinTable, on)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
func (m *arModel) Fields(fields string) *arModel {
	return &arModel{m.M.Fields(fields)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
func (m *arModel) FieldsEx(fields string) *arModel {
	return &arModel{m.M.FieldsEx(fields)}
}

// Option sets the extra operation option for the model.
func (m *arModel) Option(option int) *arModel {
	return &arModel{m.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (m *arModel) OmitEmpty() *arModel {
	return &arModel{m.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (m *arModel) Filter() *arModel {
	return &arModel{m.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (m *arModel) Where(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.Where(where, args...)}
}

// And adds "AND" condition to the where statement.
func (m *arModel) And(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (m *arModel) Or(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.Or(where, args...)}
}

// GroupBy sets the "GROUP BY" statement for the model.
func (m *arModel) GroupBy(groupBy string) *arModel {
	return &arModel{m.M.GroupBy(groupBy)}
}

// OrderBy sets the "ORDER BY" statement for the model.
func (m *arModel) OrderBy(orderBy string) *arModel {
	return &arModel{m.M.OrderBy(orderBy)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (m *arModel) Limit(limit ...int) *arModel {
	return &arModel{m.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (m *arModel) Offset(offset int) *arModel {
	return &arModel{m.M.Offset(offset)}
}

// ForPage sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (m *arModel) ForPage(page, limit int) *arModel {
	return &arModel{m.M.ForPage(page, limit)}
}

// Batch sets the batch operation number for the model.
func (m *arModel) Batch(batch int) *arModel {
	return &arModel{m.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (m *arModel) Cache(expire time.Duration, name ...string) *arModel {
	return &arModel{m.M.Cache(expire, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (m *arModel) Data(data ...interface{}) *arModel {
	return &arModel{m.M.Data(data...)}
}

// Insert does "INSERT INTO ..." statement for the model.
// The optional parameter <data> is the same as the parameter of Model.Data function,
// see Model.Data.
func (m *arModel) Insert(data ...interface{}) (result sql.Result, err error) {
	return m.M.Insert(data...)
}

// Replace does "REPLACE INTO ..." statement for the model.
// The optional parameter <data> is the same as the parameter of Model.Data function,
// see Model.Data.
func (m *arModel) Replace(data ...interface{}) (result sql.Result, err error) {
	return m.M.Replace(data...)
}

// Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the model.
// It updates the record if there's primary or unique index in the saving data,
// or else it inserts a new record into the table.
//
// The optional parameter <data> is the same as the parameter of Model.Data function,
// see Model.Data.
func (m *arModel) Save(data ...interface{}) (result sql.Result, err error) {
	return m.M.Save(data...)
}

// Update does "UPDATE ... " statement for the model.
//
// If the optional parameter <dataAndWhere> is given, the dataAndWhere[0] is the updated
// data field, and dataAndWhere[1:] is treated as where condition fields.
// Also see Model.Data and Model.Where functions.
func (m *arModel) Update(dataAndWhere ...interface{}) (result sql.Result, err error) {
	return m.M.Update(dataAndWhere...)
}

// Delete does "DELETE FROM ... " statement for the model.
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
func (m *arModel) Delete(where ...interface{}) (result sql.Result, err error) {
	return m.M.Delete(where...)
}

// Count does "SELECT COUNT(x) FROM ..." statement for the model.
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
func (m *arModel) Count(where ...interface{}) (int, error) {
	return m.M.Count(where...)
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*Entity.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
func (m *arModel) All(where ...interface{}) ([]*Entity, error) {
	all, err := m.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*Entity
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *Entity.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
func (m *arModel) One(where ...interface{}) (*Entity, error) {
	one, err := m.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *Entity
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// Value retrieves a specified record value from table and returns the result as interface type.
// It returns nil if there's no record found with the given conditions from table.
//
// If the optional parameter <fieldsAndWhere> is given, the fieldsAndWhere[0] is the selected fields
// and fieldsAndWhere[1:] is treated as where condition fields.
// Also see Model.Fields and Model.Where functions.
func (m *arModel) Value(fieldsAndWhere ...interface{}) (gdb.Value, error) {
	return m.M.Value(fieldsAndWhere...)
}

// Chunk iterates the table with given size and callback function.
func (m *arModel) Chunk(limit int, callback func(entities []*Entity, err error) bool) {
	m.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*Entity
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}
