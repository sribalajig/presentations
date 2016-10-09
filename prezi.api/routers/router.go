// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"infra-balaji-rao/prezi.api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/presentation", &controllers.PresentationController{}, "get:Get")
}
