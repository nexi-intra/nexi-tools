    
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/       
"use client";
import { z } from "zod";
// spunk
// Attachment
export interface AttachmentItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    sortorder : string ;
    filename : string ;
    link : string ;
    mimetype : string ;
    tool_id : number ;

}


// Attachment
export const AttachmentSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    sortorder : z.string(), 
    filename : z.string(), 
    link : z.string(), 
    mimetype : z.string(), 
    tool_id : z.number(), 

});

