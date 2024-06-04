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
    import {CategoryForm} from "./form";
    
    import {CategoryItem} from "../applogic/model";
    export default function CreateCategory() {
       
        const category = {} as CategoryItem;
        return (
          <div>{category && 
          <CategoryForm category={category} editmode="create"/>}
         
          </div>
        );
    }
