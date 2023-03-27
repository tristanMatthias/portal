import { useEffect, useState } from "react";
import { ActionModelsList } from "../../../wailsjs/go/model/model";
import { model } from "../../../wailsjs/go/models";
import ModelCard from "./Model.card";

export default function ModelsGrid() {
  const [models, setModels] = useState<model.EModel[]>([])
  useEffect(() => {
    ActionModelsList().then(setModels)
  }, []);

  return <div className="grid-models">
    {models.map((model) => <ModelCard model={model} key={model.id} />)}
  </div>
}
