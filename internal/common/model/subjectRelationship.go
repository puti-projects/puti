package model

// SubjectRelationshipsModel `pt_subject_relationships` 's struct
type SubjectRelationshipsModel struct {
	ObjectID  uint64 `gorm:"column:object_id;not null;primary_key"`
	SubjectID uint64 `gorm:"column:subject_id;not null;primary_key"`
	OrderNum  string `gorm:"column:order_num;not null"`
}

// TableName SubjectRelationshipsModel's binding db name
func (c *SubjectRelationshipsModel) TableName() string {
	return "pt_subject_relationships"
}

// GetArticleSubject get article's connection subject
func GetArticleSubject(articleID uint64) ([]*SubjectRelationshipsModel, error) {
	var subjectRelationships []*SubjectRelationshipsModel
	result := DB.Local.Where("`object_id` = ?", articleID).Find(&subjectRelationships)

	return subjectRelationships, result.Error
}
