    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/koksmat/useservice";
    import { useState } from "react";
    import {IconForm} from "./form";
    
    import {IconItem} from "../applogic/model";
    export default function CreateIcon() {
       
        const icon = {} as IconItem;
        return (
          <div>{icon && 
          <IconForm icon={icon} editmode="create"/>}
         
          </div>
        );
    }
