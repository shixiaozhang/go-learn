package repositories

import (
	"bbs-go/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var OperateLogRepository = newOperateLogRepository()

func newOperateLogRepository() *operateLogRepository {
	return &operateLogRepository{}
}

type operateLogRepository struct {
}

func (r *operateLogRepository) Get(db *gorm.DB, id int64) *model.OperateLog {
	ret := &model.OperateLog{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *operateLogRepository) Take(db *gorm.DB, where ...interface{}) *model.OperateLog {
	ret := &model.OperateLog{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *operateLogRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.OperateLog) {
	cnd.Find(db, &list)
	return
}

func (r *operateLogRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.OperateLog {
	ret := &model.OperateLog{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *operateLogRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.OperateLog, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *operateLogRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.OperateLog, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.OperateLog{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *operateLogRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.OperateLog{})
}

func (r *operateLogRepository) Create(db *gorm.DB, t *model.OperateLog) (err error) {
	err = db.Create(t).Error
	return
}

func (r *operateLogRepository) Update(db *gorm.DB, t *model.OperateLog) (err error) {
	err = db.Save(t).Error
	return
}

func (r *operateLogRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.OperateLog{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *operateLogRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.OperateLog{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *operateLogRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.OperateLog{}, "id = ?", id)
}
