export interface ModelCardProps {
  model: string
}

import { useState } from "react";
import { downloadModel } from "../../lib/models";
import "./Model.card.scss";

export default function ModelCard({
  model
}: ModelCardProps) {
  const [downloadState, setDownloadState] = useState<any | null>(null);

  const download = () => downloadModel(model, setDownloadState);

  return <div className="card-model">
    {model}
  </div>
}
