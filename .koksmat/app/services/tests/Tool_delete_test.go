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
        "github.com/magicbutton/nexi-tools/services/endpoints/tool"
        
        "github.com/stretchr/testify/assert"
    )
    
    func TestTooldelete(t *testing.T) {
                // noma4.1.1
        
        err := tool.ToolDelete(".")
        if err != nil {
            t.Errorf("Error %s", err)
        }
        assert.True(t, true) // for additional tests
       
        
    
    }
    