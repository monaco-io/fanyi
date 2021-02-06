package src

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
)

type Response struct {
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []TransResult `json:"trans_result"`
	ErrorCode   int64         `json:"error_code"`
}

type TransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type Query struct {
	Q     string `json:"q"`
	From  string `json:"from"`
	To    string `json:"to"`
	AppID string `json:"appid"`
	Salt  string `json:"salt"`
	Sign  string `json:"sign"`
}

func (q *Query) BuildFromAndTo() {
	if regexp.MustCompile("[a-zA-Z]+").MatchString(q.Q) {
		q.From = "en"
		q.To = "zh"
		return
	}
	q.From = "zh"
	q.To = "en"
}

func (q Query) ToMap() map[string]string {
	mp := make(map[string]string)

	mp["q"] = q.Q
	mp["from"] = q.From
	mp["to"] = q.To
	mp["appid"] = q.AppID
	mp["salt"] = q.Salt
	mp["sign"] = q.Sign

	return mp
}

func Encode(word string) map[string]string {
	var query Query

	query.Q = word
	query.BuildFromAndTo()
	query.AppID = AppID
	query.Salt = fmt.Sprint(rand.Int())

	// sum md5
	signString := query.AppID + query.Q + query.Salt + AppSecrect
	h := md5.New()
	h.Write([]byte(signString))
	query.Sign = hex.EncodeToString(h.Sum(nil))

	return query.ToMap()
}
