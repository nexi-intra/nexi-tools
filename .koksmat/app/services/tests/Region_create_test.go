    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
    //generator:  noma4.1
    package tests
    import (
        "testing"
        "github.com/magicbutton/nexi-tools/services/endpoints/region"
                    "github.com/magicbutton/nexi-tools/services/models/regionmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestRegioncreate(t *testing.T) {
                                // transformer v1
            object := regionmodel.Region{}
         
            result,err := region.RegionCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
