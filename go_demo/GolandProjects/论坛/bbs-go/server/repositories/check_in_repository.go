package repositories

import (
	"bbs-go/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var CheckInRepository = newCheckInRepository()

func newCheckInRepository() *checkInRepository {
	return &checkInRepository{}
}

type checkInRepository struct {
}

func (r *checkInRepository) Get(db *gorm.DB, id int64) *model.CheckIn {
	ret := &model.CheckIn{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *checkInRepository) Take(db *gorm.DB, where ...interface{}) *model.CheckIn {
	ret := &model.CheckIn{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *checkInRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.CheckIn) {
	cnd.Find(db, &list)
	return
}

func (r *checkInRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.CheckIn {
	ret := &model.CheckIn{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *checkInRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.CheckIn, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *checkInRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.CheckIn, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.CheckIn{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *checkInRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.CheckIn{})
}

func (r *checkInRepository) Create(db *gorm.DB, t *model.CheckIn) (err error) {
	err = db.Create(t).Error
	return
}

func (r *checkInRepository) Update(db *gorm.DB, t *model.CheckIn) (err error) {
	err = db.Save(t).Error
	return
}

func (r *checkInRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.CheckIn{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *checkInRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.CheckIn{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *checkInRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.CheckIn{}, "id = ?", id)
}
