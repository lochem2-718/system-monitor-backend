package models

type IConfigProvider interface {
	GetSetting(key string) string
}
