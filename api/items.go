package api

import (
	"database/sql"
	"ecom/db/sqlc"
	"ecom/db/util"
	"ecom/token"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

type createItemRequest struct {
	Name     string `form:"name" binding:"required,alphanum"`
	Price    string `form:"price" binding:"required,numeric"`
	Category string `form:"category" binding:"required,alpha"`
}
type createItemResponse struct {
	Urls     []string `json:"urls"`
	Name     string   `json:"name"`
	Price    string   `json:"price"`
	Category string   `json:"category"`
}

type GetItemResponse struct {
	Urls      []string `json:"urls"`
	Name      string   `json:"name"`
	Price     string   `json:"price"`
	Category  string   `json:"category"`
	CreatedBy string   `json:"created_by"`
}

func (s *Server) createItem(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var req createItemRequest
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Print("err is ", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	param := sqlc.CreateItemParams{
		Name:      req.Name,
		Price:     req.Price,
		CreatedBy: authPayload.Username,
		Category:  req.Category,
	}
	item, err := s.store.GetItemByName(ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			item, err = s.store.CreateItem(ctx, param)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	}

	s3URl := []string{}
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	files := form.File["file"]
	for _, file := range files {
		if !util.ValidateMIMEType(file.Header.Get("Content-Type")) {
			ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid file type")))
			return
		}
		openfile, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		result, err := s.s3uploader.Upload(ctx,
			&s3.PutObjectInput{
				Bucket: &s.config.S3BUCKETNAME,
				Key:    aws.String(req.Name + "/" + file.Filename),
				Body:   openfile,
				Metadata: map[string]string{
					"created_by": authPayload.Username,
					"item_name":  item.Name,
				},
			},
		)
		if err != nil {
			log.Println("file uplaod fail for :", file.Filename)

		} else {
			_, err := s.store.CreateItemImage(ctx, sqlc.CreateItemImageParams{
				ItemID:   item.ID,
				ImageUrl: result.Location,
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}
			s3URl = append(s3URl, result.Location)
		}
	}
	res := createItemResponse{
		Urls:     s3URl,
		Name:     item.Name,
		Price:    item.Price,
		Category: item.Category,
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) getItem(ctx *gin.Context) {
	id := ctx.Param("id")
	itemId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid item id")))
		return
	}
	item, err := s.store.GetItemById(ctx, int32(itemId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("no item with itemid: %v", itemId)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	itemImages, _ := s.store.GetItemImagesFromItemId(ctx, int32(itemId))
	var urls []string
	if len(itemImages) > 0 {
		for _, v := range itemImages {
			urls = append(urls, v.ImageUrl)
		}
	}
	res := GetItemResponse{
		Urls:      urls,
		Name:      item.Name,
		Price:     item.Price,
		Category:  item.Category,
		CreatedBy: item.CreatedBy,
	}
	ctx.JSON(http.StatusOK, res)
}
