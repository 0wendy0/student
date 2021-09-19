package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AdminMgr struct {
	*_BaseMgr
}

// AdminMgr open func
func AdminMgr(db *gorm.DB) *_AdminMgr {
	if db == nil {
		panic(fmt.Errorf("AdminMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdminMgr{_BaseMgr: &_BaseMgr{DB: db.Table("admin"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdminMgr) GetTableName() string {
	return "admin"
}

// Get 获取
func (obj *_AdminMgr) Get() (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AdminMgr) Gets() (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AdminMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Admin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AdminMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取
func (obj *_AdminMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取
func (obj *_AdminMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithSalt salt获取
func (obj *_AdminMgr) WithSalt(salt string) Option {
	return optionFunc(func(o *options) { o.query["salt"] = salt })
}

// WithCreatetime createtime获取
func (obj *_AdminMgr) WithCreatetime(createtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}

// WithUpdatetime updatetime获取
func (obj *_AdminMgr) WithUpdatetime(updatetime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updatetime"] = updatetime })
}

// GetByOption 功能选项模式获取
func (obj *_AdminMgr) GetByOption(opts ...Option) (result Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AdminMgr) GetByOptions(opts ...Option) (results []*Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_AdminMgr) GetFromID(id uint) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AdminMgr) GetBatchFromID(ids []uint) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容
func (obj *_AdminMgr) GetFromUsername(username string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找
func (obj *_AdminMgr) GetBatchFromUsername(usernames []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_AdminMgr) GetFromPassword(password string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_AdminMgr) GetBatchFromPassword(passwords []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromSalt 通过salt获取内容
func (obj *_AdminMgr) GetFromSalt(salt string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`salt` = ?", salt).Find(&results).Error

	return
}

// GetBatchFromSalt 批量查找
func (obj *_AdminMgr) GetBatchFromSalt(salts []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`salt` IN (?)", salts).Find(&results).Error

	return
}

// GetFromCreatetime 通过createtime获取内容
func (obj *_AdminMgr) GetFromCreatetime(createtime time.Time) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`createtime` = ?", createtime).Find(&results).Error

	return
}

// GetBatchFromCreatetime 批量查找
func (obj *_AdminMgr) GetBatchFromCreatetime(createtimes []time.Time) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`createtime` IN (?)", createtimes).Find(&results).Error

	return
}

// GetFromUpdatetime 通过updatetime获取内容
func (obj *_AdminMgr) GetFromUpdatetime(updatetime time.Time) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`updatetime` = ?", updatetime).Find(&results).Error

	return
}

// GetBatchFromUpdatetime 批量查找
func (obj *_AdminMgr) GetBatchFromUpdatetime(updatetimes []time.Time) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`updatetime` IN (?)", updatetimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AdminMgr) FetchByPrimaryKey(id uint) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`id` = ?", id).Find(&result).Error

	return
}
