package model

import "github.com/go-sql-driver/mysql"

// SubjectModel the definition of subject model
type SubjectModel struct {
	Model

	ParentID    uint64         `gorm:"column:parent_id;not null"`
	Name        string         `gorm:"column:name;not null"`
	Slug        string         `gorm:"column:slug;not null"`
	Description string         `gorm:"column:description;not null"`
	CoverImage  uint64         `gorm:"column:cover_image;not null"`
	IsEnd       uint64         `gorm:"column:is_end;not null"`
	Count       uint64         `gorm:"column:count;not null"`
	LastUpdated mysql.NullTime `gorm:"column:last_updated"`
}

// SubjectRelationshipsModel `pt_subject_relationships` 's struct
type SubjectRelationshipsModel struct {
	ObjectID  uint64 `gorm:"column:object_id;not null;primary_key"`
	SubjectID uint64 `gorm:"column:subject_id;not null;primary_key"`
	OrderNum  string `gorm:"column:order_num;not null"`
}

// TableName is the resource table name in db
func (c *SubjectModel) TableName() string {
	return "pt_subject"
}

// TableName SubjectRelationshipsModel's binding db name
func (c *SubjectRelationshipsModel) TableName() string {
	return "pt_subject_relationships"
}

// Create creates a new subject
func (c *SubjectModel) Create() error {
	return DB.Local.Create(&c).Error
}

// Update update subject
func (c *SubjectModel) Update() (err error) {
	if err = DB.Local.Model(&SubjectModel{}).Save(c).Error; err != nil {
		return err
	}

	return nil
}

// GetSubjectByID get subject info by ID
func GetSubjectByID(id uint64) (*SubjectModel, error) {
	s := &SubjectModel{}
	result := DB.Local.Where("id = ?", id).First(&s)
	return s, result.Error
}

// GetAllSubjects get all subjects
func GetAllSubjects() ([]*SubjectModel, error) {
	var subjects []*SubjectModel
	result := DB.Local.Find(&subjects)

	return subjects, result.Error
}

// SubjectCheckNameExist check the subject name if is already exist
func SubjectCheckNameExistWhileCreate(name string) bool {
	count := 0
	subjectModel := &SubjectModel{}
	DB.Local.Table(subjectModel.TableName()).
		Where("`name` = ?", name).
		Count(&count)

	if count > 0 {
		return true
	}

	return false
}

// SubjectCheckNameExistWhileUpdate check the subject name if is already exist without itself
func SubjectCheckNameExistWhileUpdate(subjectID uint64, name string) bool {
	count := 0
	subjectModel := &SubjectModel{}
	DB.Local.Table(subjectModel.TableName()).
		Where("`id` != ? AND `name` = ?", subjectID, name).
		Count(&count)

	if count > 0 {
		return true
	}

	return false
}
