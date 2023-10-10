package services

import (
	"errors"
	"gin-go-bl/internal/errmsg"
	"gin-go-bl/internal/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type ArticleServiceImpl struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{db: db}
}
func (art *ArticleServiceImpl) GetAllInfo(pageSize int, pageNum int) (errmsg.Error, int) {
	limit := pageSize
	offset := pageSize * (pageNum - 1)
	total := int64(0)
	list := []models.Article{}
	err := art.db.Model(&models.Article{}).Count(&total).Error
	err = art.db.Preload("Category").Raw("SELECT * FROM article LIMIT ?,?", offset, limit).Scan(&list).Error
	if err != nil {
		log.Println(err)
		return errmsg.ErrServer, http.StatusInternalServerError
	}
	return errmsg.OK.WithData(map[string]any{
		"list":  list,
		"total": total,
	}), http.StatusOK
}

func (art ArticleServiceImpl) GetIDArticle(id int) (errmsg.Error, int) {
	var article models.Article
	err := art.db.Preload("Category").Where("aid = ?", id).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errmsg.ErrArticleNotExist, http.StatusOK
		}
		return errmsg.ErrServer, http.StatusInternalServerError
	}
	return errmsg.OK.WithData(article), http.StatusOK

}

//func (art ArticleServiceImpl) GetUserArticle(uid int, pageSize int, pageNum int) (errmsg.Error, int) {
//	var list []models.Article
//	var total int64
//	limit := pageSize
//	offset := pageSize * (pageNum - 1)
//	art.db.Model(&models.Article{}).Where("user_id = ?", uid).Count(&total)
//	err := art.db.Preload("Category").Limit(limit).Offset(offset).Where("user_id = ?", uid).Find(&article).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return errmsg.ErrArticleNotExist, http.StatusOK
//		}
//		return errmsg.ErrServer, http.StatusInternalServerError
//	}
//	return errmsg.OK.WithData(map[string]any{
//		"list":  list,
//		"total": total,
//	}), http.StatusOK
//
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
