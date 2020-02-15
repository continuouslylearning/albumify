package album

import (
	"io/ioutil"
	"net/http"

	. "albumify/db"
	"context"
	"sync"

	"github.com/gin-gonic/gin"
)

func deleteImage(c *gin.Context) {
	username := c.MustGet("username").(string)
	s3 := c.MustGet("s3").(*S3Handler)
	redis := c.MustGet("redis").(*RedisHandler)

	var req DeleteRequest
	e := c.ShouldBind(&req)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Must specify an object key",
		})
		return
	}

	key := req.Key
	e = s3.DeleteImage(key, username)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	e = redis.RemoveKey(username, key)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func getAlbum(c *gin.Context) {
	username := c.MustGet("username").(string)
	s3 := c.MustGet("s3").(*S3Handler)
	redis := c.MustGet("redis").(*RedisHandler)

	keys, ok, e := redis.GetCachedAlbum(username)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if !ok {
		keys, e = s3.GetAlbumKeys(username)
		if e != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		e = redis.CacheAlbum(username, keys)
		if e != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	imageURLs, e := redis.GetCachedURLs(keys)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	imageURLs, e = s3.GetURLs(username, keys, imageURLs)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	e = redis.CacheURLs(keys, imageURLs)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, imageURLs)
}

func postImages(c *gin.Context) {
	username := c.MustGet("username").(string)
	s3 := c.MustGet("s3").(*S3Handler)
	redis := c.MustGet("redis").(*RedisHandler)
	form, _ := c.MultipartForm()

	var images [][]byte
	var keys []string
	for key, file := range form.File {
		fileContent, _ := file[0].Open()
		image, _ := ioutil.ReadAll(fileContent)
		images = append(images, image)
		keys = append(keys, key)
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(2)

	go func() {
		wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
		}

		e := s3.UploadImages(keys, images, username)
		if e != nil {
			cancel()
			return
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
		}
		e := redis.AddKeys(username, keys)
		if e != nil {
			cancel()
			return
		}
	}()

	wg.Wait()
	if ctx.Err() != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// wg.Add(len(images) + 1)

	// for i, image := range images {
	// 	go func(key string, image []byte) {
	// 		defer wg.Done()

	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		default:
	// 		}

	// 		e := s3.UploadImage(key, image, username)
	// 		if e != nil {
	// 			cancel()
	// 			return
	// 		}
	// 	}(keys[i], image)
	// }

	// go func() {
	// 	defer wg.Done()

	// 	select {
	// 	case <-ctx.Done():
	// 		return
	// 	default:
	// 	}

	// 	e := redis.AddKeys(username, keys)
	// 	if e != nil {
	// 		cancel()
	// 		return
	// 	}
	// }()

	// wg.Wait()

	// if ctx.Err() != nil {
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	return
	// }
	c.AbortWithStatus(http.StatusCreated)
}
