/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/koksmat/useservice";
import { useState } from "react";
import {CategoryForm} from "./form";

import {CategoryItem} from "../applogic/model";
export default function UpdateCategory(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<CategoryItem>(
      "nexi-tools.category",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const category = readResult.data;
    return (
      <div>{category && 
      <CategoryForm category={category} editmode="update"/>}
     
      </div>
    );
}
