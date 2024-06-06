    
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
// Tool
export interface ToolItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    category_id : number ;
    url : string ;
    status : string ;
    icon_id : number ;

}


// Tool
export const ToolSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    category_id : z.number(), 
    url : z.string(), 
    status : z.string().optional(), 
    icon_id : z.number().optional(), 

});

