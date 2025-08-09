package store

import (
	"database/sql"
	"math/rand"
	"test/config"
)

type Img struct {
	Id     int64  `db:"id" json:"id" form:"id" example:"1"`                                       // 图片ID
	Base64 string `db:"base64" json:"base64" form:"base64" example:"https://example.com/img.png"` // 图片URL
}

func NewImg() *Img {
	return &Img{}
}

func (i *Img) TableName() string {
	return "imgs"
}

func (i *Img) GetImgBase64ById(id int64) (string, error) {
	var base64 string
	if err := config.Db.Unsafe().Get(&base64, "SELECT base64 FROM "+i.TableName()+" WHERE id = ?", id); err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return base64, nil
}

func (i *Img) GetImgById(id int64) (*Img, error) {
	img := &Img{}
	if err := config.Db.Unsafe().Get(img, "SELECT * FROM "+i.TableName()+" WHERE id = ?", id); err != nil && err != sql.ErrNoRows {
		return img, err
	}
	return img, nil
}

func (i *Img) GetRandomImgId() (int64, error) {
	// 随机获取 [10, 17] 之间的图片
	id := rand.Int63n(17-10+1) + 10
	img, err := i.GetImgById(id)
	if err != nil {
		return 0, err
	}
	return img.Id, nil
}

func (i *Img) CreateImg() error {
	query := `INSERT INTO ` + i.TableName() + ` (base64) VALUES (?)`
	res, err := config.Db.Exec(query, i.Base64)
	if err != nil {
		return err
	}
	i.Id, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
