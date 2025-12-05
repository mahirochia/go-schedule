package schedule

type QueryReq struct {
	UserID int64 `json:"user_id"`
	Year   int32 `json:"year"`
	Month  int32 `json:"month"`
	Day    int32 `json:"day"`
}

type StoreReq struct {
	Year     int32  `json:"year"`
	Month    int32  `json:"month"`
	Day      int32  `json:"day"`
	UserID   int64  `json:"user_id"`
	Content  string `json:"content"`
	Start    string `json:"start"` // 时间字符串格式
	End      string `json:"end"`   // 时间字符串格式
	Priority int    `json:"priority"`
}

type UpdateReq struct {
	ID       int    `json:"id" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	Month    int    `json:"month" binding:"required"`
	Day      int    `json:"day" binding:"required"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Content  string `json:"content"`
	Status   int    `json:"status" binding:"required"`
	UserID   int64  `json:"user_id"`
	Priority int    `json:"priority"`
}
