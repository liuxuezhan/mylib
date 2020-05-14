// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

/*
Package proto_micro is a generated protocol buffer package.

It is generated from these files:
	greeter.proto
	MG_ActivityProtocol.proto
	MG_CommonProtocol.proto
	MG_DBCacheProtocol.proto
	MG_DBProtocol.proto
	MG_GameDefine.proto
	MG_GameProtocol.proto
	MG_GlobalProtocol.proto
	MG_GMProtocol.proto
	MG_LC_OSS_PB.proto
	MG_MailProtocol.proto
	MG_MapEditor.proto
	MG_TranslateProtocol.proto
	MG_UniqueProtocol.proto

It has these top-level messages:
	Request
	Response
	ST_GS2AS_Connect
	ST_AS2GS_ActivityUpdate
	ST_GS2AS_ActivityScore
	ST_MemberJoin_Info
	ST_GS2AS_LeagueJoin
	ST_ActivityInfo_PB
	ST_AS2GS_Activity_InfoList
	ST_GS2AS_ActivityStart
	ST_GS2AS_TickTock
	ST_GS2AS_UpdataStage
	ST_GS2AS_ActivityEnd
	ST_AS2GS_ActivityRank
	ST_AS2GS_Activity_ScoreLevel
	ST_GS2RC_Activity_List
	ST_GS2AS_GM
	ST_Vector2Int_PB
	ST_Vector3Int_PB
	ST_Vector2Int64_PB
	ST_Vector4Int_PB
	ST_Vector4UInt64_PB
	ST_Uint64List_PB
	ST_Int32List_PB
	ST_Uint32List_PB
	ST_Int64List_PB
	ST_ResponPair_PB
	ST_ResponForm_PB
	ST_Vector2IntList_PB
	ST_Vector4UInt64List_PB
	MG_DBC_PING_GS_PONG
	MG_GS_PING_DBC_PONG
	MG_GS2DBC_REQ_CREATE_NEW_CHAR
	MG_GS2DBC_REQ_PULL_CHAR_ASSETS_DATA
	MG_DBC2GS_PUSH_PULL_CHAR_ASSETS_DATA
	MG_GS2DBC_REQ_PULL_WORLDDATA_ASSETS_DATA
	MG_DBC2GS_PUSH_PULL_WORLDDATA_ASSETS_DATA
	MG_DBC2GS_PUSH_PULL_ACTION_DATA
	MG_DBC2GS_PUSH_PULL_ACTION_ASSETS_DATA
	MG_GS2DBC_REQ_PULL_ACTION_ASSETS_DATA
	MG_GS2DBC_REQ_MAIL_INFO
	MG_DBC2GS_PUSH_MAIL_INFO
	MG_GS2DBC_REQ_SAVE_MAIL
	MG_LEAGUE_LIST
	MG_GS2DBC_PULL_ALL_LEAGUE
	MG_GS2DBC_REQ_MODIFY_CHAR_NAME
	MG_GS2DBC_RET_MODIFY_CHAR_NAME
	PB_User
	PB_Player
	PB_ServerList
	ST_CharBasic_Server_PB
	ST_UserServer_PB
	ST_CommanderInfo_PB
	ST_CharBasic_PB
	ST_CharExp_PB
	ST_Statistics_PB
	ST_StrollHeroEventCountDownTimeReward
	ST_StrollHeroEventAddUnitReward
	ST_StrollHeroEventRewardContent_PB
	ST_StrollHeroEventReward_PB
	ST_Building_PB
	ST_BuildingAsset_PB
	ST_LocalActionEntry_PB
	ST_GlobalActionVersion_PB
	ST_GlobalActionVersionList_PB
	ST_LeagueMassJoins
	ST_LeagueMassDetail
	ST_GlobalActionEntry_PB
	ST_LocalActionAsset_PB
	ST_GlobalActionAsset_PB
	ST_GlobalActionList_PB
	ST_ColonyAsset_PB
	ST_DefenseAsset_PB
	ST_CashAsset_PB
	ST_Market_PB
	ST_MarketAsset_PB
	ST_Unit_PB
	ST_UnitAsset_PB
	ST_UnitVague_PB
	ST_UnitVagueAsset_PB
	ST_TeamConf_PB
	ST_TeamConfList_PB
	ST_WorldDataVersion_PB
	ST_WorldDataVersionList_PB
	ST_WorldDataBasic_PB
	ST_WorldDataBasicList_PB
	ST_WorldDataSimpleInfo_PB
	ST_CashHouse_PB
	ST_WorldDataDetail_Mine_PB
	ST_ColonyCapture_PB
	ST_WorldDataDetail_Colony_PB
	ST_WorldDataDetail_NoUse_PB
	ST_WorldDataDetail_Bestationed_PB
	ST_WorldDataDetail_MutiMonster_PB
	ST_WorldDataDetail_City_PB
	ST_WorldDataDetail_MainCity_PB
	ST_WorldDataDetail_UndergroundPalace_PB
	ST_CityArchon_PB
	ST_CityOccupyInfo_PB
	ST_SC_CityWarList_PB
	ST_SC_CityWarInfo_PB
	ST_CityWar_PB
	ST_WorldDataDetailInfo_PB
	ST_WorldDataEntry_PB
	ST_WorldDataEntryList_PB
	ST_Research_PB
	ST_ResearchEntryList_PB
	ST_ItemEntry_PB
	ST_Asset_ItemPack_PB
	ST_Asset_Chat_Item_PB
	ST_ForbiddenChat_PB
	ST_PrivateChat_PB
	ST_PrivateChat_Content_PB
	ST_Asset_Chat_PB
	ST_ActiveSpy_Lord
	ST_GarrisonInfo_PB
	ST_ActiveSpy_Reinforcements
	ST_ActiveSpy_Cash
	ST_ActiveSpy_Vague
	ST_ActiveSpy_Exact
	ST_ActiveSpy_Hardiness
	ST_Radar_TechInfo
	ST_Radar_NaturalGift
	ST_Simple_Statue
	ST_Radar_Statue
	ST_ArmyHerosDetail
	ST_ActiveSpyReport_PB
	ST_ActiveSpyMonsterReport_PB
	ST_MonsterSpy_Vague
	ST_MonsterSpy_Exact
	ST_EarlyWarning_Commander
	ST_EarlyWarning_Hero
	ST_EarlyWarning_Vague
	ST_EarlyWarning_Exact
	ST_EarlyWarning_StartInfo
	ST_EarlyWarning_EndInfo
	ST_EarlyWarningReportMail_PB
	ST_BaseEarlyWarningReport_PB
	ST_MonsterAtkCityWarning_PB
	ST_EarlyWarningEx_PB
	ST_EarlyWarningGreen_PB
	ST_UnitTotal_PB
	ST_EarlyWarningTime_PB
	ST_MultiplayerEarlyWarning_PB
	ST_RadarAsset_PB
	ST_CollectReport_PB
	ST_LeagueTransportReport_PB
	ST_LeagueInvitationJoinMail_PB
	ST_MailBasicInfo_PB
	ST_MailRewardList_PB
	ST_Vector4IntList_PB
	ST_MailSimple_PB
	ST_MailData_PB
	ST_MailAsset_PB
	ST_MailAssetForClientSave_PB
	ST_MailDBAsset_PB
	ST_MailStorage_PB
	ST_LeagueDepotUsedResource_PB
	ST_LeagueInvitationReject_PB
	ST_LeagueAsset_PB
	KeyBoxInfo
	ST_LeagueHelpBaisc_PB
	ST_Gift_PB
	ST_LeagueExtra_PB
	ST_LeagueBattleNOTJoin_PB
	ST_LeagueBattle_PB
	ST_League_Member_Info
	ST_League_Simple_Info
	ST_List_Uint64
	ST_League_Operation_Result
	ST_League_Simple_Info_List
	ST_LeagueTechnology_PB
	ST_League_Help_Simple_PB
	ST_LeagueHelpDetail_PB
	ST_League_Help_Info_PB
	ST_League_Help_Asset_PB
	ST_LeagueGift_PB
	ST_LeagueGiftList_PB
	ST_LeagueGiftSpecial_PB
	ST_LeagueMemberTitleName_PB
	ST_LeagueDepot_PB
	ST_LeagueManagement_PB
	ST_League_Detail_Info
	ST_League_Member_BuffTitle
	ST_LeagueCampLeader_PB
	ST_LeagueBox_List_PB
	ST_InvitationLeague_Info
	ST_PopUpAsset_PB
	ST_LeagueMessageEntry_PB
	ST_UnitCount_PB
	ST_BuffBattleInfo_PB
	ST_HeroSkillInfo_PB
	ST_HeroBattleInfo_PB
	ST_CorpBattleInfo_PB
	ST_RoundBaseInfo_PB
	ST_BattleRoundInfo_PB
	ST_ShowUnit_PB
	ST_SimpleStartInfo_PB
	ST_BattleInfo_PB
	ST_BattleCommanderSimpleInfo_PB
	ST_HeroBattleReport_PB
	ST_GetExp_PB
	ST_UnitBattleReport_PB
	ST_WallReport_PB
	ST_OtherAddition_PB
	ST_BattleReportDetail_PB
	ST_PowerReport_PB
	ST_DefFailInfo_PB
	ST_BattleReport_PB
	ST_ReturnUnits_PB
	ST_MonsterInfo_PB
	ST_MonsterKill_PB
	ST_BattleNotice_PB
	ST_BattleNoticesAsset_PB
	ST_TileInfo_PB
	ST_TileInfo_List
	ST_Chat_Entry_PB
	ST_Chat_Title_PB
	ST_Chat_League_PB
	ST_Chat_TitleList_PB
	ST_Chat_EntryList_PB
	ST_Chat_ChannelInfo_PB
	ST_WorldChatTitleInfo_PB
	ST_Chat_ReplyList_PB
	ST_Hot_Chat_Item_PB
	ST_Hot_Chat_Info_PB
	ST_Chat_ChannelInfoList_PB
	ST_Chat_TitleInfo_PB
	ST_Chat_TitleInfoList_PB
	ST_HeroEntry_PB
	ST_HeroSkill_PB
	ST_Prison_PB
	ST_PrisonAsset_PB
	ST_HeroAsset_PB
	ST_Buff_PB
	ST_BuffAsset_PB
	ST_GlobalBuff_PB
	ST_TalentAsset_PB
	ST_Skill_PB
	ST_SkillAsset_PB
	ST_Tips_PB
	ST_TipsAsset_PB
	ST_TaskAsset_PB
	ST_Vip_PB
	ST_Attribute_PB
	ST_AttributeList_PB
	ST_Asset_CommanderEquipInfo
	ST_WearEquip
	ST_Asset_FactoryPackage
	ST_EquipBasic
	ST_MaterialBase
	ST_Asset_TakeCard
	ST_TakeCard_Basic
	ST_Client_CardPool
	ST_CardPoolBasic
	ST_GiftHistory
	ST_Asset_OfferGift
	ST_OfferGiftPack
	ST_OfferGiftPackConf
	ST_Asset_GiftPack
	ST_GiftPackBasic
	ST_Asset_MonthGiftPack
	ST_SubscribeGiftPack
	ST_CharSimpleInfo
	ST_CharSimpleInfoList
	ST_WorldCharSimple_PB
	ST_WorldCharSimpleList_PB
	ST_PowerRankingListItem
	ST_GlobalRankingItem
	ST_RankingHeroInfo
	ST_PowerRankingList
	ST_RecommendUserAsset_PB
	ST_LeagueDonationRankEntry_PB
	ST_LeagueDonationRank_PB
	ST_LeagueStatusEntry_PB
	ST_LeagueStatus_PB
	ST_AchievementFinishedItem_PB
	ST_AchievementAsset_PB
	ST_LeagueCustomFlag_PB
	ST_MainCityEvent_PB
	ST_MainCityEvent_Server_Asset_PB
	ST_MainCityEventAsset_PB
	ST_DispatchTaskEntry_PB
	ST_DispatchTaskAsset_PB
	ST_Statue_PB
	ST_StatueAsset_PB
	ST_SignInAsset_PB
	ST_WorldCity_PB
	ST_WorldCityList_PB
	ST_TranslateResult_PB
	ST_RoleInformation_PB
	ST_GlobalMessage_PB
	ST_GlobalMessageList_PB
	ST_StrollHeroEventAsset_PB
	ST_StrollHeroEvent_PB
	ST_LeagueShopRecordList_PB
	ST_LeagueShopRecord_PB
	ST_EveryoneMailItem
	ST_EveryoneMailCondition
	ST_EveryoneMailContent
	ST_SettingAsset_PB
	ST_PowerChat_PB
	ST_AudioSetting_PB
	ST_NotifictionSetting_PB
	ST_FeedbackMark
	ST_UserSettingsAlterParams
	ST_BeginnerTutorialAsset_PB
	ST_FirstTutorialAsset_PB
	ST_SetMemberTitles_PB
	ST_Activity_PB
	ST_ActivityScore_PB
	ST_ActivityLeague_Task
	ST_ActivityLeague_TaskAction
	ST_MemberInfo_Base
	ST_ActivityScore_User
	ST_ActivityScore_LeagueUser
	ST_ActivityScore_UserTask
	ST_LeagueInfo_Base
	ST_ActivityScore_Base
	ST_ActivityScore_League
	ST_ActivityDetail_PB
	ST_ActivityContent_PB
	ST_ActivityRank_League
	ST_ActivityRank_LeagueUser
	ST_ActivityRank_PB
	ST_Asset_Activity_PB
	ST_Asset_Activity_Item_PB
	ST_MacActifityInfo
	ST_FoolsDay
	ST_FoolsDay_Rewards
	ST_DailyTask_PB
	ST_Lottery_PB
	ST_BookmarkEntry_PB
	ST_Asset_Bookmark_PB
	ST_MiracleWarRecord_PB
	ST_HistoryArchon_PB
	ST_OccupyInfo_PB
	ST_Title_PB
	ST_WarGift_PB
	ST_MiracleWarBuff_PB
	ST_MiracleWarRight
	ST_GlobalNotice
	ST_MiracleWar_PB
	ST_MiracleWarRecords_PB
	ST_MiracleWarHistoryArchons_PB
	ST_Asset_MiracleWar_PB
	ST_Server_PB
	ST_ServerList_PB
	ST_Asset_Account_PB
	ST_League_IndependData_Version_PB
	ST_League_Message_List_PB
	ST_BlackMarket_Item_PB
	ST_Asset_BlackMarket_PB
	ST_Asset_Supply_PB
	ST_ResourceObtain_Supply_PB
	ST_LeagueBox_Item_PB
	ST_Asset_LeagueBox_PB
	ST_ChallengeTarget_PB
	ST_ChallengeCondition_PB
	ST_ChallengeTast_PB
	ST_Asset_Challenge_PB
	ST_Asset_LuckyDog
	ST_Asset_AirdropInfo
	ST_WorldDomination_PB
	ST_Asset_MapList_PB
	ST_MapList_Info_PB
	ST_Chapter_PB
	ST_StarBox_PB
	ST_MapInfo_PB
	ST_MopUpRewards_PB
	ST_ActivityTask_Item
	ST_Asset_ActivityTask_PB
	ST_MainCitySkinInfo_PB
	ST_Asset_MainCitySkin_PB
	ST_Asset_BoonCenter_PB
	ST_GiftActivity
	ST_GiftActivity_One
	ST_BoonPayFirst_PB
	ST_BoonPaySecond_PB
	ST_BoonDuration_PB
	ST_BoonDurationData_PB
	ST_BoonDaily_PB
	ST_BoonDailyData_PB
	ST_BoonFund_PB
	ST_BoonFundData_PB
	ST_BoonRewardData_PB
	ST_BoonRewardConfig_PB
	ST_ChapterTask_PB
	ST_ChapterInfo_PB
	ST_Asset_Chapters_PB
	ST_PrisonMessage_PB
	ST_PrisonMessageBoard_PB
	ST_PrisonMessgeClient_PB
	ST_PrisonMessageList_PB
	ST_CommanderFashionInfo_PB
	ST_Asset_CommanderFashion_PB
	ST_SevenDayTask_Item
	ST_SeventDayTaskAsset_PB
	ST_NoviceTaskAsset_PB
	ST_KillWildMonster_PB
	ST_AirShipChanllange_PB
	AirShipPassDetail
	ST_UserAirShipChanllange_PB
	ST_AirShipChanllangeRet_PB
	ST_DropActivityRankItem_PB
	ST_DropActivityRank_PB
	ST_LeagueTask_Item
	ST_Asset_LeagueTask_PB
	ST_SubTask
	ST_StrongestCommander_PB
	ST_StrongestCommanderRank_PB
	ST_Asset_StrongestCommander_PB
	ST_PrepareForWar_PB
	ST_PrepareForWarRank_PB
	ST_Asset_PrepareForWar_PB
	ST_ExpeditionShop_Item_PB
	ST_Asset_ExpeditionShop_PB
	ST_ExpeditionShop_PB
	ST_WorldSituation_Entry_PB
	ST_WorldSituation_PB
	ST_WorldSituationRankSnapshot_PB
	ST_WorldSituationRankSnapshot_Entry_PB
	ST_Asset_WorldSituation_PB
	ST_Asset_LoginGift_PB
	ST_LuckyStarItem_PB
	ST_LuckyStar_PB
	ST_LuckyStarShop_PB
	ST_LuckyStarItem_Server_PB
	ST_LuckyStar_Server_PB
	ST_Asset_LuckyStar_PB
	ST_CharSimple_Server_PB
	ST_CharSimpleGarrison_PB
	ST_SyncAsset_Entry
	ST_SyncAsset_PB
	ST_MailTypeList_Entry
	ST_MailTypeList
	ST_MailAssetSummary_Entry
	ST_MailAssetSummary
	ST_MailDB_PB
	ST_ServerRank
	ST_GetRoleRst
	ST_LeagueBattle_User
	ST_LeagueBattle_Log_Entry
	ST_LeagueBattle_Log
	ST_MonsterAtkCity
	ST_GS2LS_UpdateRoleReq
	ST_RoleInfo
	ST_MiracleWarList_PB
	ST_ChatChatroom
	ST_ChatInfo
	ST_Chat_NeteaseRes
	ST_Transfer_Battle_PB
	ST_Prepare_Server_Info
	ST_Prepare_Detail
	ST_Processing_Detail
	ST_Recovering_Detail
	ST_Transfer_Battle_SubTask
	ST_Prepare_User_Info
	ST_Prepare_League_Info
	ST_Asset_TransferBattle_PB
	ST_Client_TransferBattle_Info
	ST_Client_TransferBattle_User_Rank
	ST_Client_TransferBattle_League_Rank
	ST_CollectResource_Entry
	ST_CollectResources
	ST_Sort_PB
	ST_Rank_PB
	ST_InitRanking_Req_PB
	ST_InitRanking_Respon_PB
	ST_GS2Global_Connect_Req_PB
	ST_GS2Global_Connect_Respon_PB
	ST_UpdateRankingScore_Req_PB
	ST_UpdateRankingScore_Respon_PB
	ST_RealTimeRanking_Req_PB
	ST_RealTimeRanking_Respon_PB
	ST_RankingScore_PB
	ST_GetPlayerInfo_Req_PB
	ST_GetPlayerInfo_Respon_PB
	ST_PushRanking_Req_PB
	ST_PushRanking_Respon_PB
	ST_GetRanking_Req_PB
	ST_GetRanking_Respon_PB
	ST_GM_CharBasic_PB
	ST_GM_Buildings_PB
	ST_GM_Building_PB
	ST_GM_Power_PB
	ST_AccountInfo_PB
	ST_AccountInfos_PB
	ST_GM_Forbidden_PB
	ST_GM_Resource_PB
	ST_GM_ServerStatus_PB
	ST_Get_List_Status
	ST_Get_List_Status_Result
	ST_Open_List
	ST_Open_List_Result
	ST_CS2LS_Add_List
	ST_LS2CS_Add_List_Result
	ST_CS2LS_Remove_List
	ST_LS2CS_Remove_List_Result
	ST_LS2CS_Alloc_Server
	ST_CS2LS_Alloc_Server
	ST_GS2CS_Register
	ST_Ping
	ST_CS2GS_GMCommand
	ST_GS2CS_GMCommand_Result
	ST_CS2GS_Pay
	ST_GS2CS_Pay
	ST_EveryOneMail
	ST_ViewServer_Request
	ST_ViewServer_Response
	ST_ChangeServer_Request
	ST_ChangeServer_Response
	ST_GS2CS_ServerInfo
	ST_GS2CS_GetGiftPackConfigs_Request
	ST_CS2GS_GetGiftPackConfigs_Response
	ST_GS2CS_GetOfferGift_Request
	ST_CS2GS_GetOfferGift_Response
	ST_GS2CS_GetLuckyStar_Request
	ST_CS2GS_GetLuckyStar_Response
	ST_GiftPackConfig
	ST_GiftPackCondition
	ST_CS2GS_Sync_ServerList
	ST_MailDocData_PB
	ST_SyncAccountDetailsRequest
	ST_AccountInfoPair
	ST_SyncAccountDetailsResponse
	ST_UpdateAccountDetailsRequest
	ST_UpdateAccountDetailsResponse
	ST_QueryAccountBoundRequest
	ST_QueryAccountBoundResponse
	ST_SwitchAccountRequest
	ST_SwitchAccountResponse
	GM_StartNewActivity_PB
	ST_GM2AS_ActivityView
	GM_ActivityInfo_PB
	ST_AS2GM_ResultList
	ST_GM2AS_ActivityReserve
	ST_AS2GM_ReserveResult
	ST_ModifyArchonNotice_Request
	ST_GM_ArchonNoticeInfo
	ST_GM_BoonRewardInfo
	ST_GM_BoonActivityOper
	ST_GM_BoonActivities
	ST_GS2LS_GetRoleReq
	ST_GS2LS_UpdateRoleRst
	ST_GS2LS_CreateRoleReq
	ST_GS2LS_CreateRoleRst
	ST_GS2LS_ReportOnlineReq
	ST_GS2LS_ReportOnlineRst
	ST_GS2LS_ServerListReq
	ST_ServerListInfo
	ST_ServerList
	ST_SwitchRoleRst
	ST_UserSid
	ST_TransferBattle_Group_Entry
	ST_TransferBattle_Group
	ST_GS2CS_GetTransferBattleInfo_PB
	ST_CS2GS_GetTransferBattleInfo_Response
	ST_GetTransferBattlePrepareInfo
	ST_TransferBattlePrepare_Response
	ST_GS2CS_SyncTransferBattleScore_PB
	ST_AddTransferBattleScore_Request
	ST_AddTransferBattleScore_Response
	ST_TBTransferServer_Request
	ST_TBTransferServer_Response
	ST_SyncTBMiracleWar_Request
	ST_SyncTBMiracleWar_Response
	ST_GS2CS_SyncTBStage_PB
	ST_StrongestCommander_Entry
	ST_GS2CS_GetStrongestCommanderInfo_PB
	ST_CS2GS_GetStrongestCommanderInfo_Response
	ST_GS2CS_SyncSCStage_PB
	MG_LC_OSS_PB
	ST_UpsertMail_Req_PB
	ST_UpsertMail_Respon_PB
	ST_GetMail_Req_PB
	ST_GetMail_Respon_PB
	ST_MapData_PB
	ST_TileData_PB
	ST_BlockDecoList_PB
	ST_DecoData_PB
	ST_WorldCityData_PB
	ST_WorldMapData_PB
	ST_TranslationItem_PB
	ST_Translation_PB
	ST_Translation_Req_PB
	ST_Translation_Respon_PB
	ST_GS2US_UserName_Check
	ST_GS2US_UserName_Change
	ST_GS2US_LeagueName_Check
	ST_GS2US_LeagueName_Change
	ST_GS2US_LeagueName_Delete
	ST_LeagueName_PB
	ST_GS2US_NameList
	ST_GS2US_ServerCheck
*/
package proto_micro

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "proto_micro.Request")
	proto.RegisterType((*Response)(nil), "proto_micro.Response")
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x53, 0xf1, 0xb9, 0x99,
	0xc9, 0x45, 0xf9, 0x4a, 0xb2, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x07, 0xd8,
	0xa0, 0xcc, 0xbc, 0x74, 0x09, 0x26, 0xb0, 0x1a, 0x38, 0xdf, 0xc8, 0x91, 0x8b, 0xdd, 0x1d, 0x62,
	0x89, 0x90, 0x19, 0x17, 0xab, 0x47, 0x6a, 0x4e, 0x4e, 0xbe, 0x90, 0x88, 0x1e, 0x92, 0x45, 0x7a,
	0x50, 0x5b, 0xa4, 0x44, 0xd1, 0x44, 0x21, 0x86, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0xc5, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0x1c, 0x1a, 0x75, 0xae, 0x00, 0x00, 0x00,
}
