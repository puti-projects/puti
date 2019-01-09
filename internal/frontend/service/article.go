package service

import (
	"strconv"

	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"

	"github.com/lexkong/log"
)

// GetArticles get articles list
func GetArticles(currentPage int, keyword string) (articles []*model.PostModel, pagination *utils.Pagination) {
	pageSize, _ := strconv.Atoi(optionCache.Options.Get("posts_per_page"))
	offset := (currentPage - 1) * pageSize
	count := 0

	where := "`post_type` = ? AND `parent_id` = ? AND `status` = ?"
	whereArgs := []interface{}{model.PostTypeArticle, 0, model.PostStatusPublish}
	if "" != keyword {
		where += " AND `title` LIKE ?"
		whereArgs = append(whereArgs, "%"+keyword+"%")
	}

	result := model.DB.Local.Model(&model.PostModel{}).
		Select("`id`, `title`, `content_html`, `guid`, `cover_picture`, `comment_count`, `view_count`, `posted_time`").
		Where(where, whereArgs...).Count(&count).
		Order("`if_top` DESC, `posted_time` DESC").
		Offset(offset).Limit(pageSize).
		Find(&articles)

	if err := result.Error; err != nil {
		log.Error("get articles failed:", err)
	}

	pagination = utils.GetPagination(count, currentPage, pageSize, 0)

	return
}
