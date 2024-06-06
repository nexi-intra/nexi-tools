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
        "github.com/magicbutton/nexi-tools/services/endpoints/icon"
                    "github.com/magicbutton/nexi-tools/services/models/iconmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestIconupdate(t *testing.T) {
                                // transformer v1
            object := iconmodel.Icon{}
         
            result,err := icon.IconUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
