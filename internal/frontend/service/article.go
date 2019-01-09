package article

import (
	"strconv"

	"github.com/puti-projects/puti/internal/common/model"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"

	"github.com/lexkong/log"
)

func GetArticles(page int, keyword string) (articles []*model.PostModel, pagination *util.Pagination) {
	pageSize, _ := strconv.Atoi(optionCache.Options.Get("posts_per_page"))
	offset := (page - 1) * pageSize
	count := 0

	where := "`post_type` = ? AND `parent_id` = ? AND `status` = ?"
	whereArgs := []interface{}{"articles", 0, "publish"}
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

	return
}
