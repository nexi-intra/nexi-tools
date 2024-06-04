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
        "github.com/magicbutton/nexi-tools/services/endpoints/category"
                    "github.com/magicbutton/nexi-tools/services/models/categorymodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestCategorycreate(t *testing.T) {
                                // transformer v1
            object := categorymodel.Category{}
         
            result,err := category.CategoryCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
