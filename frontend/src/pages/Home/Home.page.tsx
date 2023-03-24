import { useEffect, useMemo, useState } from "react";
import { ActionChat } from "../../../wailsjs/go/chat/Chat";
import { downloadModel } from "../../lib/models";
import './Home.page.scss';
import Sidebar from "./Sidebar/Sidebar";
import Page from "../../components/Page/Page";
import { ActionModelsList } from "../../../wailsjs/go/model/model";

export default function PageHome() {
  const [downloadState, setDownloadState] = useState<any | null>(null);
  const [model, setModel] = useState<string | null>(null);
  const [models, setModels] = useState<string[]>([]);

  const [prompt, setPrompt] = useState<string | null>(null);
  const [chatHistory, setChatHistory] = useState<string[]>([]);

  const download = () => {
    if (!model) return;
    downloadModel(model, setDownloadState);
  }

  const send = () => {
    if (!prompt) return;
    ActionChat("gpt2", prompt).then(r => {
      setChatHistory(chatHistory => [...chatHistory, r.response]);
    });
    setChatHistory(chatHistory => [...chatHistory, prompt]);
    setPrompt("");
  }

  useEffect(() => {
    ActionModelsList().then(setModels)
  }, []);

  return <Page
    id="home"
    title="Home"
    sidebar={<Sidebar />}
  >
    <h2>Download a model</h2>
    <input
      type="text"
      placeholder="Model to download"
      onChange={e => setModel(e.target.value)}
    />

    <button onClick={download} disabled={!model}>
      Download a model
    </button>

    <pre>
      {JSON.stringify(downloadState, null, 2)}
    </pre>

    <hr />

    <h2>Downloaded models</h2>
    <ul>
      {models.map((model, i) => <li key={i}>{model}</li>)}
    </ul>

    <h2>Chat with GPT2</h2>

    <input
      type="text"
      onChange={e => setPrompt(e.target.value)}
      placeholder="Prompt"
    />
    <button onClick={send}>Send</button>
    <ul>
      {chatHistory.map((chat, i) => <li key={i}>{chat}</li>)}
    </ul>
  </Page>;
}
