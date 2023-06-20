package core

import "github.com/TrHung-297/fountain/baselib/g_log"

type CoreController interface {
	InstallController()
	RegisterCallback(cb interface{})
}

type CoreControllerAfterInstalled interface {
	AfterInstalledDone()
}

var controllers = []CoreController{}

func InstallCoreControllers() []CoreController {
	//RegisterExtendControllers(f_core.Controllers)
	g_log.V(3).Infof("InstallCoreControllers Num: %d", len(controllers))

	for _, c := range controllers {
		// g_log.V(3).Infof("InstallCoreControllers Controller: %+v", c)
		c.InstallController()
		for _, c2 := range controllers {
			if c != c2 {
				c.RegisterCallback(c2)
			}
		}
	}

	for _, c := range controllers {
		if f, ok := c.(CoreControllerAfterInstalled); ok {
			f.AfterInstalledDone()
			g_log.V(1).Infof("InstallCoreControllers - %T AfterInstalledDone called", c)
		}
	}

	return controllers
}