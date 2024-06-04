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
import {RegionForm} from "./form";

import {RegionItem} from "../applogic/model";
export default function UpdateRegion(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<RegionItem>(
      "nexi-tools.region",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const region = readResult.data;
    return (
      <div>{region && 
      <RegionForm region={region} editmode="update"/>}
     
      </div>
    );
}
