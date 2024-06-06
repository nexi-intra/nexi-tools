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
        "github.com/magicbutton/nexi-tools/services/endpoints/translation"
                    "github.com/magicbutton/nexi-tools/services/models/translationmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestTranslationupdate(t *testing.T) {
                                // transformer v1
            object := translationmodel.Translation{}
         
            result,err := translation.TranslationUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
