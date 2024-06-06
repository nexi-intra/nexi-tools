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
import {LanguageForm} from "./form";

import {LanguageItem} from "../applogic/model";
export default function UpdateLanguage(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<LanguageItem>(
      "nexi-tools.language",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const language = readResult.data;
    return (
      <div>{language && 
      <LanguageForm language={language} editmode="update"/>}
     
      </div>
    );
}
