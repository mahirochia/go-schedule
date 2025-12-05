package spider

type Request struct {
	BaseReq   BaseReq `json:"base_req"`
	Forward   string  `json:"forward"`
	FlushNum  int     `json:"flush_num"`
	ChannelID string  `json:"channel_id"`
	ItemCount int     `json:"item_count"`
}

type BaseReq struct {
	From string `json:"from"`
}

// 外层请求结构
type NewsResponse struct {
	Timestamp string     `json:"timestamp"`
	TraceID   string     `json:"trace_id"`
	Data      []NewsItem `json:"data"`
}

// 新闻条目
type NewsItem struct {
	ID              string          `json:"id"`
	Articletype     string          `json:"articletype"`
	Title           string          `json:"title"`
	PicStyle        int             `json:"pic_style"`
	PublishTime     string          `json:"publish_time"`
	PicInfo         PicInfo         `json:"pic_info"`
	LinkInfo        LinkInfo        `json:"link_info"`
	MediaInfo       MediaInfo       `json:"media_info"`
	InteractionInfo InteractionInfo `json:"interation_info"` // 注意：原字段拼写错误 "interation_info"
	DocInfo         DocInfo         `json:"doc_info"`
	ShortTitle      string          `json:"short_title"`
	UpdateTime      string          `json:"update_time"`
	News724Feature  interface{}     `json:"news724_feature"` // 可为 null
	RcmInfo         RcmInfo         `json:"rcm_info"`
	Desc            string          `json:"desc"`
	SafeControl     SafeControl     `json:"safe_control"`
	Category        Category        `json:"category"`
	Declare         interface{}     `json:"declare"` // 空对象 {}
	PubInfo         PubInfo         `json:"pub_info"`
	PaymentInfo     interface{}     `json:"payment_info"` // 空对象 {}
	UserAddress     string          `json:"user_address"`
	DislikeOptions  []DislikeOption `json:"dislike_options"`
	CpInfo          CpInfo          `json:"cp_info"`
	ExtraProperty   interface{}     `json:"extra_property"` // 空对象 {}
	ShareInfo       ShareInfo       `json:"share_info"`
	LongSummary     string          `json:"long_summary"`
}

// 图片信息
type PicInfo struct {
	BigImg       []string          `json:"big_img"`
	SmallImg     []string          `json:"small_img"`
	ThreeImg     []string          `json:"three_img"`
	LsImgExpType string            `json:"ls_img_exp_type"`
	Ext          map[string]string `json:"ext"` // 动态键如 "196x130"
	ShareImg     string            `json:"share_img"`
	ImgCount     int               `json:"img_count"`
}

// 链接信息
type LinkInfo struct {
	ShareURL string `json:"share_url"`
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
	OrgURL   string `json:"org_url"`
}

// 媒体信息
type MediaInfo struct {
	ChlID         string    `json:"chl_id"`
	ChlName       string    `json:"chl_name"`
	Icon          string    `json:"icon"`
	Sicon         string    `json:"sicon"`
	Mrk           string    `json:"mrk"`
	LastArtUpdate string    `json:"last_art_update"`
	Uin           string    `json:"uin"`
	EncodedSuid   string    `json:"encoded_suid"`
	VipType       string    `json:"vip_type"`
	VipTypeNew    string    `json:"vip_type_new"`
	VipDesc       string    `json:"vip_desc"`
	VipIcon       string    `json:"vip_icon"`
	MedalInfo     MedalInfo `json:"medal_info"`
}

// 勋章信息
type MedalInfo struct {
	TypeID     int    `json:"type_id"`
	MedalID    int    `json:"medal_id"`
	MedalLevel int    `json:"medal_level"`
	MedalName  string `json:"medal_name"`
	MedalDesc  string `json:"medal_desc"`
	NightURL   string `json:"night_url"`
	DaytimeURL string `json:"daytime_url"`
}

// 互动信息（注意原字段名拼写为 interation_info）
type InteractionInfo struct {
	CommentID      string `json:"comment_id"`
	CommetNum      int    `json:"commet_num"` // 注意：可能是 comment_num 的拼写错误
	ReadNum        int    `json:"read_num"`
	LikeNum        int    `json:"like_num"`
	CollectNum     int    `json:"collect_num"`
	ShareNum       int    `json:"share_num"`
	ShareWechatNum int    `json:"share_wechat_num"`
}

// 文档分类信息
type DocInfo struct {
	FirstCateID   string `json:"first_cate_id"`
	FirstCateName string `json:"first_cate_name"`
	SecCateID     string `json:"sec_cate_id"`
	SecCateName   string `json:"sec_cate_name"`
}

// 推荐信息
type RcmInfo struct {
	ReasonInfo     string `json:"reason_info"`
	ReasonFlag     string `json:"reason_flag"`
	ReasonFlagList string `json:"reason_flag_list"`
}

// 安全控制
type SafeControl struct {
	CloseAllAd              int    `json:"close_all_ad"`
	CloseAllFavorite        int    `json:"close_all_favorite"`
	CloseAllRel             int    `json:"close_all_rel"`
	CloseSharePull          int    `json:"close_share_pull"`
	CloseRelateThing        int    `json:"close_relate_thing"`
	CloseAllEmoticonComment int    `json:"close_all_emoticon_comment"`
	PoliticalOption         string `json:"political_option"`
	PCNotDisplay            string `json:"pc_not_display"`
}

// 分类信息
type Category struct {
	Cate1Name   string `json:"cate1_name"`
	Cate1EnName string `json:"cate1_en_name"`
	Cate1ID     string `json:"cate1_id"`
	Cate2Name   string `json:"cate2_name"`
	Cate2ID     string `json:"cate2_id"`
}

// 发布信息
type PubInfo struct {
	Source    string `json:"source"`
	SubSource string `json:"sub_source"`
}

// 不喜欢选项
type DislikeOption struct {
	MenuID    string     `json:"menu_id"`
	MenuName  string     `json:"menu_name"`
	MenuItems []MenuItem `json:"menu_items"`
}

type MenuItem struct {
	Type string `json:"type"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

// CP 信息
type CpInfo struct {
	TnewsSelfMade int `json:"tnews_self_made"`
}

// 分享信息
type ShareInfo struct {
	ShareTitle    string `json:"share_title"`
	ShareSubtitle string `json:"share_subtitle"`
}
