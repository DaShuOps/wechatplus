package wechatplus

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGetTokenByCode(t *testing.T) {
	code := "rtarstarstratarstarst"

	token, err := GetTokenByCode(miniAppId, miniAppSecret, code)
	if err != nil {
		t.Fatal(err)
	}

	log.Debug(token)
}

func TestGetUserInfoByToken(t *testing.T) {
	accessToken := "rtarstarstratarstarst"
	openId := "rtarstarstratarstarst"

	userInfo, err := GetUserInfoByToken(accessToken, openId)
	if err != nil {
		t.Fatal(err)
	}

	log.Debug(userInfo)
}
