    
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
// Icon
export interface IconItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    icon : string ;
    base64 : string ;

}


// Icon
export const IconSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    icon : z.string(), 
    base64 : z.string(), 

});

