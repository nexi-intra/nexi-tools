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
import {TranslationForm} from "./form";

import {TranslationItem} from "../applogic/model";
export default function UpdateTranslation(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TranslationItem>(
      "nexi-tools.translation",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const translation = readResult.data;
    return (
      <div>{translation && 
      <TranslationForm translation={translation} editmode="update"/>}
     
      </div>
    );
}
