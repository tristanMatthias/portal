export interface ModelCardProps {
  model: model.EModel;
}

import { useState } from "react";
import { downloadModel } from "../../lib/models";
import "./Model.card.scss";
import { model } from "../../../wailsjs/go/models";

export default function ModelCard({
  model
}: ModelCardProps) {
  const [downloadState, setDownloadState] = useState<any | null>(null);

  const download = () => downloadModel(model.name, setDownloadState);

  return <div className="card-model">
    {model.name}
  </div>
}
