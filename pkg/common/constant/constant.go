package constant

const (

	//friend related
	BlackListFlag         = 1
	ApplicationFriendFlag = 0
	FriendFlag            = 1
	RefuseFriendFlag      = -1

	//Websocket Protocol
	WSGetNewestSeq     = 1001
	WSPullMsgBySeqList = 1002
	WSSendMsg          = 1003
	WSHeartbeat        = 1004
	WSSendSignalMsg    = 1004
	WSPushMsg          = 2001
	WSKickOnlineMsg    = 2002
	WsLogoutMsg        = 2003
	WSDataError        = 3001

	///ContentType
	//UserRelated
	Text           = 101
	Picture        = 102
	Voice          = 103
	Video          = 104
	File           = 105
	AtText         = 106
	Merger         = 107
	Card           = 108
	Location       = 109
	Custom         = 110
	Revoke         = 111
	HasReadReceipt = 112
	Typing         = 113
	Quote          = 114
	Common         = 200
	GroupMsg       = 201

	//SysRelated
	NotificationBegin = 1000

	FriendApplicationApprovedNotification = 1201 //add_friend_response
	FriendApplicationRejectedNotification = 1202 //add_friend_response
	FriendApplicationNotification         = 1203 //add_friend
	FriendAddedNotification               = 1204
	FriendDeletedNotification             = 1205 //delete_friend
	FriendRemarkSetNotification           = 1206 //set_friend_remark?
	BlackAddedNotification                = 1207 //add_black
	BlackDeletedNotification              = 1208 //remove_black

	ConversationOptChangeNotification = 1300 // change conversation opt

	UserNotificationBegin       = 1301
	UserInfoUpdatedNotification = 1303 //SetSelfInfoTip             = 204
	UserNotificationEnd         = 1399
	OANotification              = 1400

	GroupNotificationBegin = 1500

	GroupCreatedNotification             = 1501
	GroupInfoSetNotification             = 1502
	JoinGroupApplicationNotification     = 1503
	MemberQuitNotification               = 1504
	GroupApplicationAcceptedNotification = 1505
	GroupApplicationRejectedNotification = 1506
	GroupOwnerTransferredNotification    = 1507
	MemberKickedNotification             = 1508
	MemberInvitedNotification            = 1509
	MemberEnterNotification              = 1510
	GroupDismissedNotification           = 1511
	GroupMemberMutedNotification         = 1512
	GroupMemberCancelMutedNotification   = 1513
	GroupMutedNotification               = 1514
	GroupCancelMutedNotification         = 1515
	GroupMemberInfoSetNotification       = 1516

	SignalingNotificationBegin = 1600
	SignalingNotification      = 1601
	SignalingNotificationEnd   = 1699

	ConversationPrivateChatNotification = 1701

	OrganizationChangedNotification = 1801

	WorkMomentNotificationBegin = 1900
	WorkMomentNotification      = 1901

	NotificationEnd = 2000

	//status
	MsgNormal  = 1
	MsgDeleted = 4

	//MsgFrom
	UserMsgType = 100
	SysMsgType  = 200

	//SessionType
	SingleChatType = 1
	GroupChatType  = 2

	NotificationChatType = 4
	//token
	NormalToken  = 0
	InValidToken = 1
	KickedToken  = 2
	ExpiredToken = 3

	//MultiTerminalLogin
	//Full-end login, but the same end is mutually exclusive
	AllLoginButSameTermKick = 1
	//Only one of the endpoints can log in
	SingleTerminalLogin = 2
	//The web side can be online at the same time, and the other side can only log in at one end
	WebAndOther = 3
	//The PC side is mutually exclusive, and the mobile side is mutually exclusive, but the web side can be online at the same time
	PcMobileAndWeb = 4

	OnlineStatus  = "online"
	OfflineStatus = "offline"
	Registered    = "registered"
	UnRegistered  = "unregistered"

	//MsgReceiveOpt
	ReceiveMessage          = 0
	NotReceiveMessage       = 1
	ReceiveNotNotifyMessage = 2

	//OptionsKey
	IsHistory                  = "history"
	IsPersistent               = "persistent"
	IsOfflinePush              = "offlinePush"
	IsUnreadCount              = "unreadCount"
	IsConversationUpdate       = "conversationUpdate"
	IsSenderSync               = "senderSync"
	IsNotPrivate               = "notPrivate"
	IsSenderConversationUpdate = "senderConversationUpdate"

	//GroupStatus
	GroupOk              = 0
	GroupBanChat         = 1
	GroupStatusDismissed = 2
	GroupStatusMuted     = 3

	GroupBaned          = 3
	GroupBanPrivateChat = 4

	//UserJoinGroupSource
	JoinByAdmin = 1

	//Minio
	MinioDurationTimes = 3600

	// verificationCode used for
	VerificationCodeForRegister       = 1
	VerificationCodeForReset          = 2
	VerificationCodeForRegisterSuffix = "_forRegister"
	VerificationCodeForResetSuffix    = "_forReset"

	//callbackCommand
	CallbackBeforeSendSingleMsgCommand = "callbackBeforeSendSingleMsgCommand"
	CallbackAfterSendSingleMsgCommand  = "callbackAfterSendSingleMsgCommand"
	CallbackBeforeSendGroupMsgCommand  = "callbackBeforeSendGroupMsgCommand"
	CallbackAfterSendGroupMsgCommand   = "callbackAfterSendGroupMsgCommand"
	CallbackWordFilterCommand          = "callbackWordFilterCommand"
	//callback actionCode
	ActionAllow     = 0
	ActionForbidden = 1
	//callback callbackHandleCode
	CallbackHandleSuccess = 0
	CallbackHandleFailed  = 1

	// minioUpload
	OtherType = 1
	VideoType = 2
	ImageType = 3

	// workMoment permission
	WorkMomentPublic            = 0
	WorkMomentPrivate           = 1
	WorkMomentPermissionCanSee  = 2
	WorkMomentPermissionCantSee = 3

	// workMoment sdk notification type
	WorkMomentCommentNotification = 0
	WorkMomentLikeNotification    = 1
	WorkMomentAtUserNotification  = 2
)
