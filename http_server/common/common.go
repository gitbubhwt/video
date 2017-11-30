package common

import (
	"video/common"
	"video/db"
)

//批量保存文件路径
func BatchSaveVideoPath(videoPaths []*common.VideoPath) error {
	databaseDB := db.GetMysql().DB()
	tx, err := databaseDB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(common.VIDEO_PATH_INNERT_SQL)
	if err != nil {
		return err
	}
	for _, v := range videoPaths {
		if _, err = stmt.Exec(v.VideoId, v.Path, v.OrderNum); err != nil {
			break
		}
	}
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}
