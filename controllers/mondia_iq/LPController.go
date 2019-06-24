package mondia_iq

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

// mondia 埃及  LP页面
type MondiaIQPageController struct {
	beego.Controller
}

// LP首页
func (c *MondiaIQPageController) LpIndex() {
	fmt.Println(c.Ctx.Request.URL.String())
	requesParse := strings.Split(c.Ctx.Request.URL.String(), "?")
	requesData := ""
	if len(requesParse) == 2 {
		requesData = requesParse[1] + "&type=KKP_IQ"
	} else {
		requesData = "type=KKP_IQ"
	}

	trackID := requestTrackID(requesData)
	if trackID == "false" {
		// 404
		c.StopRun()
	}
	c.Data["trackID"] = trackID
	c.TplName = "mondia/iq/LP.html"
}

// 条款页面
func (c *MondiaIQPageController) Terms() {
	requesParse := strings.Split(c.Ctx.Request.URL.String(), "?")
	requesData := ""
	if len(requesParse) == 2 {
		requesData = requesParse[1] + "&type=KKP_IQ"
	} else {
		requesData = "type=KKP_IQ"
	}

	trackID := requestTrackID(requesData)
	c.Data["trackID"] = trackID
	c.TplName = "mondia/iq/terms.html"
}

func requestTrackID(requesParse string) (trackID string) {
	resp, err := http.Get(fmt.Sprintf("http://mm-eg.leadernethk.com/returnid?" + requesParse))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return string(body)
}
