/*
	The Service frontend is UI on a page. Each Service has a ServiceRepository
	and its usecases data. The architecture groundwork is the usecases that
	orchestrate everything relevant to the frontend. The usecases do not
	need to know underlying database nor transpotation layer. 
	At directory we run ~/usecases$ go test
*/
package usecases

import (
//	"github.com/khaiphong/mu/backend/entities"
)

func GetItems() {
}

/*
	The OmHub usecases may contain (1) coded operations of user's Pod, (2) GSLP
	operations, (3) local operations, (4) MuCo, (5) synchonization of user's
	online / offline data, (6) event processing and composition, (7) custom AI,
	(8) Open CI, (9) back office data centres management.
*/