    
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
// User
export interface UserItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    email : string ;
    firstname : string ;
    lastname : string ;
    language_id : number ;
    country_id : number ;
    region_id : number ;
    status : string ;

}


// User
export const UserSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    email : z.string(), 
    firstname : z.string(), 
    lastname : z.string(), 
    language_id : z.number(), 
    country_id : z.number(), 
    region_id : z.number(), 
    status : z.string(), 

});

