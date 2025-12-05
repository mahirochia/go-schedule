package news

import (
	"time"

	"gorm.io/gorm"
)

// News 新闻数据模型
type News struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;comment:新闻ID" json:"id"`
	NewsID      string    `gorm:"column:news_id;type:varchar(100);not null;uniqueIndex;comment:新闻唯一标识" json:"news_id"`
	Title       string    `gorm:"column:title;type:varchar(500);not null;comment:新闻标题" json:"title"`
	Creator     string    `gorm:"column:creator;type:varchar(100);comment:作者" json:"creator,omitempty"`
	Source      string    `gorm:"column:source;type:varchar(200);comment:新闻来源" json:"source,omitempty"`
	Content     string    `gorm:"column:content;type:longtext;comment:新闻内容" json:"content,omitempty"`
	PublishTime time.Time `gorm:"column:publish_time;type:datetime;comment:新闻发布时间" json:"publish_time,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null;onUpdate:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
	Cover       string    `gorm:"column:cover;type:varchar(550);default:'';not null;comment:封面图片url" json:"cover"`
	Desc        string    `gorm:"column:desc;type:varchar(1500);default:'';not null;comment:描述" json:"desc"`
	Link        string    `gorm:"column:link;type:varchar(1500);default:'';not null;comment:原文链接" json:"link"`
	CommentNum  int       `gorm:"column:comment_num;type:int;default:0;not null;comment:评论数" json:"comment_num"`
	ReadNum     int       `gorm:"column:read_num;type:int;default:0;not null;comment:浏览量" json:"read_num"`
	LikeNum     int       `gorm:"column:like_num;type:int;default:0;not null;comment:点赞数" json:"like_num"`
	CollectNum  int       `gorm:"column:collect_num;type:int;default:0;not null;comment:收藏数" json:"collect_num"`
	ShareNum    int       `gorm:"column:share_num;type:int;default:0;not null;comment:分享数" json:"share_num"`
}

// TableName 指定表名
func (News) TableName() string {
	return "news"
}

// BeforeCreate GORM 创建前的钩子
func (n *News) BeforeCreate(tx *gorm.DB) error {
	if n.CreatedAt.IsZero() {
		n.CreatedAt = time.Now()
	}
	if n.UpdatedAt.IsZero() {
		n.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate GORM 更新前的钩子
func (n *News) BeforeUpdate(tx *gorm.DB) error {
	n.UpdatedAt = time.Now()
	return nil
}
