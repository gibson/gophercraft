package dbd

import (
	"strings"
	"testing"
	"time"

	"github.com/Gophercraft/log"
)

const ChrRaces = `COLUMNS
int ID
int Flags
int<FactionTemplate::ID> FactionID
int<SoundEntries::ID> ExplorationSoundID
int<CreatureDisplayInfo::ID> MaleDisplayID
int<CreatureDisplayInfo::ID> FemaleDisplayID
string ClientPrefix
int<Languages::ID> BaseLanguage
int<CreatureType::ID> CreatureType
int<Spell::ID> ResSicknessSpellID
int<SoundEntries::ID> SplashSoundID
string ClientFileString
int<CinematicSequences::ID> CinematicSequenceID
int Alliance
locstring Name_lang // quest text $R
locstring Name_female_lang // quest text $R
locstring Name_male_lang
string FacialHairCustomization
string HairCustomization
int Race_related
int<ChrRaces::ID> UnalteredVisualRaceID
int<CreatureSoundData::ID> UaMaleCreatureSoundDataID
int<CreatureSoundData::ID> UaFemaleCreatureSoundDataID
int<CharComponentTextureLayouts::ID> CharComponentTextureLayoutID
int<ChrClasses::ID> DefaultClassID
int<FileData::ID> CreateScreenFileDataID
int<FileData::ID> SelectScreenFileDataID
float MaleCustomizeOffset
float FemaleCustomizeOffset
int<ChrRaces::ID> NeutralRaceID
int<FileData::ID> LowResScreenFileDataID
int<CreatureDisplayInfo::ID> HighResMaleDisplayID
int<CreatureDisplayInfo::ID> HighResFemaleDisplayID
int<CharComponentTextureLayouts::ID> CharComponentTexLayoutHiResID
int AlteredFormFinishVisualKitID
int AlteredFormStartVisualKitID
int DisplayRaceID
int<FileData::ID> FemaleSkeletonFileDataID
int HeritageArmorAchievementID
int<FileData::ID> MaleSkeletonFileDataID
locstring Name_female_lowercase_lang // quest text $r
locstring Name_lowercase_lang // quest text $r
int StartingLevel
int UiDisplayOrder
int Required_expansion
float MountScale
int<Spell::ID> LoginEffectSpellID
int<Spell::ID> CombatStunSpellID
int StartingTaxiNodes
float FemaleModelFallbackArmor2Scale
int FemaleModelFallbackRaceID
int FemaleModelFallbackSex
int FemaleTextureFallbackRaceID
int FemaleTextureFallbackSex
float MaleModelFallbackArmor2Scale
int MaleModelFallbackRaceID
int MaleModelFallbackSex
int MaleTextureFallbackRaceID
int MaleTextureFallbackSex
int<ChrRaces::ID> BaseRaceID?
int TransmogrifyDisabledSlotMask?
locstring Name_RS_lang? // quest text $RS
locstring Name_RS_female_lang? // quest text $RS
locstring Name_RS_lowercase_lang? // quest text $rs
locstring Name_RS_female_lowercase_lang? // quest text $rs
locstring Name_RL_lang? // quest text $RL
locstring Name_RL_female_lang? // quest text $RL
locstring Name_RL_lowercase_lang? // quest text $rl
locstring Name_RL_female_lowercase_lang? // quest text $rl
int PlayableRaceBit?
locstring RaceFantasyDescription_lang?
int<ChrRaces::ID> UnalteredVisualCustomizationRaceID?
float AlteredFormCustomizeOffsetFallback? // fallback for CharStartKit::AlteredFormCustomizeOffset
float AlteredFormCustomizeRotationFallback? // fallback for CharStartKit::AlteredFormCustomizeRotation
float Field_9_0_1_35256_033? // probably: fallback for CharStartKit::Field_9_0_1_35256_002
float Field_9_0_1_35256_034? // probably: fallback for CharStartKit::Field_9_0_1_35256_003
int<ChrRaces::ID> HelmetAnimScalingRaceID?
float Field_9_1_0_38312_030?
float Field_9_1_0_38312_031?
int Field_9_1_0_38783_022?
int Field_9_1_0_38783_026?
int Field_9_1_0_38783_033?
int Field_9_1_0_38783_034?
int Field_9_1_0_38783_035?
int Field_9_1_0_38783_036?
int Field_9_1_0_38783_038?
int Field_9_1_0_38783_040?
int Field_9_1_0_38783_047?
int Field_9_1_0_38783_048?
int Field_9_1_0_38783_049?
int Field_9_1_0_38783_050?

BUILD 0.5.3.3368
BUILD 0.5.3.3368-0.6.0.3592
$id$ID<32>
Flags<32>
FactionID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
MountScale
BaseLanguage<32>
CreatureType<32>
LoginEffectSpellID<32>
CombatStunSpellID<32>
ResSicknessSpellID<32>
SplashSoundID<32>
StartingTaxiNodes<32>
ClientFileString
CinematicSequenceID<32>
Name_lang

BUILD 0.7.0.3694-0.9.1.3810
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
MountScale
BaseLanguage<32>
CreatureType<32>
LoginEffectSpellID<32>
CombatStunSpellID<32>
ResSicknessSpellID<32>
SplashSoundID<32>
StartingTaxiNodes<32>
ClientFileString
CinematicSequenceID<32>
Name_lang

BUILD 2.0.0.5610-2.0.0.5666
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
MountScale
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Name_lang
FacialHairCustomization[2]
HairCustomization

BUILD 2.0.0.5991
BUILD 1.0.0.3980-1.12.2.6005
BUILD 0.10.0.3892-0.12.0.3988
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
MountScale
BaseLanguage<32>
CreatureType<32>
LoginEffectSpellID<32>
CombatStunSpellID<32>
ResSicknessSpellID<32>
SplashSoundID<32>
StartingTaxiNodes<32>
ClientFileString
CinematicSequenceID<32>
Name_lang
FacialHairCustomization[2]
HairCustomization

BUILD 2.0.0.6080-2.3.3.7799
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
MountScale
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Name_lang
FacialHairCustomization[2]
HairCustomization
Required_expansion<32>

LAYOUT 8BE33A75
BUILD 2.5.1.38043, 2.5.1.38061, 2.5.1.38093, 2.5.1.38118, 2.5.1.38119, 2.5.1.38169, 2.5.1.38225, 2.5.1.38285, 2.5.1.38339, 2.5.1.38364, 2.5.1.38401, 2.5.1.38453, 2.5.1.38502, 2.5.1.38521, 2.5.1.38537, 2.5.1.38548, 2.5.1.38561, 2.5.1.38598, 2.5.1.38644, 2.5.1.38677, 2.5.1.38692, 2.5.1.38707, 2.5.1.38739, 2.5.1.38741, 2.5.1.38757, 2.5.1.38835, 2.5.1.38892, 2.5.1.38921, 2.5.1.38988, 2.5.1.39170
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
$id$ID<32>
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
ResSicknessSpellID<32>
SplashSoundID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
FemaleSkeletonFileDataID<32>
MaleSkeletonFileDataID<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
AlteredFormCustomizeOffsetFallback[3]
AlteredFormCustomizeRotationFallback
Field_9_0_1_35256_033[3]
Field_9_0_1_35256_034[3]
FactionID<16>
CinematicSequenceID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
UnalteredVisualCustomizationRaceID<8>

BUILD 3.0.1.8303-3.1.3.9947
BUILD 2.4.0.8089-2.4.3.8606
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
MountScale
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Required_expansion<32>

BUILD 3.2.0.10192-3.2.0.10314
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Required_expansion<32>

BUILD 3.2.2.10482-3.3.5.12340
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Alliance<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Required_expansion<32>

BUILD 4.0.0.11792
BUILD 4.0.0.11927-4.0.0.12319
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Alliance<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Required_expansion<32>
UnalteredVisualRaceID<32>

BUILD 4.0.0.12911-4.3.4.15595
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Alliance<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Race_related<32>
UnalteredVisualRaceID<32>
UaMaleCreatureSoundDataID<32>
UaFemaleCreatureSoundDataID<32>

BUILD 5.0.1.15464-5.0.1.15662
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Alliance<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Race_related<32>
UnalteredVisualRaceID<32>
UaMaleCreatureSoundDataID<32>
UaFemaleCreatureSoundDataID<32>
CharComponentTextureLayoutID<32>

BUILD 5.0.1.15668-5.0.1.15699
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Alliance<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Race_related<32>
UnalteredVisualRaceID<32>
UaMaleCreatureSoundDataID<32>
UaFemaleCreatureSoundDataID<32>
CharComponentTextureLayoutID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
NeutralRaceID<32>

BUILD 5.0.1.15726-5.0.1.15799
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Alliance<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Race_related<32>
UnalteredVisualRaceID<32>
UaMaleCreatureSoundDataID<32>
UaFemaleCreatureSoundDataID<32>
CharComponentTextureLayoutID<32>
DefaultClassID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
NeutralRaceID<32>

BUILD 5.0.2.15827-5.4.8.18414
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Alliance<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Race_related<32>
UnalteredVisualRaceID<32>
UaMaleCreatureSoundDataID<32>
UaFemaleCreatureSoundDataID<32>
CharComponentTextureLayoutID<32>
DefaultClassID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
NeutralRaceID<32>
LowResScreenFileDataID<32>

BUILD 6.0.1.18179
$id$ID<32>
Flags<32>
FactionID<32>
ExplorationSoundID<32>
MaleDisplayID<32>
FemaleDisplayID<32>
ClientPrefix
BaseLanguage<32>
CreatureType<32>
ResSicknessSpellID<32>
SplashSoundID<32>
ClientFileString
CinematicSequenceID<32>
Alliance<32>
Name_lang
Name_female_lang
Name_male_lang
FacialHairCustomization[2]
HairCustomization
Race_related<32>
UnalteredVisualRaceID<32>
UaMaleCreatureSoundDataID<32>
UaFemaleCreatureSoundDataID<32>
CharComponentTextureLayoutID<32>
DefaultClassID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
NeutralRaceID<32>
LowResScreenFileDataID<32>
HighResMaleDisplayID<32>
HighResFemaleDisplayID<32>
CharComponentTexLayoutHiResID<32>

LAYOUT 608397F0
BUILD 7.3.2.25079, 7.3.2.25163, 7.3.2.25196, 7.3.2.25208, 7.3.2.25255, 7.3.2.25326, 7.3.2.25383, 7.3.2.25442, 7.3.2.25455, 7.3.2.25477, 7.3.2.25480, 7.3.2.25497, 7.3.2.25516, 7.3.2.25549
BUILD 7.3.0.24473, 7.3.0.24484, 7.3.0.24492, 7.3.0.24500, 7.3.0.24539, 7.3.0.24563, 7.3.0.24608, 7.3.0.24651, 7.3.0.24681, 7.3.0.24692, 7.3.0.24700, 7.3.0.24715, 7.3.0.24727, 7.3.0.24730, 7.3.0.24738, 7.3.0.24744, 7.3.0.24758, 7.3.0.24759, 7.3.0.24781, 7.3.0.24793, 7.3.0.24829, 7.3.0.24834, 7.3.0.24843, 7.3.0.24845, 7.3.0.24852, 7.3.0.24853, 7.3.0.24864, 7.3.0.24878, 7.3.0.24887, 7.3.0.24896, 7.3.0.24904, 7.3.0.24920, 7.3.0.24931, 7.3.0.24956, 7.3.0.24970, 7.3.0.24974, 7.3.0.25021, 7.3.0.25195
$noninline,id$ID<32>
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
FacialHairCustomization[2]
HairCustomization
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
FactionID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
CinematicSequenceID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
DefaultClassID<8>
NeutralRaceID<8>
DisplayRaceID<8>
CharComponentTexLayoutHiResID<8>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]

LAYOUT 51C511F9
BUILD 8.0.1.25902, 8.0.1.25976, 8.0.1.26010, 8.0.1.26032, 8.0.1.26095, 8.0.1.26131, 8.0.1.26141, 8.0.1.26175
BUILD 7.3.5.25928, 7.3.5.25937, 7.3.5.25944, 7.3.5.25946, 7.3.5.25950, 7.3.5.25961, 7.3.5.25996, 7.3.5.26124, 7.3.5.26365, 7.3.5.26654, 7.3.5.26755, 7.3.5.26822, 7.3.5.26899, 7.3.5.26972
BUILD 1.13.0.28211
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
StartingLevel<32>
UiDisplayOrder<32>
FactionID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
CinematicSequenceID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
DefaultClassID<8>
NeutralRaceID<8>
DisplayRaceID<8>
CharComponentTexLayoutHiResID<8>
$id$ID<32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
HeritageArmorAchievementID<32>
MaleSkeletonFileDataID<32>
FemaleSkeletonFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]

LAYOUT 7E728145
BUILD 8.0.1.26231
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
StartingLevel<32>
UiDisplayOrder<32>
MaleModelFallbackArmor2Scale
FemaleModelFallbackArmor2Scale
FactionID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
CinematicSequenceID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
DefaultClassID<8>
NeutralRaceID<8>
CharComponentTexLayoutHiResID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
$id$ID<32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
HeritageArmorAchievementID<32>
MaleSkeletonFileDataID<32>
FemaleSkeletonFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]

LAYOUT 7AC8831B
BUILD 8.0.1.26287, 8.0.1.26297, 8.0.1.26310, 8.0.1.26321, 8.0.1.26367
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
StartingLevel<32>
UiDisplayOrder<32>
MaleModelFallbackArmor2Scale
FemaleModelFallbackArmor2Scale
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
CharComponentTexLayoutHiResID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
$id$ID<32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
HeritageArmorAchievementID<32>
FemaleSkeletonFileDataID<32>
MaleSkeletonFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]

LAYOUT BCC93784
BUILD 8.0.1.26433
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
StartingLevel<32>
UiDisplayOrder<32>
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
CharComponentTexLayoutHiResID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
$id$ID<32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
HeritageArmorAchievementID<32>
FemaleSkeletonFileDataID<32>
MaleSkeletonFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]

LAYOUT E7078C91
BUILD 8.0.1.26476, 8.0.1.26491, 8.0.1.26522, 8.0.1.26530, 8.0.1.26557, 8.0.1.26567, 8.0.1.26585, 8.0.1.26604, 8.0.1.26610
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
StartingLevel<32>
UiDisplayOrder<32>
BaseRaceID<32>
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
CharComponentTexLayoutHiResID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
$id$ID<32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
HeritageArmorAchievementID<32>
FemaleSkeletonFileDataID<32>
MaleSkeletonFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]

LAYOUT 2B2A4FD3
BUILD 8.0.1.26624, 8.0.1.26629, 8.0.1.26637, 8.0.1.26640, 8.0.1.26683, 8.0.1.26707, 8.0.1.26714, 8.0.1.26715, 8.0.1.26734, 8.0.1.26766
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
UiDisplayOrder<32>
BaseRaceID<32>
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
CharComponentTexLayoutHiResID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
$id$ID<32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
HeritageArmorAchievementID<32>
StartingLevel<32>
FemaleSkeletonFileDataID<32>
MaleSkeletonFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]

LAYOUT C8BCDC89, C0CADF15
BUILD 8.3.0.32044, 8.3.0.32151, 8.3.0.32203, 8.3.0.32218, 8.3.0.32272, 8.3.0.32414, 8.3.0.32489
BUILD 8.2.5.31337, 8.2.5.31401, 8.2.5.31521, 8.2.5.31599, 8.2.5.31812, 8.2.5.31884, 8.2.5.31919, 8.2.5.31921, 8.2.5.31958, 8.2.5.31960, 8.2.5.31961, 8.2.5.31984, 8.2.5.32028, 8.2.5.32079, 8.2.5.32144, 8.2.5.32185, 8.2.5.32265, 8.2.5.32294, 8.2.5.32305, 8.2.5.32494, 8.2.5.32580, 8.2.5.32638, 8.2.5.32722, 8.2.5.32750, 8.2.5.32978
BUILD 8.2.0.30080, 8.2.0.30093, 8.2.0.30096, 8.2.0.30108, 8.2.0.30123, 8.2.0.30168, 8.2.0.30170, 8.2.0.30203, 8.2.0.30262, 8.2.0.30329, 8.2.0.30430, 8.2.0.30495, 8.2.0.30613, 8.2.0.30634, 8.2.0.30669, 8.2.0.30774, 8.2.0.30827, 8.2.0.30888, 8.2.0.30898, 8.2.0.30918, 8.2.0.30920, 8.2.0.30948, 8.2.0.30993, 8.2.0.31229, 8.2.0.31429, 8.2.0.31478
BUILD 8.1.5.28938, 8.1.5.29141, 8.1.5.29220, 8.1.5.29281, 8.1.5.29310, 8.1.5.29352, 8.1.5.29385, 8.1.5.29418, 8.1.5.29484, 8.1.5.29495, 8.1.5.29558, 8.1.5.29599, 8.1.5.29620, 8.1.5.29634, 8.1.5.29664, 8.1.5.29683, 8.1.5.29701, 8.1.5.29704, 8.1.5.29705, 8.1.5.29718, 8.1.5.29732, 8.1.5.29737, 8.1.5.29814, 8.1.5.29869, 8.1.5.29896, 8.1.5.29981, 8.1.5.30477, 8.1.5.30706
BUILD 8.1.0.27826, 8.1.0.27934, 8.1.0.27985, 8.1.0.28048, 8.1.0.28085, 8.1.0.28151, 8.1.0.28186, 8.1.0.28202, 8.1.0.28294, 8.1.0.28366, 8.1.0.28440, 8.1.0.28485, 8.1.0.28616, 8.1.0.28657, 8.1.0.28710, 8.1.0.28724, 8.1.0.28768, 8.1.0.28807, 8.1.0.28822, 8.1.0.28833, 8.1.0.29088, 8.1.0.29139, 8.1.0.29235, 8.1.0.29285, 8.1.0.29297, 8.1.0.29482, 8.1.0.29600, 8.1.0.29621
BUILD 8.0.1.26788, 8.0.1.26806, 8.0.1.26812, 8.0.1.26838, 8.0.1.26871, 8.0.1.26877, 8.0.1.26892, 8.0.1.26903, 8.0.1.26918, 8.0.1.26926, 8.0.1.26936, 8.0.1.26949, 8.0.1.26970, 8.0.1.26999, 8.0.1.27004, 8.0.1.27009, 8.0.1.27019, 8.0.1.27026, 8.0.1.27075, 8.0.1.27089, 8.0.1.27101, 8.0.1.27138, 8.0.1.27144, 8.0.1.27165, 8.0.1.27178, 8.0.1.27219, 8.0.1.27291, 8.0.1.27326, 8.0.1.27353, 8.0.1.27355, 8.0.1.27356, 8.0.1.27366, 8.0.1.27377, 8.0.1.27404, 8.0.1.27481, 8.0.1.27547, 8.0.1.27602, 8.0.1.27791, 8.0.1.27843, 8.0.1.27980, 8.0.1.28153
BUILD 1.13.7.37279, 1.13.7.37429, 1.13.7.37892, 1.13.7.38178, 1.13.7.38296, 1.13.7.38363, 1.13.7.38386, 1.13.7.38475, 1.13.7.38631, 1.13.7.38704
BUILD 1.13.6.36231, 1.13.6.36310, 1.13.6.36324, 1.13.6.36497, 1.13.6.36524, 1.13.6.36611, 1.13.6.36670, 1.13.6.36714, 1.13.6.36935, 1.13.6.37497
BUILD 1.13.5.34000, 1.13.5.34097, 1.13.5.34138, 1.13.5.34713, 1.13.5.34911, 1.13.5.35000, 1.13.5.35186, 1.13.5.35395, 1.13.5.35663, 1.13.5.35705, 1.13.5.35753, 1.13.5.36035, 1.13.5.36307, 1.13.5.36325
BUILD 1.13.4.33491, 1.13.4.33598, 1.13.4.33645, 1.13.4.33728, 1.13.4.33920, 1.13.4.34219, 1.13.4.34266, 1.13.4.34600, 1.13.4.34835
BUILD 1.13.3.32790, 1.13.3.32836, 1.13.3.32887, 1.13.3.33155, 1.13.3.33302, 1.13.3.33485, 1.13.3.33526
BUILD 1.13.2.30073, 1.13.2.30112, 1.13.2.30113, 1.13.2.30128, 1.13.2.30172, 1.13.2.30232, 1.13.2.30265, 1.13.2.30287, 1.13.2.30406, 1.13.2.30550, 1.13.2.30682, 1.13.2.30786, 1.13.2.30862, 1.13.2.30901, 1.13.2.30979, 1.13.2.31043, 1.13.2.31118, 1.13.2.31209, 1.13.2.31402, 1.13.2.31407, 1.13.2.31446, 1.13.2.31650, 1.13.2.31687, 1.13.2.31727, 1.13.2.31830, 1.13.2.31882, 1.13.2.32089, 1.13.2.32421, 1.13.2.32600
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
$id$ID<32>
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
FemaleSkeletonFileDataID<32>
MaleSkeletonFileDataID<32>
BaseRaceID<32>
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
CharComponentTexLayoutHiResID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>

LAYOUT C1BC2966
BUILD 8.3.7.34872, 8.3.7.35114, 8.3.7.35249, 8.3.7.35284, 8.3.7.35435, 8.3.7.35662
BUILD 8.3.0.32593, 8.3.0.32712, 8.3.0.32805, 8.3.0.32846, 8.3.0.32861, 8.3.0.32976, 8.3.0.33010, 8.3.0.33051, 8.3.0.33062, 8.3.0.33073, 8.3.0.33084, 8.3.0.33095, 8.3.0.33115, 8.3.0.33169, 8.3.0.33183, 8.3.0.33237, 8.3.0.33369, 8.3.0.33528, 8.3.0.33724, 8.3.0.33775, 8.3.0.33941, 8.3.0.34220, 8.3.0.34601, 8.3.0.34769, 8.3.0.34963
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
$id$ID<32>
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
FemaleSkeletonFileDataID<32>
MaleSkeletonFileDataID<32>
BaseRaceID<32>
TransmogrifyDisabledSlotMask<32>
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
CharComponentTexLayoutHiResID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>

LAYOUT 76D456C2
BUILD 9.0.1.33978, 9.0.1.34003
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
$id$ID<32>
Flags<32>
MaleDisplayID<u32>
FemaleDisplayID<u32>
HighResMaleDisplayID<u32>
HighResFemaleDisplayID<u32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
MaleCustomizeOffset[3]
FemaleCustomizeOffset[3]
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
FemaleSkeletonFileDataID<32>
MaleSkeletonFileDataID<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
CharComponentTextureLayoutID<8>
CharComponentTexLayoutHiResID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>

LAYOUT 6692653E
BUILD 9.0.1.34081, 9.0.1.34098, 9.0.1.34137, 9.0.1.34199, 9.0.1.34278, 9.0.1.34324, 9.0.1.34365, 9.0.1.34392, 9.0.1.34490, 9.0.1.34615
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
$id$ID<32>
Flags<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>

LAYOUT 561D0FB5
BUILD 9.0.1.34714, 9.0.1.34821, 9.0.1.34902, 9.0.1.34972
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
$id$ID<32>
Flags<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
UnalteredVisualCustomizationRaceID<8>

LAYOUT 04D6D8BD
BUILD 9.0.1.35078
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
$id$ID<32>
Flags<32>
BaseLanguage<32>
ResSicknessSpellID<32>
SplashSoundID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
FactionID<16>
CinematicSequenceID<16>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
UnalteredVisualCustomizationRaceID<8>

LAYOUT A21EDF78
BUILD 9.0.1.35256, 9.0.1.35282, 9.0.1.35360, 9.0.1.35432, 9.0.1.35482
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
$id$ID<32>
Flags<32>
BaseLanguage<32>
ResSicknessSpellID<32>
SplashSoundID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
AlteredFormCustomizeOffsetFallback[3]
AlteredFormCustomizeRotationFallback
Field_9_0_1_35256_033[3]
Field_9_0_1_35256_034[3]
FactionID<16>
CinematicSequenceID<16>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
UnalteredVisualCustomizationRaceID<8>

LAYOUT 85A50CA9
BUILD 9.0.5.37503, 9.0.5.37623, 9.0.5.37705, 9.0.5.37774, 9.0.5.37844, 9.0.5.37862, 9.0.5.37864, 9.0.5.37893, 9.0.5.37899, 9.0.5.37988, 9.0.5.38134, 9.0.5.38556
BUILD 9.0.2.35854, 9.0.2.35938, 9.0.2.35978, 9.0.2.36037, 9.0.2.36086, 9.0.2.36165, 9.0.2.36206, 9.0.2.36267, 9.0.2.36294, 9.0.2.36401, 9.0.2.36512, 9.0.2.36532, 9.0.2.36599, 9.0.2.36639, 9.0.2.36665, 9.0.2.36671, 9.0.2.36710, 9.0.2.36734, 9.0.2.36751, 9.0.2.36753, 9.0.2.36839, 9.0.2.36949, 9.0.2.37106, 9.0.2.37130, 9.0.2.37142, 9.0.2.37176, 9.0.2.37415, 9.0.2.37474
BUILD 9.0.1.35167, 9.0.1.35213, 9.0.1.35522, 9.0.1.35598, 9.0.1.35679, 9.0.1.35707, 9.0.1.35755, 9.0.1.35789, 9.0.1.35828, 9.0.1.35853, 9.0.1.35897, 9.0.1.35917, 9.0.1.35944, 9.0.1.35989, 9.0.1.36021, 9.0.1.36036, 9.0.1.36074, 9.0.1.36131, 9.0.1.36163, 9.0.1.36200, 9.0.1.36216, 9.0.1.36228, 9.0.1.36230, 9.0.1.36247, 9.0.1.36272, 9.0.1.36286, 9.0.1.36322, 9.0.1.36372, 9.0.1.36492, 9.0.1.36577
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
$id$ID<32>
Flags<32>
BaseLanguage<32>
ResSicknessSpellID<32>
SplashSoundID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
AlteredFormCustomizeOffsetFallback[3]
AlteredFormCustomizeRotationFallback
FactionID<16>
CinematicSequenceID<16>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
UnalteredVisualCustomizationRaceID<8>

LAYOUT C0605CB2
BUILD 9.1.0.38312, 9.1.0.38394, 9.1.0.38511, 9.1.0.38524
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
$id$ID<32>
Flags<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
AlteredFormCustomizeOffsetFallback[3]
AlteredFormCustomizeRotationFallback
Field_9_1_0_38312_030[3]
Field_9_1_0_38312_031[3]
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
UnalteredVisualCustomizationRaceID<8>

LAYOUT 7AFDB53D
BUILD 9.1.0.38549, 9.1.0.38600, 9.1.0.38627, 9.1.0.38709
$noninline,id$ID<32>
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
Flags<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<u32>[3]
AlteredFormFinishVisualKitID<u32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
PlayableRaceBit<32>
HelmetAnimScalingRaceID<32>
TransmogrifyDisabledSlotMask<32>
AlteredFormCustomizeOffsetFallback[3]
AlteredFormCustomizeRotationFallback
Field_9_1_0_38312_030[3]
Field_9_1_0_38312_031[3]
FactionID<16>
CinematicSequenceID<16>
ResSicknessSpellID<16>
SplashSoundID<16>
BaseLanguage<8>
CreatureType<8>
Alliance<8>
Race_related<8>
UnalteredVisualRaceID<8>
DefaultClassID<8>
NeutralRaceID<8>
MaleModelFallbackRaceID<8>
MaleModelFallbackSex<8>
FemaleModelFallbackRaceID<8>
FemaleModelFallbackSex<8>
MaleTextureFallbackRaceID<8>
MaleTextureFallbackSex<8>
FemaleTextureFallbackRaceID<8>
FemaleTextureFallbackSex<8>
UnalteredVisualCustomizationRaceID<8>

LAYOUT 0F5197A5
BUILD 9.1.0.38783, 9.1.0.38802, 9.1.0.38872, 9.1.0.38950, 9.1.0.39015, 9.1.0.39069, 9.1.0.39121, 9.1.0.39136, 9.1.0.39172
$noninline,id$ID<32>
ClientPrefix
ClientFileString
Name_lang
Name_female_lang
Name_lowercase_lang
Name_female_lowercase_lang
Name_RS_lang
Name_RS_female_lang
Name_RS_lowercase_lang
Name_RS_female_lowercase_lang
RaceFantasyDescription_lang
Name_RL_lang
Name_RL_female_lang
Name_RL_lowercase_lang
Name_RL_female_lowercase_lang
Flags<32>
FactionID<32>
CinematicSequenceID<32>
ResSicknessSpellID<32>
SplashSoundID<32>
Alliance<32>
Race_related<32>
Field_9_1_0_38783_022<32>
DefaultClassID<32>
CreateScreenFileDataID<32>
SelectScreenFileDataID<32>
Field_9_1_0_38783_026<32>
LowResScreenFileDataID<32>
AlteredFormStartVisualKitID<32>[3]
AlteredFormFinishVisualKitID<32>[3]
HeritageArmorAchievementID<32>
StartingLevel<32>
UiDisplayOrder<32>
Field_9_1_0_38783_033<32>
Field_9_1_0_38783_034<32>
Field_9_1_0_38783_035<32>
Field_9_1_0_38783_036<32>
PlayableRaceBit<32>
Field_9_1_0_38783_038<32>
TransmogrifyDisabledSlotMask<32>
Field_9_1_0_38783_040<32>
AlteredFormCustomizeOffsetFallback[3]
AlteredFormCustomizeRotationFallback
Field_9_1_0_38312_030[3]
Field_9_1_0_38312_031[3]
BaseLanguage<8>
CreatureType<8>
Field_9_1_0_38783_047<8>
Field_9_1_0_38783_048<8>
Field_9_1_0_38783_049<8>
Field_9_1_0_38783_050<8>`

const AreaTable = `COLUMNS
int ID
int<Map::ID> ContinentID
int<AreaTable::ID> ParentAreaID
int AreaBit
int Flags
int<SoundProviderPreferences::ID> SoundProviderPref
int<SoundProviderPreferences::ID> SoundProviderPrefUnderwater
int<SoundAmbience::ID> AmbienceID
int<ZoneMusic::ID> ZoneMusic
string ZoneName
int<ZoneIntroMusicTable::ID> IntroSound
int ExplorationLevel
locstring AreaName_lang
int FactionGroupMask
int<LiquidType::ID> LiquidTypeID
float MinElevation
float Ambient_multiplier
int<Light::ID> LightID
int MountFlags
int<ZoneIntroMusicTable::ID> UwIntroSound
int<ZoneMusic::ID> UwZoneMusic
int<SoundAmbience::ID> UwAmbience
int World_pvp_ID
int<WorldState::ID> PvpCombatWorldStateID
int WildBattlePetLevelMin
int WildBattlePetLevelMax
int<WindSettings::ID> WindSettingsID
int AreaNumber // area  = (m_AreaNumber >> 16), subarea = (m_AreaNumber & 0xFFFF)
int<AreaTable::ID> ParentAreaNum
int MIDIAmbience
int MIDIAmbienceUnderwater
int IntroPriority
int<ContentTuning::ID> ContentTuningID

BUILD 0.5.3.3368
$id$ID<32>
AreaNumber<32>
ContinentID<32>
ParentAreaNum<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
MIDIAmbience<32>
MIDIAmbienceUnderwater<32>
ZoneMusic<32>
IntroSound<32>
IntroPriority<32>
AreaName_lang

BUILD 0.5.5.3494
$id$ID<32>
AreaNumber<32>
ContinentID<32>
ParentAreaNum<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
MIDIAmbience<32>
MIDIAmbienceUnderwater<32>
ZoneMusic<32>
IntroSound<32>
IntroPriority<32>
ExplorationLevel<32>
AreaName_lang

BUILD 0.6.0.3592
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
MIDIAmbience<32>
MIDIAmbienceUnderwater<32>
ZoneMusic<32>
IntroSound<32>
IntroPriority<32>
ExplorationLevel<32>
AreaName_lang

BUILD 0.7.0.3694-0.8.0.3734
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
IntroPriority<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>

BUILD 1.11.0.5344-1.12.2.6005
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]

BUILD 2.0.0.5610-2.0.0.5666
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation

BUILD 2.0.0.5991
BUILD 1.0.0.3980-1.10.2.5302
BUILD 0.9.0.3807-0.12.0.3988
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>

LAYOUT 2AFB00FF
BUILD 2.5.4.42581, 2.5.4.42695, 2.5.4.42757, 2.5.4.42800, 2.5.4.42869, 2.5.4.42870, 2.5.4.42873, 2.5.4.42917, 2.5.4.42940, 2.5.4.43400, 2.5.4.43638, 2.5.4.43861, 2.5.4.44036, 2.5.4.44171, 2.5.4.44400, 2.5.4.44833
BUILD 2.5.3.41402, 2.5.3.41531, 2.5.3.41750, 2.5.3.41812, 2.5.3.42083, 2.5.3.42328, 2.5.3.42598
BUILD 2.5.2.39570, 2.5.2.39618, 2.5.2.39658, 2.5.2.39688, 2.5.2.39757, 2.5.2.39801, 2.5.2.39832, 2.5.2.39926, 2.5.2.40011, 2.5.2.40045, 2.5.2.40203, 2.5.2.40260, 2.5.2.40422, 2.5.2.40488, 2.5.2.40617, 2.5.2.40892, 2.5.2.41446, 2.5.2.41510
BUILD 1.14.3.42770, 1.14.3.42926, 1.14.3.43037, 1.14.3.43086, 1.14.3.43154, 1.14.3.43401, 1.14.3.43639, 1.14.3.44016, 1.14.3.44170, 1.14.3.44403, 1.14.3.44834, 1.14.3.45437, 1.14.3.46575
BUILD 1.14.2.41858, 1.14.2.41959, 1.14.2.42065, 1.14.2.42082, 1.14.2.42214, 1.14.2.42597
BUILD 1.14.1.40487, 1.14.1.40594, 1.14.1.40666, 1.14.1.40688, 1.14.1.40800, 1.14.1.40818, 1.14.1.40926, 1.14.1.40962, 1.14.1.41009, 1.14.1.41030, 1.14.1.41077, 1.14.1.41137, 1.14.1.41243, 1.14.1.41511, 1.14.1.41794, 1.14.1.42032
BUILD 1.14.0.39802, 1.14.0.39958, 1.14.0.40140, 1.14.0.40179, 1.14.0.40237, 1.14.0.40347, 1.14.0.40441, 1.14.0.40618
$noninline,id$ID<32>
ZoneName
AreaName_lang
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
AmbienceID<u16>
UwAmbience<u16>
ZoneMusic<u16>
UwZoneMusic<u16>
ExplorationLevel<8>
IntroSound<u16>
UwIntroSound<u32>
FactionGroupMask<u8>
Ambient_multiplier
MountFlags<u8>
PvpCombatWorldStateID<16>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
Flags<32>[2]
LiquidTypeID<u16>[4]

BUILD 3.0.1.8303-3.0.1.8471
BUILD 2.0.0.6080-2.4.3.8606
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier

LAYOUT 84FCA6CF
BUILD 3.4.0.43659, 3.4.0.43682, 3.4.0.43746, 3.4.0.43788, 3.4.0.43866, 3.4.0.43881, 3.4.0.43929, 3.4.0.43955, 3.4.0.44064, 3.4.0.44184, 3.4.0.44250, 3.4.0.44301, 3.4.0.44369, 3.4.0.44463, 3.4.0.44547, 3.4.0.44644, 3.4.0.44701, 3.4.0.44729, 3.4.0.44832, 3.4.0.44925, 3.4.0.44930, 3.4.0.44996, 3.4.0.45043, 3.4.0.45064, 3.4.0.45102, 3.4.0.45166, 3.4.0.45189, 3.4.0.45264, 3.4.0.45327, 3.4.0.45435, 3.4.0.45506, 3.4.0.45572, 3.4.0.45613, 3.4.0.45704, 3.4.0.45770, 3.4.0.45772, 3.4.0.45854, 3.4.0.45942, 3.4.0.46158, 3.4.0.46182, 3.4.0.46248, 3.4.0.46269, 3.4.0.46368
$noninline,id$ID<32>
ZoneName
AreaName_lang
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
AmbienceID<u16>
UwAmbience<u16>
ZoneMusic<u16>
UwZoneMusic<u16>
ExplorationLevel<8>
IntroSound<u16>
UwIntroSound<u32>
FactionGroupMask<u8>
Ambient_multiplier
MountFlags<u8>
PvpCombatWorldStateID<16>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
Flags<32>[2]
LiquidTypeID<u16>[4]

BUILD 4.0.0.11792
BUILD 4.0.0.11792-4.0.0.12232
BUILD 3.0.1.8622-3.3.5.12340
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>

BUILD 4.0.0.12266
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>
MountFlags<32>

BUILD 4.0.0.12319-4.0.6.13596
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>
MountFlags<32>
UwIntroSound<32>
UwZoneMusic<32>
UwAmbience<32>

BUILD 4.0.6.13623-4.1.0.14007
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>
MountFlags<32>
UwIntroSound<32>
UwZoneMusic<32>
UwAmbience<32>
PvpCombatWorldStateID<32>

BUILD 4.2.0.14333
BUILD 4.2.2.14522-4.3.4.15595
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>
MountFlags<32>
UwIntroSound<32>
UwZoneMusic<32>
UwAmbience<32>
World_pvp_ID<32>
PvpCombatWorldStateID<32>

BUILD 5.0.1.15464-5.0.1.15726
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
ZoneName
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>
MountFlags<32>
UwIntroSound<32>
UwZoneMusic<32>
UwAmbience<32>
World_pvp_ID<32>
PvpCombatWorldStateID<32>

BUILD 5.0.1.15739-5.0.5.16135
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>[2]
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
ZoneName
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>
MountFlags<32>
UwIntroSound<32>
UwZoneMusic<32>
UwAmbience<32>
World_pvp_ID<32>
PvpCombatWorldStateID<32>

BUILD 5.1.0.16309-5.4.8.18414
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>[2]
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
ZoneName
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>
MountFlags<32>
UwIntroSound<32>
UwZoneMusic<32>
UwAmbience<32>
World_pvp_ID<32>
PvpCombatWorldStateID<32>
WildBattlePetLevelMin<32>
WildBattlePetLevelMax<32>

BUILD 6.0.1.18179
$id$ID<32>
ContinentID<32>
ParentAreaID<32>
AreaBit<32>
Flags<32>[2]
SoundProviderPref<32>
SoundProviderPrefUnderwater<32>
AmbienceID<32>
ZoneMusic<32>
ZoneName
IntroSound<32>
ExplorationLevel<32>
AreaName_lang
FactionGroupMask<32>
LiquidTypeID<32>[4]
MinElevation
Ambient_multiplier
LightID<32>
MountFlags<32>
UwIntroSound<32>
UwZoneMusic<32>
UwAmbience<32>
World_pvp_ID<32>
PvpCombatWorldStateID<32>
WildBattlePetLevelMin<32>
WildBattlePetLevelMax<32>
WindSettingsID<32>

LAYOUT 81DA5CCB
BUILD 7.3.2.25079, 7.3.2.25163, 7.3.2.25196, 7.3.2.25208, 7.3.2.25255, 7.3.2.25326, 7.3.2.25383, 7.3.2.25442, 7.3.2.25455, 7.3.2.25477, 7.3.2.25480, 7.3.2.25497, 7.3.2.25516, 7.3.2.25549
BUILD 7.3.0.24473, 7.3.0.24484, 7.3.0.24492, 7.3.0.24500, 7.3.0.24539, 7.3.0.24563, 7.3.0.24608, 7.3.0.24651, 7.3.0.24681, 7.3.0.24692, 7.3.0.24700, 7.3.0.24715, 7.3.0.24727, 7.3.0.24730, 7.3.0.24738, 7.3.0.24744, 7.3.0.24758, 7.3.0.24759, 7.3.0.24781, 7.3.0.24793, 7.3.0.24829, 7.3.0.24834, 7.3.0.24843, 7.3.0.24845, 7.3.0.24852, 7.3.0.24853, 7.3.0.24864, 7.3.0.24878, 7.3.0.24887, 7.3.0.24896, 7.3.0.24904, 7.3.0.24920, 7.3.0.24931, 7.3.0.24956, 7.3.0.24970, 7.3.0.24974, 7.3.0.25021, 7.3.0.25195
$noninline,id$ID<32>
Flags<32>[2]
ZoneName
Ambient_multiplier
AreaName_lang
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
AmbienceID<u16>
ZoneMusic<u16>
IntroSound<u16>
LiquidTypeID<u16>[4]
UwZoneMusic<u16>
UwAmbience<u16>
PvpCombatWorldStateID<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
ExplorationLevel<8>
FactionGroupMask<u8>
MountFlags<u8>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
UwIntroSound<u32>

LAYOUT 0CA01129
BUILD 8.0.1.25902, 8.0.1.25976, 8.0.1.26010, 8.0.1.26032, 8.0.1.26095, 8.0.1.26131, 8.0.1.26141, 8.0.1.26175, 8.0.1.26231
BUILD 7.3.5.25928, 7.3.5.25937, 7.3.5.25944, 7.3.5.25946, 7.3.5.25950, 7.3.5.25961, 7.3.5.25996, 7.3.5.26124, 7.3.5.26365, 7.3.5.26654, 7.3.5.26755, 7.3.5.26822, 7.3.5.26899, 7.3.5.26972
BUILD 1.13.0.28211
$noninline,id$ID<32>
ZoneName
AreaName_lang
Flags<32>[2]
Ambient_multiplier
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
AmbienceID<u16>
ZoneMusic<u16>
IntroSound<u16>
LiquidTypeID<u16>[4]
UwZoneMusic<u16>
UwAmbience<u16>
PvpCombatWorldStateID<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
ExplorationLevel<8>
FactionGroupMask<u8>
MountFlags<u8>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
UwIntroSound<u32>

LAYOUT D5D5033B
BUILD 8.0.1.26287, 8.0.1.26297, 8.0.1.26310, 8.0.1.26321, 8.0.1.26367, 8.0.1.26433, 8.0.1.26476, 8.0.1.26491, 8.0.1.26522, 8.0.1.26530, 8.0.1.26557, 8.0.1.26567, 8.0.1.26585, 8.0.1.26604, 8.0.1.26610, 8.0.1.26624, 8.0.1.26629, 8.0.1.26637, 8.0.1.26640, 8.0.1.26683, 8.0.1.26707, 8.0.1.26714, 8.0.1.26715, 8.0.1.26734, 8.0.1.26766
$noninline,id$ID<32>
ZoneName
AreaName_lang
Flags<32>[2]
Ambient_multiplier
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
AmbienceID<u16>
UwAmbience<u16>
ZoneMusic<u16>
UwZoneMusic<u16>
IntroSound<u16>
LiquidTypeID<u16>[4]
PvpCombatWorldStateID<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
ExplorationLevel<8>
FactionGroupMask<u8>
MountFlags<u8>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
UwIntroSound<u32>

LAYOUT 22229BE7, 9944A57F
BUILD 8.3.7.34872, 8.3.7.35114, 8.3.7.35249, 8.3.7.35284, 8.3.7.35435, 8.3.7.35662
BUILD 8.3.0.32044, 8.3.0.32151, 8.3.0.32203, 8.3.0.32218, 8.3.0.32272, 8.3.0.32414, 8.3.0.32489, 8.3.0.32593, 8.3.0.32712, 8.3.0.32805, 8.3.0.32846, 8.3.0.32861, 8.3.0.32976, 8.3.0.33010, 8.3.0.33051, 8.3.0.33062, 8.3.0.33073, 8.3.0.33084, 8.3.0.33095, 8.3.0.33115, 8.3.0.33169, 8.3.0.33183, 8.3.0.33237, 8.3.0.33369, 8.3.0.33528, 8.3.0.33724, 8.3.0.33775, 8.3.0.33941, 8.3.0.34220, 8.3.0.34601, 8.3.0.34769, 8.3.0.34963
BUILD 8.2.5.31337, 8.2.5.31401, 8.2.5.31521, 8.2.5.31599, 8.2.5.31812, 8.2.5.31884, 8.2.5.31919, 8.2.5.31921, 8.2.5.31958, 8.2.5.31960, 8.2.5.31961, 8.2.5.31984, 8.2.5.32028, 8.2.5.32079, 8.2.5.32144, 8.2.5.32185, 8.2.5.32265, 8.2.5.32294, 8.2.5.32305, 8.2.5.32494, 8.2.5.32580, 8.2.5.32638, 8.2.5.32722, 8.2.5.32750, 8.2.5.32978
BUILD 8.2.0.30080, 8.2.0.30093, 8.2.0.30096, 8.2.0.30108, 8.2.0.30123, 8.2.0.30168, 8.2.0.30170, 8.2.0.30203, 8.2.0.30262, 8.2.0.30329, 8.2.0.30430, 8.2.0.30495, 8.2.0.30613, 8.2.0.30634, 8.2.0.30669, 8.2.0.30774, 8.2.0.30827, 8.2.0.30888, 8.2.0.30898, 8.2.0.30918, 8.2.0.30920, 8.2.0.30948, 8.2.0.30993, 8.2.0.31229, 8.2.0.31429, 8.2.0.31478
BUILD 8.1.5.28938, 8.1.5.29141, 8.1.5.29220, 8.1.5.29281, 8.1.5.29310, 8.1.5.29352, 8.1.5.29385, 8.1.5.29418, 8.1.5.29484, 8.1.5.29495, 8.1.5.29558, 8.1.5.29599, 8.1.5.29620, 8.1.5.29634, 8.1.5.29664, 8.1.5.29683, 8.1.5.29701, 8.1.5.29704, 8.1.5.29705, 8.1.5.29718, 8.1.5.29732, 8.1.5.29737, 8.1.5.29814, 8.1.5.29869, 8.1.5.29896, 8.1.5.29981, 8.1.5.30477, 8.1.5.30706
BUILD 8.1.0.27826, 8.1.0.27934, 8.1.0.27985, 8.1.0.28048, 8.1.0.28085, 8.1.0.28151, 8.1.0.28186, 8.1.0.28202, 8.1.0.28294, 8.1.0.28366, 8.1.0.28440, 8.1.0.28485, 8.1.0.28616, 8.1.0.28657, 8.1.0.28710, 8.1.0.28724, 8.1.0.28768, 8.1.0.28807, 8.1.0.28822, 8.1.0.28833, 8.1.0.29088, 8.1.0.29139, 8.1.0.29235, 8.1.0.29285, 8.1.0.29297, 8.1.0.29482, 8.1.0.29600, 8.1.0.29621
BUILD 8.0.1.26788, 8.0.1.26806, 8.0.1.26812, 8.0.1.26838, 8.0.1.26871, 8.0.1.26877, 8.0.1.26892, 8.0.1.26903, 8.0.1.26918, 8.0.1.26926, 8.0.1.26936, 8.0.1.26949, 8.0.1.26970, 8.0.1.26999, 8.0.1.27004, 8.0.1.27009, 8.0.1.27019, 8.0.1.27026, 8.0.1.27075, 8.0.1.27089, 8.0.1.27101, 8.0.1.27138, 8.0.1.27144, 8.0.1.27165, 8.0.1.27178, 8.0.1.27219, 8.0.1.27291, 8.0.1.27326, 8.0.1.27353, 8.0.1.27355, 8.0.1.27356, 8.0.1.27366, 8.0.1.27377, 8.0.1.27404, 8.0.1.27481, 8.0.1.27547, 8.0.1.27602, 8.0.1.27791, 8.0.1.27843, 8.0.1.27980, 8.0.1.28153
BUILD 2.5.1.38043, 2.5.1.38061, 2.5.1.38093, 2.5.1.38118, 2.5.1.38119, 2.5.1.38169, 2.5.1.38225, 2.5.1.38285, 2.5.1.38339, 2.5.1.38364, 2.5.1.38401, 2.5.1.38453, 2.5.1.38502, 2.5.1.38521, 2.5.1.38537, 2.5.1.38548, 2.5.1.38561, 2.5.1.38598, 2.5.1.38644, 2.5.1.38677, 2.5.1.38692, 2.5.1.38707, 2.5.1.38739, 2.5.1.38741, 2.5.1.38757, 2.5.1.38835, 2.5.1.38892, 2.5.1.38921, 2.5.1.38988, 2.5.1.39170, 2.5.1.39399, 2.5.1.39475, 2.5.1.39603, 2.5.1.39640
BUILD 1.13.7.37279, 1.13.7.37429, 1.13.7.37892, 1.13.7.38178, 1.13.7.38296, 1.13.7.38363, 1.13.7.38386, 1.13.7.38475, 1.13.7.38631, 1.13.7.38704, 1.13.7.39605, 1.13.7.39692
BUILD 1.13.6.36231, 1.13.6.36310, 1.13.6.36324, 1.13.6.36497, 1.13.6.36524, 1.13.6.36611, 1.13.6.36670, 1.13.6.36714, 1.13.6.36935, 1.13.6.37497
BUILD 1.13.5.34000, 1.13.5.34097, 1.13.5.34138, 1.13.5.34713, 1.13.5.34911, 1.13.5.35000, 1.13.5.35186, 1.13.5.35395, 1.13.5.35663, 1.13.5.35705, 1.13.5.35753, 1.13.5.36035, 1.13.5.36307, 1.13.5.36325
BUILD 1.13.4.33491, 1.13.4.33598, 1.13.4.33645, 1.13.4.33728, 1.13.4.33920, 1.13.4.34219, 1.13.4.34266, 1.13.4.34600, 1.13.4.34835
BUILD 1.13.3.32790, 1.13.3.32836, 1.13.3.32887, 1.13.3.33155, 1.13.3.33302, 1.13.3.33485, 1.13.3.33526
BUILD 1.13.2.30073, 1.13.2.30112, 1.13.2.30113, 1.13.2.30128, 1.13.2.30172, 1.13.2.30232, 1.13.2.30265, 1.13.2.30287, 1.13.2.30406, 1.13.2.30550, 1.13.2.30682, 1.13.2.30786, 1.13.2.30862, 1.13.2.30901, 1.13.2.30979, 1.13.2.31043, 1.13.2.31118, 1.13.2.31209, 1.13.2.31402, 1.13.2.31407, 1.13.2.31446, 1.13.2.31650, 1.13.2.31687, 1.13.2.31727, 1.13.2.31830, 1.13.2.31882, 1.13.2.32089, 1.13.2.32421, 1.13.2.32600
$noninline,id$ID<32>
ZoneName
AreaName_lang
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
AmbienceID<u16>
UwAmbience<u16>
ZoneMusic<u16>
UwZoneMusic<u16>
ExplorationLevel<8>
IntroSound<u16>
UwIntroSound<u32>
FactionGroupMask<u8>
Ambient_multiplier
MountFlags<u8>
PvpCombatWorldStateID<16>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
Flags<32>[2]
LiquidTypeID<u16>[4]

LAYOUT 7EF13BBB
BUILD 9.1.0.38312, 9.1.0.38394, 9.1.0.38511, 9.1.0.38524, 9.1.0.38549, 9.1.0.38600, 9.1.0.38627, 9.1.0.38709, 9.1.0.38783, 9.1.0.38802, 9.1.0.38872, 9.1.0.38950, 9.1.0.39015, 9.1.0.39069, 9.1.0.39121, 9.1.0.39136, 9.1.0.39172, 9.1.0.39185, 9.1.0.39226, 9.1.0.39229, 9.1.0.39262, 9.1.0.39282, 9.1.0.39289, 9.1.0.39291, 9.1.0.39318, 9.1.0.39335, 9.1.0.39413, 9.1.0.39427, 9.1.0.39497, 9.1.0.39498, 9.1.0.39584, 9.1.0.39617, 9.1.0.39653, 9.1.0.39804, 9.1.0.40000, 9.1.0.40120, 9.1.0.40443, 9.1.0.40593, 9.1.0.40725
BUILD 9.0.5.37503, 9.0.5.37623, 9.0.5.37705, 9.0.5.37774, 9.0.5.37844, 9.0.5.37862, 9.0.5.37864, 9.0.5.37893, 9.0.5.37899, 9.0.5.37988, 9.0.5.38134, 9.0.5.38556
BUILD 9.0.2.35854, 9.0.2.35938, 9.0.2.35978, 9.0.2.36037, 9.0.2.36086, 9.0.2.36165, 9.0.2.36206, 9.0.2.36267, 9.0.2.36294, 9.0.2.36401, 9.0.2.36512, 9.0.2.36532, 9.0.2.36599, 9.0.2.36639, 9.0.2.36665, 9.0.2.36671, 9.0.2.36710, 9.0.2.36734, 9.0.2.36751, 9.0.2.36753, 9.0.2.36839, 9.0.2.36949, 9.0.2.37106, 9.0.2.37130, 9.0.2.37142, 9.0.2.37176, 9.0.2.37415, 9.0.2.37474
BUILD 9.0.1.33978, 9.0.1.34003, 9.0.1.34081, 9.0.1.34098, 9.0.1.34137, 9.0.1.34199, 9.0.1.34278, 9.0.1.34324, 9.0.1.34365, 9.0.1.34392, 9.0.1.34490, 9.0.1.34615, 9.0.1.34714, 9.0.1.34821, 9.0.1.34902, 9.0.1.34972, 9.0.1.35078, 9.0.1.35167, 9.0.1.35213, 9.0.1.35256, 9.0.1.35282, 9.0.1.35360, 9.0.1.35432, 9.0.1.35482, 9.0.1.35522, 9.0.1.35598, 9.0.1.35679, 9.0.1.35707, 9.0.1.35755, 9.0.1.35789, 9.0.1.35828, 9.0.1.35853, 9.0.1.35897, 9.0.1.35917, 9.0.1.35944, 9.0.1.35989, 9.0.1.36021, 9.0.1.36036, 9.0.1.36074, 9.0.1.36131, 9.0.1.36163, 9.0.1.36200, 9.0.1.36216, 9.0.1.36228, 9.0.1.36230, 9.0.1.36247, 9.0.1.36272, 9.0.1.36286, 9.0.1.36322, 9.0.1.36372, 9.0.1.36492, 9.0.1.36577
$noninline,id$ID<32>
ZoneName
AreaName_lang
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
AmbienceID<u16>
UwAmbience<u16>
ZoneMusic<u16>
UwZoneMusic<u16>
IntroSound<u16>
UwIntroSound<u32>
FactionGroupMask<u8>
Ambient_multiplier
MountFlags<u8>
PvpCombatWorldStateID<16>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
ContentTuningID<32>
Flags<32>[2]
LiquidTypeID<u16>[4]

LAYOUT 6AA8A23B
BUILD 9.2.5.42850, 9.2.5.43022, 9.2.5.43057, 9.2.5.43254, 9.2.5.43412
BUILD 9.2.0.41089, 9.2.0.41257, 9.2.0.41360, 9.2.0.41462, 9.2.0.41726, 9.2.0.41827, 9.2.0.41962, 9.2.0.42069, 9.2.0.42174, 9.2.0.42257, 9.2.0.42277, 9.2.0.42354, 9.2.0.42399, 9.2.0.42423, 9.2.0.42488, 9.2.0.42521, 9.2.0.42538, 9.2.0.42560, 9.2.0.42578, 9.2.0.42614, 9.2.0.42698, 9.2.0.42825, 9.2.0.42852, 9.2.0.42937, 9.2.0.42979, 9.2.0.43114, 9.2.0.43206, 9.2.0.43340, 9.2.0.43345
BUILD 9.1.5.39977, 9.1.5.40071, 9.1.5.40078, 9.1.5.40196, 9.1.5.40290, 9.1.5.40383, 9.1.5.40496, 9.1.5.40622, 9.1.5.40696, 9.1.5.40738, 9.1.5.40772, 9.1.5.40843, 9.1.5.40871, 9.1.5.40906, 9.1.5.40944, 9.1.5.40966, 9.1.5.41031, 9.1.5.41079, 9.1.5.41288, 9.1.5.41323, 9.1.5.41359, 9.1.5.41488, 9.1.5.41793, 9.1.5.42010
$noninline,id$ID<32>
ZoneName
AreaName_lang
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
AmbienceID<u16>
UwAmbience<u16>
ZoneMusic<u16>
UwZoneMusic<u16>
IntroSound<u16>
UwIntroSound<u32>
FactionGroupMask<u8>
Ambient_multiplier
MountFlags<u8>
PvpCombatWorldStateID<16>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
ContentTuningID<32>
Flags<32>[2]
LiquidTypeID<u16>[4]

LAYOUT 16AE706B
BUILD 10.0.2.45480, 10.0.2.45505, 10.0.2.45569, 10.0.2.45632, 10.0.2.45698, 10.0.2.45746, 10.0.2.45779, 10.0.2.45969, 10.0.2.46091, 10.0.2.46092, 10.0.2.46144, 10.0.2.46157, 10.0.2.46259, 10.0.2.46420, 10.0.2.46456, 10.0.2.46479, 10.0.2.46619
BUILD 10.0.0.44649, 10.0.0.44707, 10.0.0.44795, 10.0.0.44895, 10.0.0.44999, 10.0.0.45141, 10.0.0.45232, 10.0.0.45335, 10.0.0.45454, 10.0.0.45570, 10.0.0.45661, 10.0.0.45697, 10.0.0.45970, 10.0.0.46047, 10.0.0.46112, 10.0.0.46181, 10.0.0.46247, 10.0.0.46270, 10.0.0.46293, 10.0.0.46313, 10.0.0.46340, 10.0.0.46366, 10.0.0.46455, 10.0.0.46547, 10.0.0.46549, 10.0.0.46597
BUILD 9.2.7.44444, 9.2.7.44767, 9.2.7.44931, 9.2.7.44981, 9.2.7.45114, 9.2.7.45161, 9.2.7.45338, 9.2.7.45745
BUILD 9.2.5.43519, 9.2.5.43630, 9.2.5.43741, 9.2.5.43810, 9.2.5.43903, 9.2.5.43971, 9.2.5.44015, 9.2.5.44061, 9.2.5.44127, 9.2.5.44232, 9.2.5.44325, 9.2.5.44730, 9.2.5.44908
$noninline,id$ID<32>
ZoneName
AreaName_lang
ContinentID<u16>
ParentAreaID<u16>
AreaBit<16>
SoundProviderPref<u8>
SoundProviderPrefUnderwater<u8>
AmbienceID<u16>
UwAmbience<u16>
ZoneMusic<u16>
UwZoneMusic<u16>
IntroSound<u16>
UwIntroSound<u32>
FactionGroupMask<u8>
Ambient_multiplier
MountFlags<u8>
PvpCombatWorldStateID<16>
WildBattlePetLevelMin<u8>
WildBattlePetLevelMax<u8>
WindSettingsID<u8>
ContentTuningID<32>
Flags<32>[2]
LiquidTypeID<u16>[4]`

func TestDecode(t *testing.T) {
	def, err := DecodeDefinition("ChrRaces", strings.NewReader(ChrRaces))
	if err != nil {
		t.Fatal(err)
	}

	cl := def.Column("ID")

	if cl.HintBits != 32 {
		t.Fatal("fail")
	}

	log.Dump("def", def)

	def, err = DecodeDefinition("AreaTable", strings.NewReader(AreaTable))
	if err != nil {
		t.Fatal(err)
	}

	cl = def.Column("ID")

	if cl.HintBits != 32 {
		t.Fatal("fail")
	}

	log.Dump("def", def)
	time.Sleep(3 * time.Second)
}
