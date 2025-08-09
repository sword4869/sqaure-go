package store

import (
	"database/sql"
	"math/rand"
	"test/config"
)

type UserRsp struct {
	Id       int64  `json:"id" example:"1"`
	UserName string `json:"user_name" example:"春日收藏家"`
	Avatar   string `json:"avatar" example:"https://example.com/avatar.png"`
}

type User struct {
	Id       int64  `db:"id" json:"id" form:"id" example:"1"`                          // 用户ID
	UserName string `db:"user_name" json:"user_name" form:"user_name" example:"春日收藏家"` // 用户名
	AvatarId int64  `db:"avatar_id" json:"avatar_id" form:"avatar_id" example:"1"`     // 头像ID
}

func NewUser() *User {
	return &User{}
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetUserAvatarById(id int64) (*UserRsp, error) {
	user := &User{}
	if err := config.Db.Unsafe().Get(user, "SELECT * FROM "+u.TableName()+" WHERE id = ?", id); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	img := NewImg()
	imgBase64, err := img.GetImgBase64ById(user.AvatarId)
	if err != nil {
		return nil, err
	}
	return &UserRsp{
		Id:       user.Id,
		UserName: user.UserName,
		Avatar:   imgBase64,
	}, nil
}

func (u *User) CreateUser() error {
	for i := 0; i < 10; i++ {
		if u.UserName == "" {
			u.UserName = u.getRandomUserName()
		}
		user, err := u.GetByUserName(u.UserName)
		if err != nil {
			return err
		}
		if user != nil && user.Id > 0 {
			continue
		}
		break
	}
	if u.AvatarId == 0 {
		avatarId, err := u.getRandomAvatar()
		if err != nil {
			return err
		}
		u.AvatarId = avatarId
	}
	query := `INSERT INTO ` + u.TableName() + ` (user_name, avatar_id) VALUES (?, ?)`
	res, err := config.Db.Exec(query, u.UserName, u.AvatarId)
	if err != nil {
		return err
	}
	u.Id, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) getRandomUserName() string {
	names := []string{
		"春日收藏家",
		"相思故",
		"蓝眼睛不忧郁",
		"寄信给风",
		"世界等同你",
		"不眠日记",
		"岁月不休",
		"孤芳又自赏",
		"思念彼岸",
		"荆棘原野",
		"快乐很简单",
		"一页",
		"海边做诗意",
		"无心",
		"陌上花开可缓归",
		"一碗人间",
		"故事与月有关",
		"月窗染了杉",
		"浮生若舢",
		"拥抱人间烟火",
		"山河岁月空惆怅",
		"胭脂尽时桃花开",
		"寄意",
		"听萧与声",
		"冰河铺子",
		"自在安然",
		"千诗可叙",
		"芳春柳摇染花香",
		"风起风花落",
		"人间不值得",
		"青舟弄酒",
		"江心雾",
		"集市漫过街巷",
		"七里安黥",
		"北风陪我梦南柯",
		"云风未归",
		"青澜饮舟",
		"春日山杏",
		"树在夜里",
		"追赶日月",
		"游离子",
		"晚春里",
		"风起半山",
		"忘川在川",
		"月光衣我以华裳",
		"路过四月桃林",
		"夜莺与鲸",
		"我与春风皆过客",
		"我不会写诗",
		"心头的小情儿",
		"欲往",
		"一支云烟",
		"沿街灯火",
		"烟雾扰山河",
		"林中雨亭",
		"白首有我共你",
		"北境子栀",
		"月亮的根据地",
		"在爱的路上等你",
		"凉风有信",
		"花不解语",
		"守护在此方",
		"历尽山河走向你",
		"共枕拥眠",
		"念起便是柔情",
		"月上云朵",
		"雕刻成花",
		"林间鹿",
		"七月别困",
		"槐序廿柒",
		"月上溪溪",
		"梦明",
		"小草泠泠",
		"鱼沉秋水",
		"星月满屋",
		"浮世",
		"醉里烟波梦里寒",
		"朝暮最相思",
		"迷路的信",
		"小晴日记",
		"枕头说它不想醒",
		"月下客",
		"落日在山时",
		"世界不够温柔",
		"彼岸无岸",
		"山中雾",
		"清悠野鹤",
		"浮世三月",
		"一花一树开",
		"客情寄风絮",
		"屋顶上的小猫咪",
		"于暮夏风中",
		"小故事里的海",
		"风间白鹿",
		"山海和月",
		"寒岛春信",
		"月亮枕枝头",
		"夜止月明",
		"花伞没有雨",
		"花遇和风",
	}
	randomIndex := rand.Intn(len(names))
	userName := names[randomIndex]
	return userName
}

func (u *User) getRandomAvatar() (int64, error) {
	img := NewImg()
	imgId, err := img.GetRandomImgId()
	if err != nil {
		return 0, err
	}
	return imgId, nil
}

func (u *User) GetByUserName(userName string) (*User, error) {
	user := &User{}
	if err := config.Db.Unsafe().Get(user, "SELECT * FROM "+u.TableName()+" WHERE user_name = ?", userName); err != nil && err != sql.ErrNoRows {
		return user, err
	}
	return user, nil
}
