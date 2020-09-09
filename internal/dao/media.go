package dao

import (
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/utils"
)

// CreateMedia crate new media record after uploaded
func (d *Dao) CreateMedia(uID int, filename, fileNameWithoutExt, fileExt, pathName, usage string) (uint64, string, error) {
	media := &model.Media{
		UserID:   uint64(uID),
		Title:    filename,
		Slug:     fileNameWithoutExt,
		GUID:     pathName,
		MimeType: utils.GetFileMimeTypeByExt(fileExt),
		Usage:    usage,
	}

	if err := media.Create(d.db); err != nil {
		return 0, "", errno.New(errno.ErrDatabase, err)
	}
	return media.ID, media.GUID, nil
}

// UpdateMedia update media info
func (d *Dao) UpdateMedia(ID uint64, title, slug, description string) error {
	// Get old media info
	oldMedia := &model.Media{
		Model: model.Model{ID: ID},
	}
	err := oldMedia.GetByID(d.db)
	if err != nil {
		return err
	}

	// Set new status values
	oldMedia.Title = title
	oldMedia.Slug = slug
	oldMedia.Description = description

	err = oldMedia.Update(d.db)
	return err
}

// ListMedia get media list
func (d *Dao) ListMedia(limit, page int) ([]*model.Media, int64, error) {
	where := "deleted_time is null"

	media := &model.Media{}
	count, err := media.Count(d.db, where, nil)
	if err != nil {
		return nil, count, err
	}

	offset := (page - 1) * limit
	medias, err := media.List(d.db, where, nil, offset, limit)
	if err != nil {
		return nil, count, err
	}

	return medias, count, nil
}

// GetMediaByID get media by ID
func (d *Dao) GetMediaByID(ID uint64) (*model.Media, error) {
	media := &model.Media{
		Model: model.Model{ID: ID},
	}
	err := media.GetByID(d.db)
	return media, err
}

// DeleteMediaByID delete media by ID
func (d *Dao) DeleteMediaByID(ID uint64) error {
	media := &model.Media{}
	media.ID = ID
	return media.Delete(d.db)
}
