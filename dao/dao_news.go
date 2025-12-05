package dao

import (
	"errors"
	"fmt"
	"go-film-demo/model/news"
	"go-film-demo/plugin/db"
	"time"

	"gorm.io/gorm"
)

// NewsRepository 新闻仓储层
type NewsRepository struct {
}

// NewNewsRepository 创建新闻仓储实例
func NewNewsRepository() *NewsRepository {
	return &NewsRepository{}
}

// Create 创建新闻
func (r *NewsRepository) Create(news *news.News) error {
	result := db.Mdb.Create(news)
	if result.Error != nil {
		return fmt.Errorf("failed to create news: %w", result.Error)
	}
	return nil
}

// CreateBatch 批量创建新闻
func (r *NewsRepository) CreateBatch(newsList []*news.News) error {
	return db.Mdb.Transaction(func(tx *gorm.DB) error {
		for _, news := range newsList {
			if err := tx.Create(news).Error; err != nil {
				return fmt.Errorf("failed to batch create news: %w", err)
			}
		}
		return nil
	})
}

// GetByID 根据ID获取新闻
func (r *NewsRepository) GetByID(id uint64) (*news.News, error) {
	var news news.News
	result := db.Mdb.First(&news, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("news not found with id: %d", id)
		}
		return nil, fmt.Errorf("failed to get news by id: %w", result.Error)
	}
	return &news, nil
}

// GetByNewsID 根据新闻ID获取新闻
func (r *NewsRepository) GetByNewsID(newsID string) (*news.News, error) {
	var news news.News
	result := db.Mdb.Where("news_id = ?", newsID).First(&news)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("news not found with news_id: %s", newsID)
		}
		return nil, fmt.Errorf("failed to get news by news_id: %w", result.Error)
	}
	return &news, nil
}

// Update 更新新闻
func (r *NewsRepository) Update(news *news.News) error {
	result := db.Mdb.Save(news)
	if result.Error != nil {
		return fmt.Errorf("failed to update news: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("news not found for update")
	}
	return nil
}

// UpdateFields 更新指定字段
func (r *NewsRepository) UpdateFields(id uint64, updates map[string]interface{}) error {
	result := db.Mdb.Model(&news.News{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("failed to update news fields: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("news not found for field update")
	}
	return nil
}

// Delete 删除新闻（软删除 - 如果需要硬删除直接使用 Unscoped().Delete）
func (r *NewsRepository) Delete(id uint64) error {
	result := db.Mdb.Delete(&news.News{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete news: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("news not found for deletion")
	}
	return nil
}

// HardDelete 硬删除新闻
func (r *NewsRepository) HardDelete(id uint64) error {
	result := db.Mdb.Unscoped().Delete(&news.News{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to hard delete news: %w", result.Error)
	}
	return nil
}

// List 获取新闻列表（带分页和条件）
func (r *NewsRepository) List(page, pageSize int, conditions map[string]interface{}) ([]*news.News, int64, error) {
	var newsList []*news.News
	var total int64

	query := db.Mdb.Model(&news.News{})

	// 应用查询条件
	for field, value := range conditions {
		query = query.Where(field+" = ?", value)
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count news: %w", err)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.
		Order("publish_time DESC, created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&newsList).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to list news: %w", err)
	}

	return newsList, total, nil
}

// ListByCreator 根据作者查询新闻
func (r *NewsRepository) ListByCreator(creator string, page, pageSize int) ([]*news.News, int64, error) {
	conditions := map[string]interface{}{
		"creator": creator,
	}
	return r.List(page, pageSize, conditions)
}

// ListBySource 根据来源查询新闻
func (r *NewsRepository) ListBySource(source string, page, pageSize int) ([]*news.News, int64, error) {
	conditions := map[string]interface{}{
		"source": source,
	}
	return r.List(page, pageSize, conditions)
}

// ListByPublishTimeRange 根据发布时间范围查询
func (r *NewsRepository) ListByPublishTimeRange(startTime, endTime time.Time) ([]*news.News, int64, error) {
	var newsList []*news.News
	var total int64

	query := db.Mdb.Model(&news.News{}).
		Where("publish_time BETWEEN ? AND ?", startTime, endTime)

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count news by time range: %w", err)
	}

	// 分页查询
	err = query.
		Order("publish_time DESC").
		Find(&newsList).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to list news by time range: %w", err)
	}

	return newsList, total, nil
}

// SearchByTitle 根据标题搜索新闻
func (r *NewsRepository) SearchByTitle(keyword string, page, pageSize int) ([]*news.News, int64, error) {
	var newsList []*news.News
	var total int64

	query := db.Mdb.Model(&news.News{}).
		Where("title LIKE ?", "%"+keyword+"%")

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count news by title search: %w", err)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.
		Order("publish_time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&newsList).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to search news by title: %w", err)
	}

	return newsList, total, nil
}

func (r *NewsRepository) GetNewsByDateRange(targetDate time.Time) ([]news.News, error) {
	var newsList []news.News

	// 计算当天的开始和结束时间
	startOfDay := time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(), 0, 0, 0, 0, targetDate.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := db.Mdb.Where("publish_time >= ? AND publish_time < ?", startOfDay, endOfDay).
		Order("publish_time DESC").
		Find(&newsList).Error

	return newsList, err
}
