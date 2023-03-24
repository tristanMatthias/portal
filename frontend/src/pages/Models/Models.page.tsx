import { Button } from 'base/components/Button/Button';
import { TextField } from 'base/components/TextField/TextField';
import { useEffect, useState } from "react";
import { ActionSearchModels } from "../../../wailsjs/go/huggingface/Huggingface";
import Page from "../../components/Page/Page";
import ModelsGrid from "./Models.grid";
import { huggingface } from '../../../wailsjs/go/models';

export default function PageModels() {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState<huggingface.HFModel[]>([]);

  useEffect(() => {
    if (!query) return;
    ActionSearchModels(query).then(setResults);
  }, [query]);

  return <Page id="models" title="Models">
    {results?.map((r) => <div key={r.id}>{r.id}</div>)}
    <TextField onChange={e => setQuery(e.target.value)} />
    <Button text="Install" />
    <ModelsGrid />
  </Page>
}
