package biz

type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

type AppealParam struct {
	ReviewID  int64
	StoreID   int64
	Reason    string
	Content   string
	PicInfo   string
	VideoInfo string
}
