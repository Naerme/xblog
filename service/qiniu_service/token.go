package qiniu_service

import (
	"blogx_server/global"
	"context"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"time"
)

func GenToken() (token string, err error) {
	mac := credentials.NewCredentials(global.Conifg.QiNiu.AccessKey, global.Conifg.QiNiu.SecretKey)
	putPolicy, err := uptoken.NewPutPolicy(global.Conifg.QiNiu.Bucket, time.Now().Add(time.Duration(global.Conifg.QiNiu.Expiry)*time.Second))
	if err != nil {
		return
	}
	token, err = uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
	if err != nil {
		return
	}
	return
}
