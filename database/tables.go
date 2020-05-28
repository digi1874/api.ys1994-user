/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-04 13:44:29
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-09 12:52:57
 */

package database

/* User start */

	// User 用户
	type User struct {
		Model
		Name          string `gorm:"comment:'昵称'"`
		UserAvatarID  uint   `gorm:"comment:'头像id'"`
		Birthday      uint   `gorm:"comment:'生日'"`
		Autograph     string `gorm:"comment:'签名'"`
		PlayTime      uint   `gorm:"comment:'播放时长；偷懒一下，用来当等级'"`
		Sex           uint8  `gorm:"DEFAULT:1;comment:'1: 保密；2: 男；3: 女'"`
		State         uint8  `gorm:"DEFAULT:1;comment:'1: 正常；2: 禁用'"`
	}

	// UserAvatar 头像
	type UserAvatar struct {
		Model
		UserID uint   `gorm:"not null;comment:'用户id'"`
		Image  string `gorm:"not null;comment:'头像图片'"`
		UseNum uint   `gorm:"DEFAULT:1;comment:'使用次数'"`
		State  uint8  `gorm:"DEFAULT:1;comment:'1: 正常；2: 禁用'"`
	}

	// UserLogin 用户登录记录
	type UserLogin struct {
		Model
		Signature  string     `gorm:"comment:'JWT的签名'"`
		IP         string     `gorm:"comment:'IP地址'"`
		Country    string     `gorm:"comment:'IP地址地区'"`
		WebsiteID  uint       `gorm:"comment:'网站id'"`
		State      uint8      `gorm:"DEFAULT:1;comment:'1: 正常；2: 退出'"`
	}
/* User end */

/* Vod start */

	// Vod Vod
	type Vod struct {
		Model
		TypeID          uint8       `gorm:"comment:'类型id'"`
		TypePID         uint8       `gorm:"comment:'类型pid'"`
		Name            string      `gorm:"comment:'视频名'"`
		SubName         string      `gorm:"comment:'视频别名'"`
		PY              string      `gorm:"comment:'视频名拼音'"`
		Pic             string      `gorm:"comment:'封面'"`
		Actor           string      `gorm:"comment:'演员'"`
		Director        string      `gorm:"comment:'导演'"`
		Serial          string      `gorm:"comment:'最近更新'"`
		Area            string      `gorm:"comment:'地区'"`
		Lang            string      `gorm:"comment:'语言'"`
		Year            uint16      `gorm:"comment:'年'"`
		Content         string      `gorm:"type:varchar(5000);comment:'绍介'"`
		State           uint8       `gorm:"DEFAULT:1;comment:'1: 正常；2: 禁用'"`
	}

	// VodM3u8 影片m3u8地址
	type VodM3u8 struct {
		Model
		VodID           uint        `gorm:"not null;comment:'视频id'"`
		Name            string      `gorm:"comment:'链接名'"`
		URL             string      `gorm:"comment:'链接地址'"`
	}

	// VodData 影片数据
	type VodData struct {
		Model
		VodID                   uint        `gorm:"not null;comment:'影片id'"`
		LikeCount               uint        `gorm:"DEFAULT:0;comment:'收藏量'"`
		PlayCount               uint        `gorm:"DEFAULT:0;comment:'播放量'"`
		VodStar                 uint8       `gorm:"DEFAULT:0;comment:'影片评分: 0~5'"`
		VodStarCount            uint        `gorm:"DEFAULT:0;comment:'影片评分用户数量'"`
		DirectorStar            uint8       `gorm:"DEFAULT:0;comment:'导演评分'"`
		DirectorStarCount       uint8       `gorm:"DEFAULT:0;comment:'导演评分用户数量'"`
		ActorStar               uint8       `gorm:"DEFAULT:0;comment:'男主评分'"`
		ActorStarCount          uint8       `gorm:"DEFAULT:0;comment:'男主评分用户数量'"`
		ActressStar             uint8       `gorm:"DEFAULT:0;comment:'女主评分'"`
		ActressStarCount        uint8       `gorm:"DEFAULT:0;comment:'女主评分用户数量'"`
		SuppActorStar           uint8       `gorm:"DEFAULT:0;comment:'男配评分'"`
		SuppActorStarCount      uint8       `gorm:"DEFAULT:0;comment:'男配评分用户数量'"`
		SuppActressStar         uint8       `gorm:"DEFAULT:0;comment:'女配评分'"`
		SuppActressStarCount    uint8       `gorm:"DEFAULT:0;comment:'女配评分用户数量'"`
		ScreenplayStar          uint8       `gorm:"DEFAULT:0;comment:'剧本评分'"`
		ScreenplayStarCount     uint8       `gorm:"DEFAULT:0;comment:'剧本评分用户数量'"`
		CinematographyStar      uint8       `gorm:"DEFAULT:0;comment:'摄影评分'"`
		CinematographyStarCount uint8       `gorm:"DEFAULT:0;comment:'摄影评分用户数量'"`
		EditStar                uint8       `gorm:"DEFAULT:0;comment:'剪辑评分'"`
		EditStarCount           uint8       `gorm:"DEFAULT:0;comment:'剪辑评分用户数量'"`
		SoundStar               uint8       `gorm:"DEFAULT:0;comment:'音效评分'"`
		SoundStarCount          uint8       `gorm:"DEFAULT:0;comment:'音效评分用户数量'"`
		VisualStar              uint8       `gorm:"DEFAULT:0;comment:'视觉评分'"`
		VisualStarCount         uint8       `gorm:"DEFAULT:0;comment:'视觉评分用户数量'"`
		MakeupStar              uint8       `gorm:"DEFAULT:0;comment:'化妆评分'"`
		MakeupStarCount         uint8       `gorm:"DEFAULT:0;comment:'化妆评分用户数量'"`
		CostumeStar             uint8       `gorm:"DEFAULT:0;comment:'服装评分'"`
		CostumeStarCount        uint8       `gorm:"DEFAULT:0;comment:'服装评分用户数量'"`
		MusicStar               uint8       `gorm:"DEFAULT:0;comment:'音乐评分'"`
		MusicStarCount          uint8       `gorm:"DEFAULT:0;comment:'音乐评分用户数量'"`
	}

	// VodM3u8Data 影片m3u8地址数据
	type VodM3u8Data struct {
		Model
		VodM3u8ID       uint        `gorm:"not null;comment:'视频id'"`
		PlayCount       uint        `gorm:"DEFAULT:1;comment:'播放量'"`
		LikeCount       uint        `gorm:"DEFAULT:0;comment:'收藏量'"`
	}

	// UserLikeVod 用户收藏影片
	type UserLikeVod struct {
		Model
		VodID      uint       `gorm:"not null;comment:'影片id'"`
		UserID     uint       `gorm:"not null;comment:'账号id'"`
		State      uint8      `gorm:"DEFAULT:1;comment:'1: 收藏；2: 不收藏'"`
	}

	// UserLikeVodM3u8 用户收藏影片
	type UserLikeVodM3u8 struct {
		Model
		VodM3u8ID  uint       `gorm:"not null;comment:'视频id'"`
		UserID     uint       `gorm:"not null;comment:'账号id'"`
		State      uint8      `gorm:"DEFAULT:1;comment:'1: 收藏；2: 不收藏'"`
	}

	// UserPlayM3u8 影片播放记录
	type UserPlayM3u8 struct {
		Model
		UserID           uint        `gorm:"not null;comment:'账号id'"`
		VodID            uint        `gorm:"not null;comment:'影片id'"`
		VodM3u8ID        uint        `gorm:"not null;comment:'链接id'"`
		Time             uint        `gorm:"comment:'播放时间'"`
	}

/* Vod end */

/* UserVodComment start */

	// UserVodComment 影片评论
	type UserVodComment struct {
		Model
		UserID             uint        `gorm:"not null;comment:'账号id'"`
		Grade              uint8       `gorm:"not null;comment:'账号当时等级'"`
		VodID              uint        `gorm:"comment:'影片ID'"`
		Content            string      `gorm:"type:varchar(2000);comment:'内容'"`
		VodStar            uint8       `gorm:"DEFAULT:0;comment:'影片评分'"`
		DirectorStar       uint8       `gorm:"DEFAULT:0;comment:'导演评分'"`
		ActorStar          uint8       `gorm:"DEFAULT:0;comment:'男主评分'"`
		ActressStar        uint8       `gorm:"DEFAULT:0;comment:'女主评分'"`
		SuppActorStar      uint8       `gorm:"DEFAULT:0;comment:'男配评分'"`
		SuppActressStar    uint8       `gorm:"DEFAULT:0;comment:'女配评分'"`
		ScreenplayStar     uint8       `gorm:"DEFAULT:0;comment:'剧本评分'"`
		CinematographyStar uint8       `gorm:"DEFAULT:0;comment:'摄影评分'"`
		EditStar           uint8       `gorm:"DEFAULT:0;comment:'剪辑评分'"`
		SoundStar          uint8       `gorm:"DEFAULT:0;comment:'音效评分'"`
		VisualStar         uint8       `gorm:"DEFAULT:0;comment:'视觉评分'"`
		MakeupStar         uint8       `gorm:"DEFAULT:0;comment:'化妆评分'"`
		CostumeStar        uint8       `gorm:"DEFAULT:0;comment:'服装评分'"`
		MusicStar          uint8       `gorm:"DEFAULT:0;comment:'音乐评分'"`
		Like               uint        `gorm:"DEFAULT:0;comment:'点赞数量'"`
		Dislike            uint        `gorm:"DEFAULT:0;comment:'贬赞数量'"`
		Reply              uint        `gorm:"DEFAULT:0;comment:'回复数量'"`
	}

	// UserVodCommentAt user 收到的 影评@
	type UserVodCommentAt struct {
		Model
		UserID           uint        `gorm:"not null;comment:'账号id'"`
		UserVodCommentID uint        `gorm:"not null;comment:'影片评论id'"`
		Read             uint8       `gorm:"DEFAULT:1;comment:'1:未读；2:已读'"`
	}

	// UserLikeVodComment 影片评论点赞贬赞
	type UserLikeVodComment struct {
		Model
		UserID           uint        `gorm:"not null;comment:'账号id'"`
		UserVodCommentID uint        `gorm:"not null;comment:'评论id'"`
		State            uint8       `gorm:"DEFAULT:0;comment:'1:点赞；2:贬赞；3:取消'"`
	}

	// UserVodCommentReply 影片评论回复
	type UserVodCommentReply struct {
		Model
		UserID           uint        `gorm:"not null;comment:'账号id'"`
		Grade            uint8       `gorm:"not null;comment:'账号当时等级'"`
		UserVodCommentID uint        `gorm:"not null;comment:'影片评论id'"`
		Content          string      `gorm:"type:varchar(400);comment:'内容'"`
		Like             uint        `gorm:"DEFAULT:0;comment:'点赞数量'"`
		Dislike          uint        `gorm:"DEFAULT:0;comment:'贬赞数量'"`
	}

	// UserVodCommentReplyAt user 收到的 影评回复@
	type UserVodCommentReplyAt struct {
		Model
		UserID                uint        `gorm:"not null;comment:'账号id'"`
		UserVodCommentReplyID uint        `gorm:"not null;comment:'影片评论回复id'"`
		Read                  uint8       `gorm:"DEFAULT:1;comment:'1:未读；2:已读'"`
	}

	// UserLikeVodCommentReply 影片评论回复点赞贬赞
	type UserLikeVodCommentReply struct {
		Model
		UserID                uint        `gorm:"not null;comment:'账号id'"`
		UserVodCommentReplyID uint        `gorm:"not null;comment:'评论id'"`
		State                 uint8       `gorm:"not null;comment:'1:点赞；2:贬赞；3:取消'"`
	}

/* UserVodComment end */

/* VodCollection start */

	// UserVodCollection 影片集合
	type UserVodCollection struct {
		Model
		UserID          uint      `gorm:"not null;comment:'用户id'"`
		Name            string    `gorm:"not null;comment:'集合名'"`
		Pic             string    `gorm:"not null;comment:'封面图片'"`
		Content         string    `gorm:"type:varchar(5000);comment:'绍介'"`
		Share           uint8     `gorm:"DEFAULT:1;comment:'1: 分享；2: 不分享'"`
		State           uint8     `gorm:"DEFAULT:1;comment:'1: 正常；2: 禁用'"`
		LikeCount       uint      `gorm:"DEFAULT:0;comment:'收藏量'"`
	}

	// UserVodCollectionVod 集合的影片
	type UserVodCollectionVod struct {
		Model
		UserVodCollectionID   uint      `gorm:"not null;comment:'集合id'"`
		VodID                 uint      `gorm:"not null;comment:'影片id'"`
	}

	// UserLikeVodCollection 收藏影片集合
	type UserLikeVodCollection struct {
		Model
		UserVodCollectionID  uint       `gorm:"not null;comment:'影片集合id'"`
		UserID               uint       `gorm:"not null;comment:'账号id'"`
		State                uint8      `gorm:"DEFAULT:1;comment:'1: 收藏；2: 不收藏'"`
	}

	// UserVodCollectionComment 影片集合评论
	type UserVodCollectionComment struct {
		Model
		UserID              uint        `gorm:"not null;comment:'账号id'"`
		Grade               uint8       `gorm:"not null;comment:'账号当时等级'"`
		UserVodCollectionID uint        `gorm:"not null;comment:'影片集合id'"`
		Content             string      `gorm:"type:varchar(400);comment:'内容'"`
		Like                uint        `gorm:"DEFAULT:0;comment:'点赞数量'"`
		Dislike             uint        `gorm:"DEFAULT:0;comment:'贬赞数量'"`
	}

	// UserVodCollectionCommentAt 影片集合评论@
	type UserVodCollectionCommentAt struct {
		Model
		UserID                     uint        `gorm:"not null;comment:'账号id'"`
		UserVodCollectionCommentID uint        `gorm:"not null;comment:'影片集合评论id'"`
		Read                       uint8       `gorm:"DEFAULT:1;comment:'1:未读；2:已读'"`
	}

	// UserLikeVodCollectionComment 影片集合评论点赞贬赞
	type UserLikeVodCollectionComment struct {
		Model
		UserID                     uint        `gorm:"not null;comment:'账号id'"`
		UserVodCollectionCommentID uint        `gorm:"not null;comment:'评论id'"`
		State                      uint8       `gorm:"not null;comment:'1:点赞；2:贬赞；3:取消'"`
	}
/* VodCollection end */

// autoMigrate 迁移表
func autoMigrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserLogin{})
	DB.AutoMigrate(&UserAvatar{})

	DB.AutoMigrate(&VodData{})
	DB.AutoMigrate(&UserLikeVod{})
	DB.AutoMigrate(&UserPlayM3u8{})
	DB.AutoMigrate(&VodM3u8Data{})
	DB.AutoMigrate(&UserLikeVodM3u8{})

	DB.AutoMigrate(&UserVodComment{})
	DB.AutoMigrate(&UserVodCommentAt{})
	DB.AutoMigrate(&UserLikeVodComment{})

	DB.AutoMigrate(&UserVodCommentReply{})
	DB.AutoMigrate(&UserVodCommentReplyAt{})
	DB.AutoMigrate(&UserLikeVodCommentReply{})

	DB.AutoMigrate(&UserVodCollection{})
	DB.AutoMigrate(&UserVodCollectionVod{})
	DB.AutoMigrate(&UserLikeVodCollection{})
	DB.AutoMigrate(&UserVodCollectionComment{})
	DB.AutoMigrate(&UserVodCollectionCommentAt{})
	DB.AutoMigrate(&UserLikeVodCollectionComment{})
}
