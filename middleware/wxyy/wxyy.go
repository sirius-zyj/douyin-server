package wxyy

import (
	base64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type wxyyInstance struct {
	accessToken string
}

var (
	instance *wxyyInstance
	once     sync.Once
)

// GetInstance 返回单例对象的实例
func GetInstance() *wxyyInstance {
	once.Do(func() {
		instance = &wxyyInstance{
			accessToken: GetAccessToken(),
		}
	})
	return instance
}

func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s",
		"SN7Floc182ieoakGdCXZOVo9", "wWn8EtsvTafEsNnE21BesyFtong6K9Yb")
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}

type AudioRequestBody struct {
	Format  string `json:"format"`
	Rate    int    `json:"rate"`
	Channel int    `json:"channel"`
	Cuid    string `json:"cuid"`
	Token   string `json:"token"`
	DevPid  int    `json:"dev_pid"`
	Speech  string `json:"speech"`
	Len     int    `json:"len"`
}

type AudioRespBody struct {
	ErrNo    int      `json:"err_no"`
	ErrMsg   string   `json:"err_msg"`
	CorpurNo string   `json:"corpus_no"`
	Sn       string   `json:"sn"`
	Result   []string `json:"result"`
}

func (it *wxyyInstance) Video2Txt(videoFilePath string) (text []string, err error) {
	filename := filepath.Base(videoFilePath)
	extension := filepath.Ext(videoFilePath)
	audioFilePath := filename[:len(filename)-len(extension)] + ".mp3"
	err = ffmpeg.Input(videoFilePath).
		Output(audioFilePath, ffmpeg.KwArgs{"acodec": "libmp3lame"}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		return
	}
	defer os.Remove(audioFilePath)

	audioData, _ := ioutil.ReadFile(audioFilePath)
	len := len(audioData)
	speech := base64.StdEncoding.EncodeToString(audioData)

	url := "https://vop.baidu.com/pro_api"
	bodyFmt := AudioRequestBody{
		Format:  "m4a",
		Rate:    16000,
		Channel: 1,
		Cuid:    "0zlmXj5Pym8EcTLE1Wbkfn3qHLTZhIvt",
		Token:   it.accessToken,
		DevPid:  80001,
		Speech:  speech,
		Len:     len,
	}

	srcByte, err := json.Marshal(bodyFmt)
	payload := strings.NewReader(string(srcByte))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	textObj := &AudioRespBody{}
	json.Unmarshal([]byte(body), &textObj)
	if textObj.ErrMsg != "success." {
		return text, fmt.Errorf(textObj.ErrMsg)
	}
	text = textObj.Result
	return
}

type KeywordRequestBody struct {
	Text []string `json:"text"`
	Num  int      `json:"num"`
}

type Keyword struct {
	Score float64 `json:"score"`
	Word  string  `json:"word"`
}
type KeywordRespBody struct {
	LogId   int       `json:"log_id"`
	ErrCode int       `json:"error_code"`
	ErrMsg  string    `json:"error_msg"`
	Results []Keyword `json:"results"`
}

func (it *wxyyInstance) Txt2Keyword(text []string) (keyword []Keyword, err error) {
	url := "https://aip.baidubce.com/rpc/2.0/nlp/v1/txt_keywords_extraction?access_token=" + it.accessToken
	bodyFmt := KeywordRequestBody{
		Text: text,
		Num:  5,
	}
	srcByte, err := json.Marshal(bodyFmt)
	payload := strings.NewReader(string(srcByte))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	textObj := &KeywordRespBody{}
	json.Unmarshal([]byte(body), &textObj)
	if textObj.ErrCode != 0 {
		return keyword, fmt.Errorf("Txt2Gorse Failed: %s", textObj.ErrMsg)
	}
	keyword = textObj.Results
	return
}

type summaryRequestBody struct {
	Content       string `json:"content"`
	MaxSummaryLen int    `json:"max_summary_len"`
}

type summaryRespBody struct {
	LogId   int    `json:"log_id"`
	ErrCode int    `json:"error_code"`
	ErrMsg  string `json:"error_msg"`
	Summary string `json:"summary"`
}

func (it *wxyyInstance) Txt2Summary(text []string) (summary string, err error) {
	url := "https://aip.baidubce.com/rpc/2.0/nlp/v1/news_summary?charset=UTF-8&access_token=" + it.accessToken
	bodyFmt := summaryRequestBody{
		Content:       strings.Join(text, "\n"),
		MaxSummaryLen: 30,
	}
	srcByte, err := json.Marshal(bodyFmt)
	payload := strings.NewReader(string(srcByte))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	textObj := &summaryRespBody{}
	json.Unmarshal([]byte(body), &textObj)
	if textObj.ErrCode != 0 {
		return summary, fmt.Errorf("Txt2Gorse Failed: %s", textObj.ErrMsg)
	}
	summary = textObj.Summary
	return
}
