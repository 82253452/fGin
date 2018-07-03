package controller

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"github.com/tidwall/gjson"
	. "fGin/ov"
)

var baseUrl = "http://fbservice.qiuzhang.com"

func GetSaichengList(c *gin.Context) {
	c.JSON(http.StatusOK, getsaichenng())
}
func GetJifenList(c *gin.Context) {
	c.String(http.StatusOK, getjifenn())
}
func getFootBallList() {
	resp, _ := http.Get("http://www.maimaiche.com/loginRegister/login.do")
	body, _ := ioutil.ReadAll(resp.Body)
	print(string(body))
}
func getSessionId() string {
	resp, _ := http.Get(baseUrl + "/api/data/footballrank/GetSeasonList?leagueId=5a698341360c3f41541b3bb2")
	body, _ := ioutil.ReadAll(resp.Body)
	sessionid := gjson.GetBytes(body, "result.0.id").String()
	return sessionid
}

func getsaichenng() []FootballInfo {
	re := []FootballInfo{}
	resp, _ := http.Get(baseUrl + "/api/data/footballrank/GetRoundList?leagueId=5a698341360c3f41541b3bb2&seasonId=" + getSessionId())
	body, _ := ioutil.ReadAll(resp.Body)
	for _, bodyArray := range gjson.GetBytes(body, "result").Array() {
		roundId := bodyArray.Get("id").String()
		name := bodyArray.Get("name").String()
		selected := bodyArray.Get("selected").String()

		resp2, _ := http.Get(baseUrl + "/api/data/footballrank/GetScheduleList?leagueId=5a698341360c3f41541b3bb2&seasonId=" + getSessionId() + "&roundId=" + roundId + "&pageNumber=1&pageSize=20")
		body2, _ := ioutil.ReadAll(resp2.Body)
		body2Sc := []FootballNatchDateInfo{}
		for _, bodyArray2 := range gjson.GetBytes(body2, "result.result").Array() {
			dateInfo := FootballNatchDateInfo{Date: bodyArray2.Get("date").String()}
			s := []FootballNatchInfo{}
			for _, bodyArray3 := range bodyArray2.Get("scheduleList").Array() {
				ss := FootballNatchInfo{Id: bodyArray3.Get("id").String(),
					Date:                   bodyArray3.Get("date").String(),
					StartTime:              bodyArray3.Get("startTime").String(),
					HomeName:               bodyArray3.Get("homeName").String(),
					HomeLogo:               bodyArray3.Get("homeLogo").String(),
					SeasonId:               bodyArray3.Get("seasonId").String(),
					RoundId:                bodyArray3.Get("roundId").String(),
					HasProcessLogo:         bodyArray3.Get("hasProcessLogo").String(),
					GuestName:              bodyArray3.Get("guestName").String(),
					GuestLogo:              bodyArray3.Get("guestLogo").String(),
					GuestGoal:              bodyArray3.Get("guestGoal").String(),
					LeagueId:               bodyArray3.Get("leagueId").String(),
				}
				s = append(s, ss)
			}
			dateInfo.Info = s
			body2Sc = append(body2Sc, dateInfo)
		}
		ff := FootballInfo{Name: name, Id: roundId, Selected: selected, Info: body2Sc}
		re = append(re, ff)
	}
	return re
}
func getjifenn() string {
	resp, _ := http.Get(baseUrl + "/api/data/footballrank/GetOuGuanTeamScoreRankList?leagueId=5a698341360c3f41541b3bb2&seasonId=" + getSessionId())
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

