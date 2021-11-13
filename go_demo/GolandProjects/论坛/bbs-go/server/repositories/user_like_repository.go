package repositories

import (
	"github.com/mlogclub/simple"
	"gorm.io/gorm"

	"bbs-go/model"
)

var UserLikeRepository = newUserLikeRepository()

func newUserLikeRepository() *userLikeRepository {
	return &userLikeRepository{}
}

type userLikeRepository struct {
}

func (r *userLikeRepository) Get(db *gorm.DB, id int64) *model.UserLike {
	ret := &model.UserLike{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userLikeRepository) Take(db *gorm.DB, where ...interface{}) *model.UserLike {
	ret := &model.UserLike{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userLikeRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserLike) {
	cnd.Find(db, &list)
	return
}

func (r *userLikeRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserLike {
	ret := &model.UserLike{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userLikeRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserLike, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userLikeRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserLike, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserLike{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userLikeRepository) Create(db *gorm.DB, t *model.UserLike) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userLikeRepository) Update(db *gorm.DB, t *model.UserLike) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userLikeRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserLike{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userLikeRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserLike{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userLikeRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserLike{}, "id = ?", id)
}
