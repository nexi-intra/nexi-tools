
    "use client";

    import { Input } from "@/components/ui/input";
import { useService } from "@/koksmat/useservice";
import { useSQLSelect } from "@/koksmat/usesqlselect";
import { set } from "date-fns";
import { tr } from "date-fns/locale";
import { useMemo, useState } from "react";
import CategorySmallCard from "./smallcard";
import {CategoryItem} from "../applogic/model";

  
    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    /* guldbar */

   

    export interface Root {
        totalPages: number
        totalItems: number
        currentPage: number
        items: CategoryItem[]
      }
   

    export default function SearchCategory(
        props: {
          onItemClick? : (item:CategoryItem) => void
        }

    ) {
        const [transactionId, settransactionId] = useState(0);
        const search = useMemo(() => {
          return { text: "" };
        }, []);
        const searchResult = useService<Root>(
          "nexi-tools.category",
          ["search", "%" + search.text + "%"],
          "",
          6000,
          transactionId.toString()
        );
        return (
          <div>
            Search
            <Input
              type="text"
              value={search.text}
              className="mx-2"
              onChange={(e) => {
                search.text = e.target.value;
                settransactionId(transactionId + 1);
              }}
            />
            {searchResult?.error && (
                <div className="text-red-500">Error: {searchResult.error}</div>
              )}
              {searchResult?.isLoading && <div>Loading</div>}
              {searchResult?.data && (
                <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 ">
                  {searchResult.data.items.map((item, index) => {
                    return <div key={index}>
                    <CategorySmallCard item={item} 
                  
                    onClick={()=>{
                        if (props.onItemClick) {
                            props.onItemClick(item)
                    }}}
                    />

                    
                    
                    
                    </div>;
                  })}
                </div>
              )}
        
         
          </div>
    );
    }
        
