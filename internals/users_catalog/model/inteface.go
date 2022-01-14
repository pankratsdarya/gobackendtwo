package model

import "github.com/google/uuid"

type CatalogItem interface {
	AddUser(User) error
	AddSetting(Setting) error
	GetUser(uuid.UUID) (User, error)
	GetSetting(uuid.UUID) (Setting, error)
	AddUserToSetting(userID, settingID uuid.UUID) error
	ExcludeUserFromSetting(userID, settingID uuid.UUID) error
	SearchUserByName(string) []User
	SearchUserBySetting(Setting) []User
	SearchSettingByName(string) []Setting
	SearchSettingByUsers([]User) []Setting
}
