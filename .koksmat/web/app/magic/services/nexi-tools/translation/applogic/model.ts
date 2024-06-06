    
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
// Translation
export interface TranslationItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    language_id : number ;
    translation : string ;
    tablename : string ;
    fieldname : string ;
    recordid : number ;

}


// Translation
export const TranslationSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    language_id : z.number(), 
    translation : z.string(), 
    tablename : z.string(), 
    fieldname : z.string(), 
    recordid : z.number(), 

});

