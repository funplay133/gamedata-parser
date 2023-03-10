package dataparse

import (
	"dataparse/dump"
)

var NewAssets = map[string][]dump.Asset{
	"6029408":    {dump.AssetName["ScDLCAvatarLevelMetaData"]},
	"11165326":   {dump.AssetName["OWActivityBossMetaData"]},
	"11507150":   {dump.AssetName["RandomPlotDataMetaData"]},
	"22366012":   {dump.AssetName["MonsterCardStarMetaData"]},
	"36243594":   {dump.AssetName["WorldMapEntryMetaData"]},
	"44405226":   {dump.AssetName["SLGBattlePointMetaData"]},
	"50428269":   {dump.AssetName["HybridSiteDataMetaData"]},
	"53370678":   {dump.AssetName["BossRushBuffMetaData"]},
	"112154430":  {dump.AssetName["JigsawGroupMetaData"]},
	"121386631":  {dump.AssetName["AvatarTrialMetaData"]},
	"125371850":  {dump.AssetName["ShopGoodsMetaData"]},
	"137850209":  {dump.AssetName["ScDLCTalentDataMetaData"]},
	"150008597":  {dump.AssetName["GeneralActivityDisplayDataMetaData"]},
	"169620141":  {dump.AssetName["GeneralScheduleMetaData"]},
	"175654298":  {dump.AssetName["GodWarEventFlagMetaData"]},
	"175952655":  {dump.AssetName["GodWarEventPlotMetaData"]},
	"201380970":  {dump.AssetName["StageDetailMonsterMetaData"]},
	"208715169":  {dump.AssetName["SinDemonExScoreRewardMetaData"]},
	"213532048":  {dump.AssetName["RpgSurvivalTraitMetaData"]},
	"220083180":  {dump.AssetName["WarehouseRequireData"]},
	"226923750":  {dump.AssetName["TowerRaidBuffMetaData"]},
	"237914729":  {dump.AssetName["SinDemonExSkillTipsMetaData"]},
	"244665338":  {dump.AssetName["ScratchTicketPlateDataMetaData"]},
	"250894170":  {dump.AssetName["BattlePassSeasonLevelConfigMetaData"]},
	"285372794":  {dump.AssetName["GreedyEndlessFloorConfigMetaData"]},
	"295065595":  {dump.AssetName["CustomHeadMetaData"]},
	"319085691":  {dump.AssetName["SignInRewardMetaData"]},
	"326943498":  {dump.AssetName["GlobalExploreEventMetaData"]},
	"337996114":  {dump.AssetName["GlobalExploreQuestMetaData"]},
	"339809134":  {dump.AssetName["GlobalExploreStageMetaData"]},
	"354247105":  {dump.AssetName["UnknownAsset1MetaData"]},
	"364255081":  {dump.AssetName["TouchMetaData"]},
	"376306182":  {dump.AssetName["OWEndlessBattleConfig"]},
	"379497939":  {dump.AssetName["AvatarFragmentMetaData"]},
	"394220128":  {dump.AssetName["BattlePassScheduleMetaData"]},
	"402624629":  {dump.AssetName["ChapterOWHeroLevelMetaData"]},
	"404941629":  {dump.AssetName["OpenWorldLocation"]},
	"408619454":  {dump.AssetName["EquipmentSetMetaData"]},
	"411717359":  {dump.AssetName["RogueTowerStageMetaData"]},
	"420652155":  {dump.AssetName["ChapterOWBuildingLevelMetaData"]},
	"438979584":  {dump.AssetName["DormitoryAvatarMetaData"]},
	"462670926":  {dump.AssetName["GreedyEndlessBattleConfigMetaData"]},
	"482711261":  {dump.AssetName["HoukaiTownBuffMetaData"]},
	"483009743":  {dump.AssetName["StageDetailMetaData"]},
	"500858231":  {dump.AssetName["ChapterMetaData"]},
	"500991781":  {dump.AssetName["GeneralActivityScoreDataMetaData"]},
	"550303508":  {dump.AssetName["AvatarSubSkillMetaData"]},
	"604807223":  {dump.AssetName["ChapterOWQuestDataMetaData"]},
	"606639383":  {dump.AssetName["DormitoryFurnitureMetaData"]},
	"609434632":  {dump.AssetName["TutorialData"]},
	"612893459":  {dump.AssetName["StigmataPositionMetaData"]},
	"619916865":  {dump.AssetName["DormitoryDecorationEffectMetaData"]},
	"623672237":  {dump.AssetName["ChapterOWTipsMetaData"]},
	"650609582":  {dump.AssetName["RpgSimpleBuffMetaData"]},
	"686454746":  {dump.AssetName["MonsterCardLevelMetaData"]},
	"693085863":  {dump.AssetName["MonsterCardSkillMetaData"]},
	"693346132":  {dump.AssetName["MonsterCardStageMetaData"]},
	"723318358":  {dump.AssetName["StageMaxScoreMetaData"]},
	"738857319":  {dump.AssetName["GeneralActivityStageMetaData"]},
	"754937353":  {dump.AssetName["LoadingTipMetaData"]},
	"761861209":  {dump.AssetName["GodWarRoleSkillDataMetaData"]},
	"766414602":  {dump.AssetName["StageRechallengeMetaData"]},
	"783766844":  {dump.AssetName["RpgTowerV2MetaData"]},
	"788785795":  {dump.AssetName["StageScoreRewardMetaData"]},
	"544413173":  EventDialogMultiLang("sl"), //Single Language (CN/ZH/JP/KO)
	"816421621":  EventDialogMultiLang("cn"),
	"816421643":  EventDialogMultiLang("de"),
	"816421683":  EventDialogMultiLang("en"),
	"816421718":  EventDialogMultiLang("fr"),
	"816421797":  EventDialogMultiLang("id"),
	"816422142":  EventDialogMultiLang("th"),
	"816422210":  EventDialogMultiLang("vn"),
	"834781912":  {dump.AssetName["QAvatarMissionMetaData"]},
	"857374862":  {dump.AssetName["MatrixSuddenEventMetaData"]},
	"872686180":  {dump.AssetName["StageRestrictMetaData"]},
	"877472804":  {dump.AssetName["VirtualAvatarMetaData"]},
	"897231668":  {dump.AssetName["GodWarCollectionDataMetaData"]},
	"897362620":  {dump.AssetName["OpenWorldEventData"]},
	"899544326":  {dump.AssetName["CompensationOfIslandMetaData"]},
	"952148573":  {dump.AssetName["OpenWorldSection"]},
	"956779179":  {dump.AssetName["MatrixGridIndexMetaData"]},
	"973890938":  {dump.AssetName["ShopGoodsTagMetaData"]},
	"979499011":  {dump.AssetName["ChapterOWTalentLevelMetaData"]},
	"983144477":  {dump.AssetName["OWEndlessInvadeMetaData"]},
	"1004459941": {dump.AssetName["DLCTowerScheduleMetaData"]},
	"1008604113": {dump.AssetName["MPTeamSkillMetaData"]},
	"1013202273": {dump.AssetName["EquipForgeMeta"]},
	"1035807515": {dump.AssetName["RpgSiteProgressConfigMetaData"]},
	"1044417225": {dump.AssetName["RogueBuffPoolMetaData"]},
	"1047012884": {dump.AssetName["ArmadaBossActivityScoreMetaData"]},
	"1071581160": {dump.AssetName["DiceyDungeonDailyScheduleMetaData"]},
	"1097457432": {dump.AssetName["PhonePendantMetaData"]},
	"1101385821": {dump.AssetName["HalfBalanceModeAttrMetaData"]},
	"1101456918": {dump.AssetName["HalfBalanceModeDataMetaData"]},
	"1104804610": {dump.AssetName["GodWarAvatarTaleScheduleMetaData"]},
	"1105592154": {dump.AssetName["ElfSkillTreeMetaData"]},
	"1152412547": {dump.AssetName["CityEventPhotoData"]},
	"1162560144": {dump.AssetName["GodWarEventStageMetaData"]},
	"1172429283": {dump.AssetName["MonopolyLevelInfoMetaData"]},
	"1185954507": {dump.AssetName["VirtualStigmataMetaData"]},
	"1194660516": {dump.AssetName["LinkDataMetaData"]},
	"1197875885": {dump.AssetName["DormitoryVoiceMetaData"]},
	"1200550125": {dump.AssetName["GodWarChallengeLevelMetaData"]},
	"1226355738": {dump.AssetName["RogueActiveCardMetaData"]},
	"1226677989": {dump.AssetName["ChapterOWHeroCardLevelMetaData"]},
	"1226967721": {dump.AssetName["CycleMissionManagementMetaData"]},
	"1227160674": {dump.AssetName["EventCollectionMetaData"]},
	"1237064793": {dump.AssetName["RpgSurvivalRoleMetaData"]},
	"1237093255": {dump.AssetName["DLCTalentLevelMetaData"]},
	"1248866144": {dump.AssetName["ThemeScheduleMetaData"]},
	"1326960348": {dump.AssetName["VirtualStigmataDetailMetaData"]},
	"1372996550": {dump.AssetName["WeaponMetaData"]},
	"1373747633": {dump.AssetName["DLCTalentMetaData"]},
	"1383713590": {dump.AssetName["MissionData"]},
	"1396039932": {dump.AssetName["GlobalExploreFlagMetaData"]},
	"1400794614": {dump.AssetName["RpgAreaMetaData"]},
	"1401298399": {dump.AssetName["RpgRoleMetaData"]},
	"1401322672": {dump.AssetName["RpgSiteMetaData"]},
	"1401344527": {dump.AssetName["RpgTaleMetaData"]},
	"1405469077": {dump.AssetName["TVTMessageMetaData"]},
	"1407107708": {dump.AssetName["AdventureQuestMetaData"]},
	"1407478510": {dump.AssetName["AdventureQuestPoolMetaData"]},
	"1413692358": {dump.AssetName["GlobalExploreActionMetaData"]},
	"1435715286": {dump.AssetName["DormitoryEventSequenceMetaData"]},
	"1438013711": {dump.AssetName["DanmakuMetaData"]},
	"1468594424": {dump.AssetName["OWEndlessMonsterData"]},
	"1538367859": {dump.AssetName["GlobalExploreEntityMetaData"]},
	"1541948906": {dump.AssetName["MonsterCardMetaData"]},
	"1546333431": {dump.AssetName["GodWarChallengeMetaData"]},
	"1554732703": {dump.AssetName["WaveRushBuffLevelGroupMetaData"]},
	"1556318929": {dump.AssetName["OptionalBuffListMetaData"]},
	"1566239312": {dump.AssetName["ConstValueMetaData"]},
	"1599871598": {dump.AssetName["StigmataRuneAffixMetaData"]},
	"1601481717": {dump.AssetName["InControlKeyInfoData"]},
	"1614253435": {dump.AssetName["RanchMonsterSkillMetaData"]},
	"1614612283": {dump.AssetName["ActChallengeMetaData"]},
	"1616735675": {dump.AssetName["NewbieActivityStageMissionMetaData"]},
	"1618516151": {dump.AssetName["ShopGoodsScheduleMetaData"]},
	"1640755418": {dump.AssetName["StigmataTagMetaData"]},
	"1646402813": {dump.AssetName["OpenWorldQuestJudgeData"]},
	"1654586486": {dump.AssetName["KingdomsStageMetaData"]},
	"1659355289": {dump.AssetName["HoukaiTownStrengthLevelMetaData"]},
	"1683941788": {dump.AssetName["TileMapInfoMetaData"]},
	"1686160113": {dump.AssetName["LinearMissionData"]},
	"1691684697": {dump.AssetName["RewardData"]},
	"1699439678": {dump.AssetName["DLCSupportDataMetaData"]},
	"1714228303": {dump.AssetName["LevelChallengeMetaData"]},
	"1724684358": {dump.AssetName["EditorUniqueMonsterMetaData"]},
	"1781437247": {dump.AssetName["GodWarBuffMetaData"]},
	"1781932595": {dump.AssetName["GodWarSiteMetaData"]},
	"1787634743": {dump.AssetName["StigmataVirtualSetMetaData"]},
	"1797372405": {dump.AssetName["MonsterProtoTypeMetaData"]},
	"1812163242": {dump.AssetName["MaterialUseMetaData"]},
	"1830025156": {dump.AssetName["CollectionFileMetaData"]},
	"1838763346": {dump.AssetName["QAvatarBattleTextIdMapMetaData"]},
	"1872287297": {dump.AssetName["OWEndlessMonsterGroupScore"]},
	"1889836997": {dump.AssetName["ChatLobbyAnnouncementMetaData"]},
	"1922042657": {dump.AssetName["MedalMetaData"]},
	"1923827730": {dump.AssetName["HoukaiTownQBattleSkillMetaData"]},
	"1926164450": {dump.AssetName["DLC2DailyQuestInfoMetaData"]},
	"2024081468": {dump.AssetName["MaterialDeleteData"]},
	"2040571594": {dump.AssetName["GodWarMissionDataMetaData"]},
	"2056165738": {dump.AssetName["WebLinkAvatarMetaData"]},
	"2059223482": {dump.AssetName["StageDetailAbilityMetaData"]},
	"2065585253": {dump.AssetName["MapSiteStageDataMetaData"]},
	"2072194043": {dump.AssetName["NPCLevelMetaData"]},
	"2079853359": {dump.AssetName["CurrencyIconMetaData"]},
	"2100673299": {dump.AssetName["AvatarCardMetaData"]},
	"2104488917": {dump.AssetName["GoBackMissionScoreMetaData"]},
	"2124307750": {dump.AssetName["CollectionMonsterMetaData"]},
	"2134817660": {dump.AssetName["AdventureWelfareMetaData"]},
	"2179819857": {dump.AssetName["ItemMetaData"]},
	"2185014648": {dump.AssetName["GeneralActivityStageGroupMetaData"]},
	"2240309350": {dump.AssetName["CollectionEventMetaData"]},
	"2245461054": {dump.AssetName["HoukaiTownBuildingMetaData"]},
	"2250293013": {dump.AssetName["MatrixFloorMetaData"]},
	"2255345609": {dump.AssetName["GodWarMainMissionNodeMetaData"]},
	"2275419147": {dump.AssetName["PlotMetaData"]},
	"2305568562": {dump.AssetName["NewbieGroupAdvData"]},
	"2323972294": {dump.AssetName["GlobalExploreGridMapMetaData"]},
	"2338384154": {dump.AssetName["ChapterOWAchievementMetaData"]},
	"2344131772": {dump.AssetName["LevelTrialMetaData"]},
	"2386659147": {dump.AssetName["ChapterOWTalentDataMetaData"]},
	"2388104161": {dump.AssetName["NewbieEquipAdvData"]},
	"2388638465": {dump.AssetName["ExaminationQuestionMetaData"]},
	"2389119108": {dump.AssetName["ExaminationAnswerMetaData"]},
	"2414780799": {dump.AssetName["PlotLineDataMetaData"]},
	"2432136154": {dump.AssetName["TeamBuffMetaData"]},
	"2470006924": {dump.AssetName["InLevelBuffUIMetaData"]},
	"2483258034": {dump.AssetName["GalEventMetaData"]},
	"2525212431": {dump.AssetName["UltraEndlessSiteMetaData"]},
	"2576386967": {dump.AssetName["OpenWorldStoryData"]},
	"2578607515": TextMapMultiLang("cn"),
	"2578607537": TextMapMultiLang("de"),
	"2578607577": TextMapMultiLang("en"),
	"2578607612": TextMapMultiLang("fr"),
	"2578607691": TextMapMultiLang("id"),
	"2578608036": TextMapMultiLang("th"),
	"2578608104": TextMapMultiLang("vn"),
	"2581767792": {dump.AssetName["StageDialogMetaData"]},
	"2590234221": {dump.AssetName["GodWarLobbyAvatarDataMetaData"]},
	"2638133815": {dump.AssetName["WeekDayActivityMetaData"]},
	"2644439028": {dump.AssetName["OpenWorldQuestRarityRewardMetaData"]},
	"2682852434": {dump.AssetName["DiceyDungeonPassiveSkillMetaData"]},
	"2696174958": {dump.AssetName["ScDLCTowerFloorMetaData"]},
	"2704317315": {dump.AssetName["BattlePassTypeMetaData"]},
	"2710207575": {dump.AssetName["DormitoryDialogMetaData"]},
	"2734189206": {dump.AssetName["RanchMonsterMetaData"]},
	"2743855919": {dump.AssetName["RichBuffMetaData"]},
	"2753470167": {dump.AssetName["OpenworldQuestBuffData"]},
	"2770137763": {dump.AssetName["ChatLobbyNPCMetaData"]},
	"2772757206": {dump.AssetName["RestaurantAvatarMetaData"]},
	"2774311853": {dump.AssetName["EquipSkillMetaData"]},
	"2777342908": {dump.AssetName["ElfSkillMetaData"]},
	"2782073203": {dump.AssetName["AddParamGroupMetaData"]},
	"2782171632": {dump.AssetName["OWEndlessInvadeBuffMetaData"]},
	"2796833292": {dump.AssetName["ChapterOWEquipmentBuffMetaData"]},
	"2812396528": {dump.AssetName["LevelMetaData"]},
	"2812422669": {dump.AssetName["TileEntityDataMetaData"]},
	"2829361031": {dump.AssetName["InLevelDialog"]},
	"2840629778": {dump.AssetName["DiceyDungeonMonsterMetaData"]},
	"2846487885": {dump.AssetName["AvatarPracticeMainMetaData"]},
	"2846684768": {dump.AssetName["AvatarPracticeStepMetaData"]},
	"2873405398": {dump.AssetName["HybridSiteContentMetaData"]},
	"2877142287": {dump.AssetName["TextMapMetaData"]}, //Single Lang (CN/ZH/JP/KO)
	"2889932192": {dump.AssetName["UnknownAsset2MetaData"]},
	"2907879079": {dump.AssetName["LevelMatrixGridIndexMetaData"]},
	"2910390514": {dump.AssetName["DialogMetaData"]},
	"2911346328": {dump.AssetName["OWEndlessMonsterInitTypeMetaData"]},
	"2917070182": {dump.AssetName["RpgSuddenEventMetaData"]},
	"2920424082": {dump.AssetName["GodWarRelationDataMetaData"]},
	"2928991298": {dump.AssetName["AvatarSkillMetaData"]},
	"2932484622": {dump.AssetName["CgMetaData"]},
	"2934471913": {dump.AssetName["CoupleTowerScoreMetaData"]},
	"2967440797": {dump.AssetName["DormitoryEventDialogMetaData"]},
	"2971210687": {dump.AssetName["RpgUIConfigMetaData"]},
	"2974904445": {dump.AssetName["ChapterOWMemorySiteMetaData"]},
	"2979402384": {dump.AssetName["SLGStageMetaData"]},
	"2980129774": {dump.AssetName["TileMapMetaData"]},
	"2986114947": {dump.AssetName["RpgMissionMetaData"]},
	"2988272078": {dump.AssetName["InControlActionMapData"]},
	"3003489637": {dump.AssetName["MPStageMetaData"]},
	"3080227030": {dump.AssetName["GodWarLobbyInteractActionMetaData"]},
	"3092346878": {dump.AssetName["WeaponTagMetaData"]},
	"3096985441": {dump.AssetName["EliteAbilityMetaData"]},
	"3127645492": {dump.AssetName["MonsterWikiDataMetaData"]},
	"3133001532": {dump.AssetName["ActMetaData"]},
	"3134841825": {dump.AssetName["JigsawMetaData"]},
	"3135063653": {dump.AssetName["OpenWorldExplorePointMetaData"]},
	"3138930165": {dump.AssetName["ScDLCTowerScheduleMetaData"]},
	"3152235461": {dump.AssetName["RestaurantQuestMetaData"]},
	"3153788212": {dump.AssetName["RestaurantSkillMetaData"]},
	"3158472400": {dump.AssetName["BurdenAlleviationMetaData"]},
	"3179169915": {dump.AssetName["MissionPanelConversationMetaData"]},
	"3184361587": {dump.AssetName["QTEndlessMonsterData"]},
	"3184927682": {dump.AssetName["QTEndlessMonsterWaveMetaData"]},
	"3201233750": {dump.AssetName["PVZQavatarMetaData"]},
	"3205845523": {dump.AssetName["UniqueMonsterMetaData"]},
	"3219510525": {dump.AssetName["RpgStageScoreMetaData"]},
	"3222135066": {dump.AssetName["AvatarModelMetaData"]},
	"3242467349": {dump.AssetName["OpenWorldArea"]},
	"3263951298": {dump.AssetName["UltraEndlessBattleConfigMetaData"]},
	"3269946612": {dump.AssetName["RecommendEquipMetaData"]},
	"3273959509": {dump.AssetName["ActivityPanelMetaData"]},
	"3279621832": {dump.AssetName["DailyRecommendMeta"]},
	"3284085781": {dump.AssetName["DormitoryDecorationMetaData"]},
	"3338147799": {dump.AssetName["RpgClientMainDataMetaData"]},
	"3347524550": {dump.AssetName["WaveRushStageWaveMetaData"]},
	"3347643230": {dump.AssetName["PicTutorialStepDataMetaData"]},
	"3355042839": {dump.AssetName["GodWarTeleportMetaData"]},
	"3358461928": {dump.AssetName["ArmadaBossActivityDataMetaData"]},
	"3359281148": {dump.AssetName["TileValueKeyMetaData"]},
	"3364732199": {dump.AssetName["MechMetaData"]},
	"3371129601": {dump.AssetName["StageDetailEffectMetaData"]},
	"3373320453": {dump.AssetName["LoadingTipStringMetaData"]},
	"3378048260": {dump.AssetName["OpenworldQuestDataCfg"]},
	"3378821031": {dump.AssetName["LevelStatisticsMetaData"]},
	"3392770755": {dump.AssetName["ClientLogDataMetaData"]},
	"3407709807": {dump.AssetName["ActivityLoginMetaData"]},
	"3424769365": {dump.AssetName["AvatarStarMetaData"]},
	"3426901140": {dump.AssetName["GodWarPunishBuffMetaData"]},
	"3427406605": {dump.AssetName["GodWarPunishStepMetaData"]},
	"3431409261": {dump.AssetName["AvatarSampleMetaData"]},
	"3435206935": {dump.AssetName["TowerStageConfigMetaData"]},
	"3456902459": {dump.AssetName["TileModelMetaData"]},
	"3467656643": {dump.AssetName["ScratchTicketItemDataMetaData"]},
	"3506541559": {dump.AssetName["GeneralActivityMetaData"]},
	"3507933901": {dump.AssetName["DormitoryEventWeightMetaData"]},
	"3509442358": {dump.AssetName["RpgCollectionRewardData"]},
	"3518340181": {dump.AssetName["GenActivityRewardShowItemsMetaData"]},
	"3519580761": {dump.AssetName["AvatarForgeWhiteListMetaData"]},
	"3525001332": {dump.AssetName["ChatLobbyVoiceMetaData"]},
	"3535691949": {dump.AssetName["ProductRecommendMetaData"]},
	"3566516475": {dump.AssetName["SinDemonExRankRewardMetaData"]},
	"3579822406": {dump.AssetName["MonsterConfigMetaData"]},
	"3600254561": {dump.AssetName["DressMetaData"]},
	"3605151818": {dump.AssetName["CollectionPictureMetaData"]},
	"3618301546": {dump.AssetName["StigmataMetaData"]},
	"3618877226": {dump.AssetName["StageDropItemDataMetaData"]},
	"3635635046": {dump.AssetName["RogueGoodsMetaData"]},
	"3647356129": {dump.AssetName["RpgNavSiteMetaData"]},
	"3653023269": {dump.AssetName["AvatarTagUnLockMetaData"]},
	"3659505636": {dump.AssetName["SinDemonExMonsterMetaData"]},
	"3679377011": {dump.AssetName["GiftPackMetaData"]},
	"3687746862": {dump.AssetName["GodWarEventMetaData"]},
	"3701749502": {dump.AssetName["RpgEventTextMetaData"]},
	"3712542493": {dump.AssetName["MonsterWikiDescConfigMetaData"]},
	"3721846672": {dump.AssetName["RichBuildingMetaData"]},
	"3750327922": {dump.AssetName["VirtualCustomLevelDataMetaData"]},
	"3750842160": {dump.AssetName["GodWarPunishRewardMetaData"]},
	"3817790113": {dump.AssetName["SeriesMetaData"]},
	"3881117200": {dump.AssetName["RpgMaterialMetaData"]},
	"3891426031": {dump.AssetName["MPRaidActMetaData"]},
	"3892272259": {dump.AssetName["AvatarMetaData"]},
	"3895069218": {dump.AssetName["ChapterActivityStageMetaData"]},
	"3909035470": {dump.AssetName["WeekDayActivityScheduleMetaData"]},
	"3924533703": {dump.AssetName["HoukaiTownMapMetaData"]},
	"3932072382": {dump.AssetName["DLCTowerFloorMetaData"]},
	"3944639586": {dump.AssetName["RpgBuffDataMetaData"]},
	"3954218449": {dump.AssetName["SinDemonExMonsterScheduleMetaData"]},
	"3963850408": {dump.AssetName["MatrixThemeMetaData"]},
	"3996597855": {dump.AssetName["UltraEndlessScheduleMetaData"]},
	"4011635046": {dump.AssetName["GoBackScheduleMetaData"]},
	"4077943566": {dump.AssetName["RogueStageMetaData"]},
	"4087158562": {dump.AssetName["ChapterOWMainLineMetaData"]},
	"4150724661": {dump.AssetName["RanchSiteDataMetaData"]},
	"4182714193": {dump.AssetName["ChapterOWSpecialStoryMetaData"]},
	"4200546953": {dump.AssetName["DiceyDungeonSkillMetaData"]},
	"4204779011": {dump.AssetName["GodWarLobbyInteractPropMetaData"]},
	"4206115175": {dump.AssetName["RandomPlotNPCMetaData"]},
	"4215361678": {dump.AssetName["NetworkErrCodeMetaData"]},
	"4218920404": {dump.AssetName["TutorialStepData"]},
	"4232884778": {dump.AssetName["ActivityBingoMetaData"]},
	"4240244261": {dump.AssetName["ThemeAvatarMetaData"]},
	"4242391060": {dump.AssetName["RandomEffectMetaData"]},
	"4243688563": {dump.AssetName["MissionGroupMetaData"]},
	"4272897265": {dump.AssetName["TowerGrowBuffConfigMetaData"]},
	"4280865965": {dump.AssetName["ScDLCTalentLevelMetaData"]},
	"4281796563": {dump.AssetName["HoukaiTownQMonsterMetaData"]},
	"4285153761": {dump.AssetName["MailDataMetaData"]},
	"4288747552": {dump.AssetName["InControlUIButtonConfig"]},
	"4291465330": {dump.AssetName["DiceyDungeonDataMetaData"]},
	"4291895614": {dump.AssetName["DiceyDungeonRoleMetaData"]},
	"4291895715": {dump.AssetName["DiceyDungeonRoomMetaData"]},

	// multiple matches
	"1715048037": {dump.AssetName["ChapterOWMoonConditionMetaData"], dump.AssetName["DLC2ConditionMetaData"]},
	"2853541098": {dump.AssetName["AiCyberMainUIRepairMetaData"], dump.AssetName["ChapterStageCompensationMetaData"], dump.AssetName["InviteeActivityGobackConfigMetaData"], dump.AssetName["InviteeActivityNewbieConfigMetaData"], dump.AssetName["OWHuntActivityMachineEventMetaData"], dump.AssetName["ReclaimRankingRewardData"], dump.AssetName["ReclaimScoreRewardData"], dump.AssetName["ScDLCMPStageDisplayMetaData"], dump.AssetName["GeneralActivityTicketMetaDataReader"]},
	"3393012976": {dump.AssetName["DLC2ConditionMetaData"], dump.AssetName["ChapterOWMoonConditionMetaData"]},
}
