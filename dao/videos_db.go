/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:00:14
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-03 11:14:57
 */

package dao

import (
	"govideo_server/model"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DelVideos() {
	db.Delete(&model.Video{})
}

func AddVideo(video *model.Video) bool {
	db.Create(&video)
	return !db.NewRecord(video)
}

func GetVideoByID(id uint) *model.Video {
	video := &model.Video{}
	db.First(video, id)
	return video
}

func GetVideoByTitle(title string) *model.Video {
	video := &model.Video{}
	db.Where("title=?", title).First(video)
	return video
}

func GetVideos() (videos []*model.Video) {
	db.Find(&videos)
	return videos
}