package repository

import (
	"cramee/codegen/template/api/repository/util"
	"cramee/codegen/template/api/repository/util/querybuilder"
	"cramee/codegen/template/api/services"
	"cramee/codegen/template/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

func NewPlaceHolderRepository(db *gorm.DB) services.PlaceHolderRepository {
	res := &placeHolderRepositoryImpl{}
	res.db = db
	res.now = time.Now
	return res
}

type placeHolderRepositoryImpl struct {
	db  *gorm.DB
	now func() time.Time
}

func (m *placeHolderRepositoryImpl) GetByID(id uint, expand ...string) (*models.PlaceHolder, error) {
	data := &models.PlaceHolder{}
	db := m.db.Unscoped()
	db, err := querybuilder.BuildExpandQuery(&models.PlaceHolder{}, expand, db, func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	})
	if err != nil {
		return nil, err
	}
	if err := db.Unscoped().Where("id = ?", id).First(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

type GetAllPlaceHolderBaseQueryBuildFunc func(db *gorm.DB) (*gorm.DB, error)

func GetAllPlaceHolderBase(config services.GetAllConfig, db *gorm.DB, queryBuildFunc GetAllPlaceHolderBaseQueryBuildFunc) ([]*models.PlaceHolder, uint, error) {
	var limit int = util.GetAllMaxLimit
	var offset int = 0
	var allCount int64
	var (
		err   error
		model []*models.PlaceHolder = []*models.PlaceHolder{}
		q     *gorm.DB              = db.Model(&models.PlaceHolder{})
	)
	if config.Limit > 0 {
		limit = int(config.Limit)
	}
	if config.Offset > 0 {
		offset = int(config.Offset)
	}
	if config.IncludeDeleted {
		q = q.Unscoped()
	}
	if config.OnlyDeleted {
		q = q.Unscoped().Where("deleted_at is not null")
	}
	q, err = querybuilder.BuildQueryQuery(&models.PlaceHolder{}, config.Query, q)
	if err != nil {
		return nil, 0, err
	}
	q, err = querybuilder.BuildOrderQuery(&models.PlaceHolder{}, config.Order, q)
	if err != nil {
		return nil, 0, err
	}
	q, err = querybuilder.BuildExpandQuery(&models.PlaceHolder{}, config.Expand, q, func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	})
	if err != nil {
		return nil, 0, err
	}
	if queryBuildFunc != nil {
		q, err = queryBuildFunc(q)
		if err != nil {
			return nil, 0, err
		}
	}
	// ??????10000??????????????????????????????????????????
	load := func() (bool, error) {
		var sub []models.PlaceHolder
		subLimit := util.GetAllSubLimit
		if limit <= subLimit {
			subLimit = limit + 1
		}
		if err := q.Offset(offset).Limit(subLimit).Find(&sub).Error; err != nil {
			return false, err
		}
		var size int
		model, size = MergePlaceHolderSlice(model, sub)
		offset += size
		limit -= size
		return size < subLimit || limit < 0, nil
	}
	for {
		shouldEnd, err := load()
		if err != nil {
			return nil, 0, err
		}
		if shouldEnd {
			break
		}
	}

	if (config.Limit > 0 && uint(len(model)) > config.Limit) || config.Offset > 0 {
		if err := q.Model(&models.PlaceHolder{}).Count(&allCount).Error; err != nil {
			return nil, 0, err
		}
	} else {
		allCount = int64(len(model))
	}
	if config.Limit > 0 && uint(len(model)) > config.Limit {
		model = model[:config.Limit]
	}
	if len(model) > util.GetAllMaxLimit {
		return nil, 0, errors.New("??????????????????????????????????????????????????????")
	}
	return model, uint(allCount), nil
}

func (m *placeHolderRepositoryImpl) GetAll(config services.GetAllConfig) ([]*models.PlaceHolder, uint, error) {
	return GetAllPlaceHolderBase(config, m.db, nil)
}

func (m *placeHolderRepositoryImpl) Create(data *models.PlaceHolder) (*models.PlaceHolder, error) {
	data = util.ShallowCopy(data).(*models.PlaceHolder)
	now := m.now()
	data.SetUpdatedAt(now)
	data.SetCreatedAt(now)
	if err := m.db.Create(data).Error; err != nil {
		return nil, err
	}
	data, err := m.GetByID(data.GetID())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *placeHolderRepositoryImpl) Update(id uint, data *models.PlaceHolder) (*models.PlaceHolder, error) {
	orgData, err := m.GetByID(id)
	if err != nil {
		return nil, err
	}
	if data.GetID() != orgData.GetID() {
		return nil, errors.New("ID????????????????????????")
	}
	if data.GetCreatedAt().UTC().Unix() != orgData.GetCreatedAt().UTC().Unix() {
		return nil, errors.New("????????????????????????????????????")
	}
	if data.GetUpdatedAt().UTC().Unix() != orgData.GetUpdatedAt().UTC().Unix() {
		return nil, errors.New("????????????????????????????????????")
	}
	if data.GetDeletedAt() != orgData.GetDeletedAt() {
		if data.GetDeletedAt() == nil && orgData.GetDeletedAt() != nil {
		} else if data.GetDeletedAt() == nil || orgData.GetDeletedAt() == nil {
			return nil, errors.New("????????????????????????????????????")
		} else if data.GetDeletedAt().UTC().Unix() != orgData.GetDeletedAt().UTC().Unix() {
			return nil, errors.New("????????????????????????????????????")
		}
	}
	data.SetUpdatedAt(m.now())
	if err := m.db.Unscoped().Save(data).Error; err != nil {
		return nil, err
	}
	data, err = m.GetByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *placeHolderRepositoryImpl) SoftDelete(id uint) (*models.PlaceHolder, error) {
	data, err := m.GetByID(id)
	if err != nil {
		return nil, err
	}
	data.SetDeletedAt(m.now())
	if err := m.db.Unscoped().Save(data).Error; err != nil {
		return nil, err
	}
	data, err = m.GetByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *placeHolderRepositoryImpl) HardDelete(id uint) (*models.PlaceHolder, error) {
	data, err := m.GetByID(id)
	if err != nil {
		return nil, err
	}
	if !data.IsDeleted() {
		return nil, errors.New("????????????????????????????????????????????????????????????????????????????????????")
	}
	if err := m.db.Unscoped().Delete(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (m *placeHolderRepositoryImpl) Restore(id uint) (*models.PlaceHolder, error) {
	data, err := m.GetByID(id)
	if err != nil {
		return nil, err
	}
	if err := m.db.Unscoped().Save(data).Error; err != nil {
		return nil, err
	}
	data, err = m.GetByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func MergePlaceHolderSlice(s []*models.PlaceHolder, t []models.PlaceHolder) ([]*models.PlaceHolder, int) {
	for i := range t {
		s = append(s, &t[i])
	}
	return s, len(t)
}
