package mondia_eg

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

// mondia 埃及  LP页面
type MondiaEGPageController struct {
	beego.Controller
}

// LP首页
func (c *MondiaEGPageController) LpIndex() {
	requesParse := strings.Split(c.Ctx.Request.URL.String(), "?")
	requesData := ""
	if len(requesParse) == 2 {
		requesData = requesParse[1] + "&type=KKP"
	} else {
		requesData = "type=KKP"
	}

	trackID := requestTrackID(requesData)
	c.Data["trackID"] = trackID
	c.TplName = "mondia/eg/LP.html"
}

// 条款页面
func (c *MondiaEGPageController) Terms() {
	requesParse := strings.Split(c.Ctx.Request.URL.String(), "?")
	requesData := ""
	if len(requesParse) == 2 {
		requesData = requesParse[1] + "&type=KKP"
	} else {
		requesData = "type=KKP"
	}

	trackID := requestTrackID(requesData)
	c.Data["trackID"] = trackID
	c.TplName = "mondia/eg/terms.html"
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