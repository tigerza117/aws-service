// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"api/model"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newTx(db *gorm.DB, opts ...gen.DOOption) tx {
	_tx := tx{}

	_tx.txDo.UseDB(db, opts...)
	_tx.txDo.UseModel(&model.Tx{})

	tableName := _tx.txDo.TableName()
	_tx.ALL = field.NewAsterisk(tableName)
	_tx.ID = field.NewUint(tableName, "id")
	_tx.CreatedAt = field.NewTime(tableName, "created_at")
	_tx.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tx.DeletedAt = field.NewField(tableName, "deleted_at")
	_tx.AccountID = field.NewUint(tableName, "account_id")
	_tx.DstAccountID = field.NewUint(tableName, "dst_account_id")
	_tx.Amount = field.NewFloat64(tableName, "amount")
	_tx.Status = field.NewInt(tableName, "status")
	_tx.Account = txBelongsToAccount{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Account", "model.Account"),
		Customer: struct {
			field.RelationField
			Accounts struct {
				field.RelationField
				Customer struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Account.Customer", "model.Customer"),
			Accounts: struct {
				field.RelationField
				Customer struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Account.Customer.Accounts", "model.AccountList"),
				Customer: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Account.Customer.Accounts.Customer", "model.Customer"),
				},
			},
		},
	}

	_tx.DstAccount = txBelongsToDstAccount{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("DstAccount", "model.Account"),
	}

	_tx.fillFieldMap()

	return _tx
}

type tx struct {
	txDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	AccountID    field.Uint
	DstAccountID field.Uint
	Amount       field.Float64
	Status       field.Int
	Account      txBelongsToAccount

	DstAccount txBelongsToDstAccount

	fieldMap map[string]field.Expr
}

func (t tx) Table(newTableName string) *tx {
	t.txDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tx) As(alias string) *tx {
	t.txDo.DO = *(t.txDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tx) updateTableName(table string) *tx {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.AccountID = field.NewUint(table, "account_id")
	t.DstAccountID = field.NewUint(table, "dst_account_id")
	t.Amount = field.NewFloat64(table, "amount")
	t.Status = field.NewInt(table, "status")

	t.fillFieldMap()

	return t
}

func (t *tx) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tx) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["account_id"] = t.AccountID
	t.fieldMap["dst_account_id"] = t.DstAccountID
	t.fieldMap["amount"] = t.Amount
	t.fieldMap["status"] = t.Status

}

func (t tx) clone(db *gorm.DB) tx {
	t.txDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tx) replaceDB(db *gorm.DB) tx {
	t.txDo.ReplaceDB(db)
	return t
}

type txBelongsToAccount struct {
	db *gorm.DB

	field.RelationField

	Customer struct {
		field.RelationField
		Accounts struct {
			field.RelationField
			Customer struct {
				field.RelationField
			}
		}
	}
}

func (a txBelongsToAccount) Where(conds ...field.Expr) *txBelongsToAccount {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a txBelongsToAccount) WithContext(ctx context.Context) *txBelongsToAccount {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a txBelongsToAccount) Model(m *model.Tx) *txBelongsToAccountTx {
	return &txBelongsToAccountTx{a.db.Model(m).Association(a.Name())}
}

type txBelongsToAccountTx struct{ tx *gorm.Association }

func (a txBelongsToAccountTx) Find() (result *model.Account, err error) {
	return result, a.tx.Find(&result)
}

func (a txBelongsToAccountTx) Append(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a txBelongsToAccountTx) Replace(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a txBelongsToAccountTx) Delete(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a txBelongsToAccountTx) Clear() error {
	return a.tx.Clear()
}

func (a txBelongsToAccountTx) Count() int64 {
	return a.tx.Count()
}

type txBelongsToDstAccount struct {
	db *gorm.DB

	field.RelationField
}

func (a txBelongsToDstAccount) Where(conds ...field.Expr) *txBelongsToDstAccount {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a txBelongsToDstAccount) WithContext(ctx context.Context) *txBelongsToDstAccount {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a txBelongsToDstAccount) Model(m *model.Tx) *txBelongsToDstAccountTx {
	return &txBelongsToDstAccountTx{a.db.Model(m).Association(a.Name())}
}

type txBelongsToDstAccountTx struct{ tx *gorm.Association }

func (a txBelongsToDstAccountTx) Find() (result *model.Account, err error) {
	return result, a.tx.Find(&result)
}

func (a txBelongsToDstAccountTx) Append(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a txBelongsToDstAccountTx) Replace(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a txBelongsToDstAccountTx) Delete(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a txBelongsToDstAccountTx) Clear() error {
	return a.tx.Clear()
}

func (a txBelongsToDstAccountTx) Count() int64 {
	return a.tx.Count()
}

type txDo struct{ gen.DO }

type ITxDo interface {
	gen.SubQuery
	Debug() ITxDo
	WithContext(ctx context.Context) ITxDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITxDo
	WriteDB() ITxDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITxDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITxDo
	Not(conds ...gen.Condition) ITxDo
	Or(conds ...gen.Condition) ITxDo
	Select(conds ...field.Expr) ITxDo
	Where(conds ...gen.Condition) ITxDo
	Order(conds ...field.Expr) ITxDo
	Distinct(cols ...field.Expr) ITxDo
	Omit(cols ...field.Expr) ITxDo
	Join(table schema.Tabler, on ...field.Expr) ITxDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITxDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITxDo
	Group(cols ...field.Expr) ITxDo
	Having(conds ...gen.Condition) ITxDo
	Limit(limit int) ITxDo
	Offset(offset int) ITxDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITxDo
	Unscoped() ITxDo
	Create(values ...*model.Tx) error
	CreateInBatches(values []*model.Tx, batchSize int) error
	Save(values ...*model.Tx) error
	First() (*model.Tx, error)
	Take() (*model.Tx, error)
	Last() (*model.Tx, error)
	Find() ([]*model.Tx, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Tx, err error)
	FindInBatches(result *[]*model.Tx, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Tx) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITxDo
	Assign(attrs ...field.AssignExpr) ITxDo
	Joins(fields ...field.RelationField) ITxDo
	Preload(fields ...field.RelationField) ITxDo
	FirstOrInit() (*model.Tx, error)
	FirstOrCreate() (*model.Tx, error)
	FindByPage(offset int, limit int) (result []*model.Tx, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITxDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t txDo) Debug() ITxDo {
	return t.withDO(t.DO.Debug())
}

func (t txDo) WithContext(ctx context.Context) ITxDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t txDo) ReadDB() ITxDo {
	return t.Clauses(dbresolver.Read)
}

func (t txDo) WriteDB() ITxDo {
	return t.Clauses(dbresolver.Write)
}

func (t txDo) Session(config *gorm.Session) ITxDo {
	return t.withDO(t.DO.Session(config))
}

func (t txDo) Clauses(conds ...clause.Expression) ITxDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t txDo) Returning(value interface{}, columns ...string) ITxDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t txDo) Not(conds ...gen.Condition) ITxDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t txDo) Or(conds ...gen.Condition) ITxDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t txDo) Select(conds ...field.Expr) ITxDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t txDo) Where(conds ...gen.Condition) ITxDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t txDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITxDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t txDo) Order(conds ...field.Expr) ITxDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t txDo) Distinct(cols ...field.Expr) ITxDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t txDo) Omit(cols ...field.Expr) ITxDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t txDo) Join(table schema.Tabler, on ...field.Expr) ITxDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t txDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITxDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t txDo) RightJoin(table schema.Tabler, on ...field.Expr) ITxDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t txDo) Group(cols ...field.Expr) ITxDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t txDo) Having(conds ...gen.Condition) ITxDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t txDo) Limit(limit int) ITxDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t txDo) Offset(offset int) ITxDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t txDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITxDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t txDo) Unscoped() ITxDo {
	return t.withDO(t.DO.Unscoped())
}

func (t txDo) Create(values ...*model.Tx) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t txDo) CreateInBatches(values []*model.Tx, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t txDo) Save(values ...*model.Tx) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t txDo) First() (*model.Tx, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tx), nil
	}
}

func (t txDo) Take() (*model.Tx, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tx), nil
	}
}

func (t txDo) Last() (*model.Tx, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tx), nil
	}
}

func (t txDo) Find() ([]*model.Tx, error) {
	result, err := t.DO.Find()
	return result.([]*model.Tx), err
}

func (t txDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Tx, err error) {
	buf := make([]*model.Tx, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t txDo) FindInBatches(result *[]*model.Tx, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t txDo) Attrs(attrs ...field.AssignExpr) ITxDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t txDo) Assign(attrs ...field.AssignExpr) ITxDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t txDo) Joins(fields ...field.RelationField) ITxDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t txDo) Preload(fields ...field.RelationField) ITxDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t txDo) FirstOrInit() (*model.Tx, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tx), nil
	}
}

func (t txDo) FirstOrCreate() (*model.Tx, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tx), nil
	}
}

func (t txDo) FindByPage(offset int, limit int) (result []*model.Tx, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t txDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t txDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t txDo) Delete(models ...*model.Tx) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *txDo) withDO(do gen.Dao) *txDo {
	t.DO = *do.(*gen.DO)
	return t
}
