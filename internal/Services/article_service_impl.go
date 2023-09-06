package Services

import (
	"gin-go-bl/internal/Models"
	"gin-go-bl/internal/errmsg"
	"log"
	"net/http"
)

type ArticleServiceImpl struct{}

func (art ArticleServiceImpl) GetAllInfo(pageSize int, pageNum int) (errmsg.Error, int) {
	db := DB
	limit := pageSize
	offset := pageSize * (pageNum - 1)
	total := int64(0)
	list := []Models.Article{}
	err = db.Model(&Models.Article{}).Count(&total).Error
	err := DB.Preload("Category").Raw("SELECT * FROM article LIMIT ?,?", offset, limit).Scan(&list).Error
	if err != nil {
		log.Println(err)
		return errmsg.ErrServer, http.StatusInternalServerError
	}
	return errmsg.OK.WithData(map[string]any{
		"list":  list,
		"total": total,
	}), http.StatusOK
}

//func (art ArticleServiceImpl) GetArticle(id int) (Article, int) {
//	var article Article
//	err := DB.Preload("Category").Where("aid = ?", id).First(&article).Error
//
//	switch {
//	case err != nil && err != gorm.ErrRecordNotFound:
//		return article, utils2.ERROR_ART_NULL
//	default:
//		return article, utils2.SUCCESS
//	}
//}

//func(art ArticleServiceImpl) GetUserArticle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
//	var article []Article
//	var total int64
//	DB.Model(&Article{}).Where("user_id = ?", id).Count(&total)
//	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("user_id = ?", id).Find(&article).Error
//	switch {
//	case err != nil && err != gorm.ErrRecordNotFound:
//		return nil, utils.ERROR_CAART_NULL, 0
//	default:
//		return article, utils.SUCCESS, total
//	}
//}
//func (art ArticleServiceImpl)GetCatArticle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
//	var article []Article
//	var total int64
//	DB.Model(&Article{}).Where("category_id = ?", id).Count(&total)
//	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("category_id = ?", id).Find(&article).Error
//	switch {
//	case err != nil && err != gorm.ErrRecordNotFound:
//		return nil, utils.ERROR_CAART_NULL, 0
//	default:
//		return article, utils.SUCCESS, total
//	}
//
//}
//func (art ArticleServiceImpl)CreatArticle(data *Article) int {
//	err := DB.Create(&data).Error
//
//	switch {
//	case err != nil:
//		fmt.Println(err)
//		return utils.ERROR
//	default:
//		return utils.SUCCESS
//	}
//
//}
//func (art ArticleServiceImpl)DeleteArticle(id int) int {
//	var article Article
//	err := DB.Where("aid = ?", id).Delete(&article).Error
//
//	switch {
//	case err != nil:
//		return utils.ERROR
//	default:
//		return utils.SUCCESS
//	}
//
//}
//func (art ArticleServiceImpl)EditArticle(id int, data *Article) int {
//	var article Article
//	var art = make(map[string]interface{})
//	art["title"] = data.Title
//	art["content"] = data.Content
//	art["img"] = data.Img
//	art["desc"] = data.Desc
//	art["category_id"] = data.CategoryId
//	//fmt.Printf("%v\n", &article)
//	err := DB.Model(&article).Where("aid = ?", id).Updates(art).Error
//	switch {
//	case err != nil:
//		return utils.ERROR
//	default:
//		return utils.SUCCESS
//	}
//
//}
