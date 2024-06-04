/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
package magicapp

import (
	"github.com/magicbutton/nexi-tools/services"
	"github.com/nats-io/nats.go/micro"
)

func RegisterServiceEndpoints(root micro.Group) {
    root.AddEndpoint("category", micro.HandlerFunc(services.HandleCategoryRequests))
        root.AddEndpoint("tool", micro.HandlerFunc(services.HandleToolRequests))
        root.AddEndpoint("region", micro.HandlerFunc(services.HandleRegionRequests))
        root.AddEndpoint("country", micro.HandlerFunc(services.HandleCountryRequests))
    }
