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
    import {ToolForm} from "./form";
    
    import {ToolItem} from "../applogic/model";
    export default function CreateTool() {
       
        const tool = {} as ToolItem;
        return (
          <div>{tool && 
          <ToolForm tool={tool} editmode="create"/>}
         
          </div>
        );
    }
