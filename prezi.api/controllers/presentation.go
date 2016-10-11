package controllers

import (
	"github.com/astaxie/beego"
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.api.contracts/response"
	serviceLayer "infra-balaji-rao/prezi.core/service"
)

type PresentationController struct {
	beego.Controller
}

/*Get returns presentations based on the filters & options*/
func (presentationController *PresentationController) Get() {
	request := presentationController.generateRequest()
	service := serviceLayer.NewPresentationService()

	presentations, _ := service.Get(request)
	count := service.Count(request)
	numberOfItems := count

	if request.PaginationOption != nil {
		numberOfItems = request.PaginationOption.NumberOfItems
	}

	presentationController.Data["json"] = response.PaginatedResponse{
		Results:      presentations,
		TotalRecords: count,
		TotalPages:   count / numberOfItems,
	}

	presentationController.ServeJSON()
}

/*generateRequest() parses the HTTP request for the relevant params*/
func (presentationController *PresentationController) generateRequest() request.Request {
	presentationRequest := request.Request{}

	paginate, _ := presentationController.GetBool("paginate")

	if paginate {
		index, _ := presentationController.GetInt("index")
		numberOfItems, _ := presentationController.GetInt("numitems")

		presentationRequest.PaginationOption = &request.PaginationOption{
			Index:         index,
			NumberOfItems: numberOfItems,
		}
	}

	sort, _ := presentationController.GetBool("sort")

	if sort {
		direction, _ := presentationController.GetInt("direction")

		presentationRequest.SortingOption = &request.SortingOption{
			Direction: direction,
			Field:     presentationController.GetString("sortby"),
		}
	}

	return presentationRequest
}
