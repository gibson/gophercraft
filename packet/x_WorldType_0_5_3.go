package packet

var worldTypeDescriptor_Alpha = &WorldTypeDescriptor{
	Codes: WorldCodes{
		CMSG_BOOTME:                                0x0001,
		CMSG_DBLOOKUP:                              0x0002,
		SMSG_DBLOOKUP:                              0x0003,
		CMSG_QUERY_OBJECT_POSITION:                 0x0004,
		SMSG_QUERY_OBJECT_POSITION:                 0x0005,
		CMSG_QUERY_OBJECT_ROTATION:                 0x0006,
		SMSG_QUERY_OBJECT_ROTATION:                 0x0007,
		CMSG_WORLD_TELEPORT:                        0x0008,
		CMSG_TELEPORT_TO_PLAYER:                    0x0009,
		CMSG_ZONE_MAP:                              0x000A,
		SMSG_ZONE_MAP:                              0x000B,
		CMSG_DEBUG_CHANGECELLZONE:                  0x000C,
		CMSG_EMBLAZON_TABARD_OBSOLETE:              0x000D,
		CMSG_UNEMBLAZON_TABARD_OBSOLETE:            0x000E,
		CMSG_RECHARGE:                              0x000F,
		CMSG_LEARN_SPELL:                           0x0010,
		CMSG_CREATEMONSTER:                         0x0011,
		CMSG_DESTROYMONSTER:                        0x0012,
		CMSG_CREATEITEM:                            0x0013,
		CMSG_CREATEGAMEOBJECT:                      0x0014,
		CMSG_MAKEMONSTERATTACKME:                   0x0015,
		CMSG_MAKEMONSTERATTACKGUID:                 0x0016,
		CMSG_ENABLEDEBUGCOMBATLOGGING:              0x0017,
		CMSG_FORCEACTION:                           0x0018,
		CMSG_FORCEACTIONONOTHER:                    0x0019,
		CMSG_FORCEACTIONSHOW:                       0x001A,
		SMSG_FORCEACTIONSHOW:                       0x001B,
		SMSG_ATTACKERSTATEUPDATEDEBUGINFO:          0x001C,
		SMSG_ATTACKERSTATEUPDATEDEBUGINFOSPELL:     0x001D,
		SMSG_ATTACKERSTATEUPDATEDEBUGINFOSPELLMISS: 0x001E,
		SMSG_DEBUG_PLAYER_RANGE:                    0x001F,
		CMSG_UNDRESSPLAYER:                         0x0020,
		CMSG_BEASTMASTER:                           0x0021,
		CMSG_GODMODE:                               0x0022,
		SMSG_GODMODE:                               0x0023,
		CMSG_CHEAT_SETMONEY:                        0x0024,
		CMSG_LEVEL_CHEAT:                           0x0025,
		CMSG_PET_LEVEL_CHEAT:                       0x0026,
		CMSG_LEVELUP_CHEAT:                         0x0027,
		CMSG_COOLDOWN_CHEAT:                        0x0028,
		CMSG_USE_SKILL_CHEAT:                       0x0029,
		CMSG_FLAG_QUEST:                            0x002A,
		CMSG_FLAG_QUEST_FINISH:                     0x002B,
		CMSG_CLEAR_QUEST:                           0x002C,
		CMSG_SEND_EVENT:                            0x002D,
		CMSG_DEBUG_AISTATE:                         0x002E,
		SMSG_DEBUG_AISTATE:                         0x002F,
		CMSG_ENABLE_PVP:                            0x0030,
		CMSG_ADVANCE_SPAWN_TIME:                    0x0031,
		CMSG_PVP_PORT:                              0x0032,
		CMSG_AUTH_SRP6_BEGIN:                       0x0033,
		CMSG_AUTH_SRP6_PROOF:                       0x0034,
		CMSG_AUTH_SRP6_RECODE:                      0x0035,
		CMSG_CHAR_CREATE:                           0x0036,
		CMSG_CHAR_ENUM:                             0x0037,
		CMSG_CHAR_DELETE:                           0x0038,
		SMSG_AUTH_SRP6_RESPONSE:                    0x0039,
		SMSG_CHAR_CREATE:                           0x003A,
		SMSG_CHAR_ENUM:                             0x003B,
		SMSG_CHAR_DELETE:                           0x003C,
		CMSG_PLAYER_LOGIN:                          0x003D,
		SMSG_NEW_WORLD:                             0x003E,
		SMSG_TRANSFER_PENDING:                      0x003F,
		SMSG_TRANSFER_ABORTED:                      0x0040,
		SMSG_CHARACTER_LOGIN_FAILED:                0x0041,
		SMSG_LOGIN_SETTIMESPEED:                    0x0042,
		SMSG_GAMETIME_UPDATE:                       0x0043,
		CMSG_GAMETIME_SET:                          0x0044,
		SMSG_GAMETIME_SET:                          0x0045,
		CMSG_GAMESPEED_SET:                         0x0046,
		SMSG_GAMESPEED_SET:                         0x0047,
		CMSG_SERVERTIME:                            0x0048,
		SMSG_SERVERTIME:                            0x0049,
		CMSG_PLAYER_LOGOUT:                         0x004A,
		CMSG_LOGOUT_REQUEST:                        0x004B,
		SMSG_LOGOUT_RESPONSE:                       0x004C,
		SMSG_LOGOUT_COMPLETE:                       0x004D,
		CMSG_LOGOUT_CANCEL:                         0x004E,
		SMSG_LOGOUT_CANCEL_ACK:                     0x004F,
		CMSG_NAME_QUERY:                            0x0050,
		SMSG_NAME_QUERY_RESPONSE:                   0x0051,
		CMSG_PET_NAME_QUERY:                        0x0052,
		SMSG_PET_NAME_QUERY_RESPONSE:               0x0053,
		CMSG_GUILD_QUERY:                           0x0054,
		SMSG_GUILD_QUERY_RESPONSE:                  0x0055,
		CMSG_ITEM_QUERY_SINGLE:                     0x0056,
		CMSG_ITEM_QUERY_MULTIPLE:                   0x0057,
		SMSG_ITEM_QUERY_SINGLE_RESPONSE:            0x0058,
		SMSG_ITEM_QUERY_MULTIPLE_RESPONSE:          0x0059,
		CMSG_PAGE_TEXT_QUERY:                       0x005A,
		SMSG_PAGE_TEXT_QUERY_RESPONSE:              0x005B,
		CMSG_QUEST_QUERY:                           0x005C,
		SMSG_QUEST_QUERY_RESPONSE:                  0x005D,
		CMSG_GAMEOBJECT_QUERY:                      0x005E,
		SMSG_GAMEOBJECT_QUERY_RESPONSE:             0x005F,
		CMSG_CREATURE_QUERY:                        0x0060,
		SMSG_CREATURE_QUERY_RESPONSE:               0x0061,
		CMSG_WHO:                                   0x0062,
		SMSG_WHO:                                   0x0063,
		CMSG_WHOIS:                                 0x0064,
		SMSG_WHOIS:                                 0x0065,
		CMSG_FRIEND_LIST:                           0x0066,
		SMSG_FRIEND_LIST:                           0x0067,
		SMSG_FRIEND_STATUS:                         0x0068,
		CMSG_ADD_FRIEND:                            0x0069,
		CMSG_DEL_FRIEND:                            0x006A,
		SMSG_IGNORE_LIST:                           0x006B,
		CMSG_ADD_IGNORE:                            0x006C,
		CMSG_DEL_IGNORE:                            0x006D,
		CMSG_GROUP_INVITE:                          0x006E,
		SMSG_GROUP_INVITE:                          0x006F,
		CMSG_GROUP_CANCEL:                          0x0070,
		SMSG_GROUP_CANCEL:                          0x0071,
		CMSG_GROUP_ACCEPT:                          0x0072,
		CMSG_GROUP_DECLINE:                         0x0073,
		SMSG_GROUP_DECLINE:                         0x0074,
		CMSG_GROUP_UNINVITE:                        0x0075,
		CMSG_GROUP_UNINVITE_GUID:                   0x0076,
		SMSG_GROUP_UNINVITE:                        0x0077,
		CMSG_GROUP_SET_LEADER:                      0x0078,
		SMSG_GROUP_SET_LEADER:                      0x0079,
		CMSG_LOOT_METHOD:                           0x007A,
		CMSG_GROUP_DISBAND:                         0x007B,
		SMSG_GROUP_DESTROYED:                       0x007C,
		SMSG_GROUP_LIST:                            0x007D,
		SMSG_PARTY_MEMBER_STATS:                    0x007E,
		SMSG_PARTY_COMMAND_RESULT:                  0x007F,
		UMSG_UPDATE_GROUP_MEMBERS:                  0x0080,
		CMSG_GUILD_CREATE:                          0x0081,
		CMSG_GUILD_INVITE:                          0x0082,
		SMSG_GUILD_INVITE:                          0x0083,
		CMSG_GUILD_ACCEPT:                          0x0084,
		CMSG_GUILD_DECLINE:                         0x0085,
		SMSG_GUILD_DECLINE:                         0x0086,
		CMSG_GUILD_INFO:                            0x0087,
		SMSG_GUILD_INFO:                            0x0088,
		CMSG_GUILD_ROSTER:                          0x0089,
		SMSG_GUILD_ROSTER:                          0x008A,
		CMSG_GUILD_PROMOTE:                         0x008B,
		CMSG_GUILD_DEMOTE:                          0x008C,
		CMSG_GUILD_LEAVE:                           0x008D,
		CMSG_GUILD_REMOVE:                          0x008E,
		CMSG_GUILD_DISBAND:                         0x008F,
		CMSG_GUILD_LEADER:                          0x0090,
		CMSG_GUILD_MOTD:                            0x0091,
		SMSG_GUILD_EVENT:                           0x0092,
		SMSG_GUILD_COMMAND_RESULT:                  0x0093,
		UMSG_UPDATE_GUILD:                          0x0094,
		CMSG_MESSAGECHAT:                           0x0095,
		SMSG_MESSAGECHAT:                           0x0096,
		CMSG_JOIN_CHANNEL:                          0x0097,
		CMSG_LEAVE_CHANNEL:                         0x0098,
		SMSG_CHANNEL_NOTIFY:                        0x0099,
		CMSG_CHANNEL_LIST:                          0x009A,
		SMSG_CHANNEL_LIST:                          0x009B,
		CMSG_CHANNEL_PASSWORD:                      0x009C,
		CMSG_CHANNEL_SET_OWNER:                     0x009D,
		CMSG_CHANNEL_OWNER:                         0x009E,
		CMSG_CHANNEL_MODERATOR:                     0x009F,
		CMSG_CHANNEL_UNMODERATOR:                   0x00A0,
		CMSG_CHANNEL_MUTE:                          0x00A1,
		CMSG_CHANNEL_UNMUTE:                        0x00A2,
		CMSG_CHANNEL_INVITE:                        0x00A3,
		CMSG_CHANNEL_KICK:                          0x00A4,
		CMSG_CHANNEL_BAN:                           0x00A5,
		CMSG_CHANNEL_UNBAN:                         0x00A6,
		CMSG_CHANNEL_ANNOUNCEMENTS:                 0x00A7,
		CMSG_CHANNEL_MODERATE:                      0x00A8,
		SMSG_UPDATE_OBJECT:                         0x00A9,
		SMSG_DESTROY_OBJECT:                        0x00AA,
		CMSG_USE_ITEM:                              0x00AB,
		CMSG_OPEN_ITEM:                             0x00AC,
		CMSG_READ_ITEM:                             0x00AD,
		SMSG_READ_ITEM_OK:                          0x00AE,
		SMSG_READ_ITEM_FAILED:                      0x00AF,
		SMSG_ITEM_COOLDOWN:                         0x00B0,
		CMSG_GAMEOBJ_USE:                           0x00B1,
		CMSG_GAMEOBJ_CHAIR_USE_OBSOLETE:            0x00B2,
		SMSG_GAMEOBJECT_CUSTOM_ANIM:                0x00B3,
		CMSG_AREATRIGGER:                           0x00B4,
		MSG_MOVE_START_FORWARD:                     0x00B5,
		MSG_MOVE_START_BACKWARD:                    0x00B6,
		MSG_MOVE_STOP:                              0x00B7,
		MSG_MOVE_START_STRAFE_LEFT:                 0x00B8,
		MSG_MOVE_START_STRAFE_RIGHT:                0x00B9,
		MSG_MOVE_STOP_STRAFE:                       0x00BA,
		MSG_MOVE_JUMP:                              0x00BB,
		MSG_MOVE_START_TURN_LEFT:                   0x00BC,
		MSG_MOVE_START_TURN_RIGHT:                  0x00BD,
		MSG_MOVE_STOP_TURN:                         0x00BE,
		MSG_MOVE_START_PITCH_UP:                    0x00BF,
		MSG_MOVE_START_PITCH_DOWN:                  0x00C0,
		MSG_MOVE_STOP_PITCH:                        0x00C1,
		MSG_MOVE_SET_RUN_MODE:                      0x00C2,
		MSG_MOVE_SET_WALK_MODE:                     0x00C3,
		MSG_MOVE_TOGGLE_LOGGING:                    0x00C4,
		MSG_MOVE_TELEPORT:                          0x00C5,
		MSG_MOVE_TELEPORT_CHEAT:                    0x00C6,
		SMSG_MOVE_WORLDPORT_ACK:                    0x00C7,
		MSG_MOVE_TOGGLE_FALL_LOGGING:               0x00C8,
		MSG_MOVE_COLLIDE_REDIRECT:                  0x00C9,
		MSG_MOVE_COLLIDE_STUCK:                     0x00CA,
		MSG_MOVE_START_SWIM:                        0x00CB,
		MSG_MOVE_STOP_SWIM:                         0x00CC,
		MSG_MOVE_SET_RUN_SPEED_CHEAT:               0x00CD,
		MSG_MOVE_SET_RUN_SPEED:                     0x00CE,
		MSG_MOVE_SET_WALK_SPEED_CHEAT:              0x00CF,
		MSG_MOVE_SET_WALK_SPEED:                    0x00D0,
		MSG_MOVE_SET_SWIM_SPEED_CHEAT:              0x00D1,
		MSG_MOVE_SET_SWIM_SPEED:                    0x00D2,
		MSG_MOVE_SET_ALL_SPEED_CHEAT:               0x00D3,
		MSG_MOVE_SET_TURN_RATE_CHEAT:               0x00D4,
		MSG_MOVE_SET_TURN_RATE:                     0x00D5,
		MSG_MOVE_TOGGLE_COLLISION_CHEAT:            0x00D6,
		MSG_MOVE_SET_FACING:                        0x00D7,
		MSG_MOVE_SET_PITCH:                         0x00D8,
		MSG_MOVE_WORLDPORT_ACK:                     0x00D9,
		SMSG_MONSTER_MOVE:                          0x00DA,
		MSG_MOVE_RESERVED_0:                        0x00DB,
		MSG_MOVE_RESERVED_1:                        0x00DC,
		MSG_MOVE_RESERVED_2:                        0x00DD,
		MSG_MOVE_RESERVED_3:                        0x00DE,
		SMSG_FORCE_SPEED_CHANGE:                    0x00DF,
		CMSG_FORCE_SPEED_CHANGE_ACK:                0x00E0,
		SMSG_FORCE_SWIM_SPEED_CHANGE:               0x00E1,
		CMSG_FORCE_SWIM_SPEED_CHANGE_ACK:           0x00E2,
		SMSG_FORCE_MOVE_ROOT:                       0x00E3,
		CMSG_FORCE_MOVE_ROOT_ACK:                   0x00E4,
		SMSG_FORCE_MOVE_UNROOT:                     0x00E5,
		CMSG_FORCE_MOVE_UNROOT_ACK:                 0x00E6,
		MSG_MOVE_ROOT:                              0x00E7,
		MSG_MOVE_UNROOT:                            0x00E8,
		MSG_MOVE_HEARTBEAT:                         0x00E9,
		CMSG_STUCK_OBSOLETE:                        0x00EA,
		CMSG_TRIGGER_CINEMATIC_CHEAT:               0x00EB,
		CMSG_OPENING_CINEMATIC:                     0x00EC,
		SMSG_TRIGGER_CINEMATIC:                     0x00ED,
		CMSG_NEXT_CINEMATIC_CAMERA:                 0x00EE,
		CMSG_COMPLETE_CINEMATIC:                    0x00EF,
		SMSG_TUTORIAL_FLAGS:                        0x00F0,
		CMSG_TUTORIAL_SHOWN:                        0x00F1,
		CMSG_TUTORIAL_CLEAR:                        0x00F2,
		CMSG_TUTORIAL_RESET:                        0x00F3,
		CMSG_STANDSTATECHANGE:                      0x00F4,
		CMSG_EMOTE:                                 0x00F5,
		SMSG_EMOTE:                                 0x00F6,
		CMSG_TEXT_EMOTE:                            0x00F7,
		SMSG_TEXT_EMOTE:                            0x00F8,
		CMSG_AUTOEQUIP_GROUND_ITEM:                 0x00F9,
		CMSG_AUTOSTORE_GROUND_ITEM:                 0x00FA,
		CMSG_AUTOSTORE_LOOT_ITEM:                   0x00FB,
		CMSG_STORE_LOOT_IN_SLOT:                    0x00FC,
		CMSG_AUTOEQUIP_ITEM:                        0x00FD,
		CMSG_AUTOSTORE_BAG_ITEM:                    0x00FE,
		CMSG_SWAP_ITEM:                             0x00FF,
		CMSG_SWAP_INV_ITEM:                         0x0100,
		CMSG_SPLIT_ITEM:                            0x0101,
		CMSG_PICKUP_ITEM:                           0x0102,
		CMSG_DROP_ITEM:                             0x0103,
		CMSG_DESTROYITEM:                           0x0104,
		SMSG_INVENTORY_CHANGE_FAILURE:              0x0105,
		SMSG_OPEN_CONTAINER:                        0x0106,
		CMSG_INSPECT:                               0x0107,
		SMSG_INSPECT:                               0x0108,
		CMSG_INITIATE_TRADE:                        0x0109,
		CMSG_BEGIN_TRADE:                           0x010A,
		CMSG_BUSY_TRADE:                            0x010B,
		CMSG_IGNORE_TRADE:                          0x010C,
		CMSG_ACCEPT_TRADE:                          0x010D,
		CMSG_UNACCEPT_TRADE:                        0x010E,
		CMSG_CANCEL_TRADE:                          0x010F,
		CMSG_SET_TRADE_ITEM:                        0x0110,
		CMSG_CLEAR_TRADE_ITEM:                      0x0111,
		CMSG_SET_TRADE_GOLD:                        0x0112,
		SMSG_TRADE_STATUS:                          0x0113,
		SMSG_TRADE_STATUS_EXTENDED:                 0x0114,
		SMSG_INITIALIZE_FACTIONS:                   0x0115,
		SMSG_SET_FACTION_VISIBLE:                   0x0116,
		SMSG_SET_FACTION_STANDING:                  0x0117,
		CMSG_SET_FACTION_ATWAR:                     0x0118,
		CMSG_SET_FACTION_CHEAT:                     0x0119,
		SMSG_SET_PROFICIENCY:                       0x011A,
		CMSG_SET_ACTION_BUTTON:                     0x011B,
		SMSG_ACTION_BUTTONS:                        0x011C,
		SMSG_INITIAL_SPELLS:                        0x011D,
		SMSG_LEARNED_SPELL:                         0x011E,
		SMSG_SUPERCEDED_SPELL:                      0x011F,
		CMSG_NEW_SPELL_SLOT:                        0x0120,
		CMSG_CAST_SPELL:                            0x0121,
		CMSG_CANCEL_CAST:                           0x0122,
		SMSG_CAST_RESULT:                           0x0123,
		SMSG_SPELL_START:                           0x0124,
		SMSG_SPELL_GO:                              0x0125,
		SMSG_SPELL_FAILURE:                         0x0126,
		SMSG_SPELL_COOLDOWN:                        0x0127,
		SMSG_COOLDOWN_EVENT:                        0x0128,
		CMSG_CANCEL_AURA:                           0x0129,
		SMSG_UPDATE_AURA_DURATION:                  0x012A,
		SMSG_PET_CAST_FAILED:                       0x012B,
		MSG_CHANNEL_START:                          0x012C,
		MSG_CHANNEL_UPDATE:                         0x012D,
		CMSG_CANCEL_CHANNELLING:                    0x012E,
		SMSG_AI_REACTION:                           0x012F,
		CMSG_SET_SELECTION:                         0x0130,
		CMSG_SET_TARGET:                            0x0131,
		CMSG_START_USING_RANGED_WEAPON:             0x0132,
		CMSG_STOP_USING_RANGED_WEAPON:              0x0133,
		CMSG_ATTACKSWING:                           0x0134,
		CMSG_ATTACKSTOP:                            0x0135,
		SMSG_ATTACKSTART:                           0x0136,
		SMSG_ATTACKSTOP:                            0x0137,
		SMSG_ATTACKSWING_NOTINRANGE:                0x0138,
		SMSG_ATTACKSWING_BADFACING:                 0x0139,
		SMSG_ATTACKSWING_NOTSTANDING:               0x013A,
		SMSG_ATTACKSWING_DEADTARGET:                0x013B,
		SMSG_ATTACKSWING_CANT_ATTACK:               0x013C,
		SMSG_ATTACKERSTATEUPDATE:                   0x013D,
		SMSG_VICTIMSTATEUPDATE_OBSOLETE:            0x013E,
		SMSG_DAMAGE_DONE:                           0x013F,
		SMSG_DAMAGE_TAKEN:                          0x0140,
		SMSG_CANCEL_COMBAT:                         0x0141,
		SMSG_PLAYER_COMBAT_XP_GAIN_OBSOLETE:        0x0142,
		SMSG_HEALSPELL_ON_PLAYER:                   0x0143,
		SMSG_HEALSPELL_ON_PLAYERS_PET:              0x0144,
		CMSG_SETSHEATHED:                           0x0145,
		CMSG_SAVE_PLAYER:                           0x0146,
		CMSG_SETDEATHBINDPOINT:                     0x0147,
		SMSG_BINDPOINTUPDATE:                       0x0148,
		CMSG_GETDEATHBINDZONE:                      0x0149,
		SMSG_BINDZONEREPLY:                         0x014A,
		SMSG_PLAYERBOUND:                           0x014B,
		SMSG_DEATH_NOTIFY:                          0x014C,
		CMSG_REPOP_REQUEST:                         0x014D,
		SMSG_RESURRECT_REQUEST:                     0x014E,
		CMSG_RESURRECT_RESPONSE:                    0x014F,
		CMSG_LOOT:                                  0x0150,
		CMSG_LOOT_MONEY:                            0x0151,
		CMSG_LOOT_RELEASE:                          0x0152,
		SMSG_LOOT_RESPONSE:                         0x0153,
		SMSG_LOOT_RELEASE_RESPONSE:                 0x0154,
		SMSG_LOOT_REMOVED:                          0x0155,
		SMSG_LOOT_MONEY_NOTIFY:                     0x0156,
		SMSG_LOOT_ITEM_NOTIFY:                      0x0157,
		SMSG_LOOT_CLEAR_MONEY:                      0x0158,
		SMSG_ITEM_PUSH_RESULT:                      0x0159,
		SMSG_DUEL_REQUESTED:                        0x015A,
		SMSG_DUEL_OUTOFBOUNDS:                      0x015B,
		SMSG_DUEL_INBOUNDS:                         0x015C,
		SMSG_DUEL_COMPLETE:                         0x015D,
		SMSG_DUEL_WINNER:                           0x015E,
		CMSG_DUEL_ACCEPTED:                         0x015F,
		CMSG_DUEL_CANCELLED:                        0x0160,
		SMSG_MOUNTRESULT:                           0x0161,
		SMSG_DISMOUNTRESULT:                        0x0162,
		SMSG_PUREMOUNT_CANCELLED:                   0x0163,
		CMSG_MOUNTSPECIAL_ANIM:                     0x0164,
		SMSG_MOUNTSPECIAL_ANIM:                     0x0165,
		SMSG_PET_TAME_FAILURE:                      0x0166,
		CMSG_PET_SET_ACTION:                        0x0167,
		CMSG_PET_ACTION:                            0x0168,
		CMSG_PET_ABANDON:                           0x0169,
		CMSG_PET_RENAME:                            0x016A,
		SMSG_PET_NAME_INVALID:                      0x016B,
		SMSG_PET_SPELLS:                            0x016C,
		CMSG_PET_CAST_SPELL_OBSOLETE:               0x016D,
		CMSG_LIST_INVENTORY:                        0x016E,
		SMSG_LIST_INVENTORY:                        0x016F,
		CMSG_SELL_ITEM:                             0x0170,
		SMSG_SELL_ITEM:                             0x0171,
		CMSG_BUY_ITEM:                              0x0172,
		CMSG_BUY_ITEM_IN_SLOT:                      0x0173,
		SMSG_BUY_ITEM:                              0x0174,
		SMSG_BUY_FAILED:                            0x0175,
		CMSG_GOSSIP_HELLO:                          0x0176,
		SMSG_GOSSIP_MESSAGE:                        0x0177,
		CMSG_NPC_TEXT_QUERY:                        0x0178,
		SMSG_NPC_TEXT_UPDATE:                       0x0179,
		CMSG_NPC_OFFER_ITEM:                        0x017A,
		SMSG_NPC_ACCEPT_ITEM:                       0x017B,
		SMSG_NPC_DECLINE_ITEM:                      0x017C,
		SMSG_NPC_WONT_TALK:                         0x017D,
		CMSG_QUESTGIVER_STATUS_QUERY:               0x017E,
		SMSG_QUESTGIVER_STATUS:                     0x017F,
		CMSG_QUESTGIVER_HELLO:                      0x0180,
		SMSG_QUESTGIVER_QUEST_LIST:                 0x0181,
		CMSG_QUESTGIVER_QUERY_QUEST:                0x0182,
		CMSG_QUESTGIVER_QUEST_AUTOLAUNCH:           0x0183,
		SMSG_QUESTGIVER_QUEST_DETAILS:              0x0184,
		CMSG_QUESTGIVER_ACCEPT_QUEST:               0x0185,
		CMSG_QUESTGIVER_COMPLETE_QUEST:             0x0186,
		SMSG_QUESTGIVER_REQUEST_ITEMS:              0x0187,
		CMSG_QUESTGIVER_REQUEST_REWARD:             0x0188,
		SMSG_QUESTGIVER_OFFER_REWARD:               0x0189,
		CMSG_QUESTGIVER_CHOOSE_REWARD:              0x018A,
		SMSG_QUESTGIVER_QUEST_INVALID:              0x018B,
		CMSG_QUESTGIVER_CANCEL:                     0x018C,
		SMSG_QUESTGIVER_QUEST_COMPLETE:             0x018D,
		SMSG_QUESTGIVER_QUEST_FAILED:               0x018E,
		CMSG_QUESTLOG_SWAP_QUEST:                   0x018F,
		CMSG_QUESTLOG_REMOVE_QUEST:                 0x0190,
		SMSG_QUESTLOG_FULL:                         0x0191,
		SMSG_QUESTUPDATE_FAILED:                    0x0192,
		SMSG_QUESTUPDATE_COMPLETE:                  0x0193,
		SMSG_QUESTUPDATE_ADD_KILL:                  0x0194,
		SMSG_QUESTUPDATE_ADD_ITEM:                  0x0195,
		CMSG_QUEST_CONFIRM_ACCEPT:                  0x0196,
		SMSG_QUEST_CONFIRM_ACCEPT:                  0x0197,
		CMSG_TAXICLEARALLNODES:                     0x0198,
		CMSG_TAXIENABLEALLNODES:                    0x0199,
		CMSG_TAXISHOWNODES:                         0x019A,
		SMSG_SHOWTAXINODES:                         0x019B,
		CMSG_TAXINODE_STATUS_QUERY:                 0x019C,
		SMSG_TAXINODE_STATUS:                       0x019D,
		CMSG_TAXIQUERYAVAILABLENODES:               0x019E,
		CMSG_ACTIVATETAXI:                          0x019F,
		SMSG_ACTIVATETAXIREPLY:                     0x01A0,
		SMSG_NEW_TAXI_PATH:                         0x01A1,
		CMSG_TRAINER_LIST:                          0x01A2,
		SMSG_TRAINER_LIST:                          0x01A3,
		CMSG_TRAINER_BUY_SPELL:                     0x01A4,
		SMSG_TRAINER_BUY_SUCCEEDED:                 0x01A5,
		SMSG_TRAINER_BUY_FAILED:                    0x01A6,
		CMSG_BINDER_ACTIVATE:                       0x01A7,
		SMSG_PLAYERBINDERROR:                       0x01A8,
		CMSG_BANKER_ACTIVATE:                       0x01A9,
		SMSG_SHOW_BANK:                             0x01AA,
		CMSG_BUY_BANK_SLOT:                         0x01AB,
		SMSG_BUY_BANK_SLOT_RESULT:                  0x01AC,
		CMSG_PETITION_SHOWLIST:                     0x01AD,
		SMSG_PETITION_SHOWLIST:                     0x01AE,
		CMSG_PETITION_BUY:                          0x01AF,
		CMSG_PETITION_SHOW_SIGNATURES:              0x01B0,
		SMSG_PETITION_SHOW_SIGNATURES:              0x01B1,
		CMSG_PETITION_SIGN:                         0x01B2,
		SMSG_PETITION_SIGN_RESULTS:                 0x01B3,
		CMSG_OFFER_PETITION:                        0x01B4,
		CMSG_TURN_IN_PETITION:                      0x01B5,
		SMSG_TURN_IN_PETITION_RESULTS:              0x01B6,
		CMSG_PETITION_QUERY:                        0x01B7,
		SMSG_PETITION_QUERY_RESPONSE:               0x01B8,
		SMSG_FISH_NOT_HOOKED:                       0x01B9,
		SMSG_FISH_ESCAPED:                          0x01BA,
		CMSG_BUG:                                   0x01BB,
		SMSG_NOTIFICATION:                          0x01BC,
		CMSG_PLAYED_TIME:                           0x01BD,
		SMSG_PLAYED_TIME:                           0x01BE,
		CMSG_QUERY_TIME:                            0x01BF,
		SMSG_QUERY_TIME_RESPONSE:                   0x01C0,
		SMSG_LOG_XPGAIN:                            0x01C1,
		MSG_SPLIT_MONEY:                            0x01C2,
		CMSG_RECLAIM_CORPSE:                        0x01C3,
		CMSG_WRAP_ITEM:                             0x01C4,
		SMSG_LEVELUP_INFO:                          0x01C5,
		MSG_MINIMAP_PING:                           0x01C6,
		SMSG_RESISTLOG:                             0x01C7,
		SMSG_ENCHANTMENTLOG:                        0x01C8,
		CMSG_SET_SKILL_CHEAT:                       0x01C9,
		SMSG_START_MIRROR_TIMER:                    0x01CA,
		SMSG_PAUSE_MIRROR_TIMER:                    0x01CB,
		SMSG_STOP_MIRROR_TIMER:                     0x01CC,
		CMSG_PING:                                  0x01CD,
		SMSG_PONG:                                  0x01CE,
		SMSG_CLEAR_COOLDOWN:                        0x01CF,
		SMSG_GAMEOBJECT_PAGETEXT:                   0x01D0,
		CMSG_SETWEAPONMODE:                         0x01D1,
		SMSG_COOLDOWN_CHEAT:                        0x01D2,
		SMSG_SPELL_DELAYED:                         0x01D3,
		CMSG_PLAYER_MACRO:                          0x01D4,
		SMSG_PLAYER_MACRO:                          0x01D5,
		CMSG_GHOST:                                 0x01D6,
		CMSG_GM_INVIS:                              0x01D7,
		CMSG_SCREENSHOT:                            0x01D8,
		MSG_GM_BIND_OTHER:                          0x01D9,
		MSG_GM_SUMMON:                              0x01DA,
		SMSG_ITEM_TIME_UPDATE:                      0x01DB,
		SMSG_ITEM_ENCHANT_TIME_UPDATE:              0x01DC,
		SMSG_AUTH_CHALLENGE:                        0x01DD,
		CMSG_AUTH_SESSION:                          0x01DE,
		SMSG_AUTH_RESPONSE:                         0x01DF,
		MSG_GM_SHOWLABEL:                           0x01E0,
		MSG_ADD_DYNAMIC_TARGET:                     0x01E1,
		MSG_SAVE_GUILD_EMBLEM:                      0x01E2,
		MSG_TABARDVENDOR_ACTIVATE:                  0x01E3,
		SMSG_PLAY_SPELL_VISUAL:                     0x01E4,
		CMSG_ZONEUPDATE:                            0x01E5,
		SMSG_PARTYKILLLOG:                          0x01E6,
		SMSG_COMPRESSED_UPDATE_OBJECT:              0x01E7,
		SMSG_MIRRORTIMERDAMAGELOG:                  0x01E8,
		SMSG_EXPLORATION_EXPERIENCE:                0x01E9,
		CMSG_GM_SET_SECURITY_GROUP:                 0x01EA,
		CMSG_GM_NUKE:                               0x01EB,
		MSG_RANDOM_ROLL:                            0x01EC,
		SMSG_ENVIRONMENTALDAMAGELOG:                0x01ED,
		CMSG_RWHOIS:                                0x01EE,
		SMSG_RWHOIS:                                0x01EF,
		MSG_LOOKING_FOR_GROUP:                      0x01F0,
		CMSG_SET_LOOKING_FOR_GROUP:                 0x01F1,
	},
}
