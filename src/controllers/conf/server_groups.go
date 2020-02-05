package confcontrollers

import (
	"github.com/linclin/gopub/src/controllers"
	"github.com/linclin/gopub/src/library/jumpserver"
)

type ServerGroupsController struct {
	controllers.BaseController
}

func (c *ServerGroupsController) Get() {
	if c.User == nil || c.User.Id == 0 {
		c.SetJson(2, nil, "not login")
		return
	}
	group2id, _ := jumpserver.GetGroups()
	c.SetJson(0, group2id, "")
	return
}
