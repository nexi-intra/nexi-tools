            /*
            File have been automatically created. To prevent the file from getting overwritten
            set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
            ---
            keep: false
            ---
            */
            //generator:  noma3.delete.v2
            package icon
            
            import (
                "log"
                "strconv"
                "github.com/magicbutton/nexi-tools/applogic"
                "github.com/magicbutton/nexi-tools/database"
                "github.com/magicbutton/nexi-tools/services/models/iconmodel"
            
            )
            
            func IconDelete(arg0 string) ( error) {
                id,_ := strconv.Atoi(arg0)
                log.Println("Calling Icondelete")
            
                return applogic.Delete[database.Icon, iconmodel.Icon](id)
            
            }
