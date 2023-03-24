import React, { useState, useEffect } from "react";
import { ActionModelsList } from "../../../wailsjs/go/model/model";
import ModelCard from "./Model.card";

export default function ModelsGrid() {
  const [models, setModels] = useState<string[]>([])
  useEffect(() => {
    ActionModelsList().then(setModels)
  }, []);

  return <div className="grid-models">
    {models.map((model) => <ModelCard model={model} key={model} />)}
  </div>
}
