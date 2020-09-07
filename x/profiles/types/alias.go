package types

// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profiles/types/models"
	"github.com/desmos-labs/desmos/x/profiles/types/msgs"
)

const (
	ModuleName          = models.ModuleName
	RouterKey           = models.RouterKey
	StoreKey            = models.StoreKey
	ActionSaveProfile   = models.ActionSaveProfile
	ActionDeleteProfile = models.ActionDeleteProfile
	QuerierRoute        = models.QuerierRoute
	QueryProfile        = models.QueryProfile
	QueryProfiles       = models.QueryProfiles
	QueryParams         = models.QueryParams
)

var (
	// functions aliases
	ProfileStoreKey       = models.ProfileStoreKey
	DtagStoreKey          = models.DtagStoreKey
	NewProfile            = models.NewProfile
	NewProfiles           = models.NewProfiles
	NewPictures           = models.NewPictures
	RegisterModelsCodec   = models.RegisterModelsCodec
	NewMsgSaveProfile     = msgs.NewMsgSaveProfile
	NewMsgDeleteProfile   = msgs.NewMsgDeleteProfile
	RegisterMessagesCodec = msgs.RegisterMessagesCodec

	// variable aliases
	ProfileStorePrefix = models.ProfileStorePrefix
	DtagStorePrefix    = models.DtagStorePrefix
	ModelsCdc          = models.ModelsCdc
	MsgsCodec          = msgs.MsgsCodec
)

type (
	MsgSaveProfile   = msgs.MsgSaveProfile
	MsgDeleteProfile = msgs.MsgDeleteProfile
	Profile          = models.Profile
	Profiles         = models.Profiles
	Pictures         = models.Pictures
)