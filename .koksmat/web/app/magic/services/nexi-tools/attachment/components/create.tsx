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
    import {AttachmentForm} from "./form";
    
    import {AttachmentItem} from "../applogic/model";
    export default function CreateAttachment() {
       
        const attachment = {} as AttachmentItem;
        return (
          <div>{attachment && 
          <AttachmentForm attachment={attachment} editmode="create"/>}
         
          </div>
        );
    }
