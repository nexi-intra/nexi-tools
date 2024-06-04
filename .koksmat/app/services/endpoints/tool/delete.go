            /*
            File have been automatically created. To prevent the file from getting overwritten
            set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
            ---
            keep: false
            ---
            */
            //generator:  noma3.delete.v2
            package tool
            
            import (
                "log"
                "strconv"
                "github.com/magicbutton/nexi-tools/applogic"
                "github.com/magicbutton/nexi-tools/database"
                "github.com/magicbutton/nexi-tools/services/models/toolmodel"
            
            )
            
            func ToolDelete(arg0 string) ( error) {
                id,_ := strconv.Atoi(arg0)
                log.Println("Calling Tooldelete")
            
                return applogic.Delete[database.Tool, toolmodel.Tool](id)
            
            }
