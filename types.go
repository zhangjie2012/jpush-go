package jpush

var (
	JPushPushUrl = "https://api.jpush.cn/v3/push"
)

type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	Abtest         []string `json:"abtest,omitempty"`
}

type IOSBase struct {
	Alert            interface{} `json:"alert"`                       // 通知内容
	Sound            interface{} `json:"sound,omitempty"`             // 通知提示声音或警告通知
	Badge            int         `json:"badge,omitempty"`             // 应用角标
	ContentAvailable bool        `json:"content-available,omitempty"` // 推送唤醒
	MutableContent   bool        `json:"mutable-content,omitempty"`   // 通知扩展
	Category         string      `json:"category,omitempty"`          //
	Extras           interface{} `json:"extras,omitempty"`            // 附加字段
	ThreadId         string      `json:"thread-id,omitempty"`         // 通知分组
}

type Notification struct {
	AIOpportunity bool        `json:"ai_opportunity"`     // 是否采用"智能时机"策略下发通知
	Alert         string      `json:"alert,omitempty"`    // 基本通知
	Android       interface{} `json:"android,omitempty"`  // TODO
	IOS           IOSBase     `json:"ios,omitempty"`      // ios
	Quickapp      interface{} `json:"quickapp,omitempty"` // TODO
	Winphone      interface{} `json:"winphone,omitempty"` // TODO
}

type Message struct {
	MsgContent  string      `json:"msg_content,omitempty"`
	Title       string      `json:"title,omitempty"`
	ContentType string      `json:"content_type,omitempty"`
	Extras      interface{} `json:"extras,omitempty"`
}

type InappMessage struct {
	InappMessage bool `json:"inapp_message"`
}

type Options struct {
	SendNo            int         `json:"sendno,omitempty"`              // 推送序号
	TimeToLive        int         `json:"time_to_live,omitempty"`        // 离线消息保留时长(秒)
	OverrideMsgId     int64       `json:"override_msg_id,omitempty"`     // 要覆盖的消息 ID
	ApnsProduction    bool        `json:"apns_production"`               // APNs 是否生产环境
	ApnsCollapseId    string      `json:"apns_collapse_id,omitempty"`    // 更新 iOS 通知的标识符
	BigPushDuration   int         `json:"big_push_duration,omitempty"`   // 定速推送时长(分钟)
	ThirdPartyChannel interface{} `json:"third_party_channel,omitempty"` // TODO
}

type PushBody struct {
	Platform        interface{}   `json:"platform"`                   // 平台: "all" | ["android", "ios", "quickapp", "winphone"]
	Audience        *Audience     `json:"audience"`                   // 推送设备(目标)
	Notification    *Notification `json:"notification,omitempty"`     // 通知内容体, 与 message 必有其一
	Message         *Message      `json:"message,omitempty"`          // 通知消息体, 与 notification 必有其一, 应用内消息
	InappMessage    *InappMessage `json:"inapp_message,omitempty"`    // 应用内提醒
	Options         *Options      `json:"options,omitempty"`          // 推送参数
	Notification3rd interface{}   `json:"notification_3rd,omitempty"` // TODO 自定义消息转厂商通知内容体。与 message 一起使用
	SmsMessage      interface{}   `json:"sms_message,omitempty"`      // TODO 短信渠道补充送达内容体
	Cid             string        `json:"cid,omitempty"`              // CID
}

type PushResp struct {
	SendNo string `json:"sendno"`
	MsgId  string `json:"msg_id"`
}
