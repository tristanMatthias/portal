import { useEffect, useState } from "react";
import { EventsEmit, EventsOn } from '../../../wailsjs/runtime';
import { Page } from "../../components/Page/Page";
import { downloadModel } from "../../lib/models";

export const PageHome = () => {
  const [downloadState, setDownloadState] = useState<any | null>(null);
  const [model, setModel] = useState<string | null>(null);

  const [prompt, setPrompt] = useState<string | null>(null);
  const [chatHistory, setChatHistory] = useState<string[]>([]);

  const download = () => {
    if (!model) return;
    downloadModel(model, setDownloadState);
  }

  const send = () => {
    if (!prompt) return;
    EventsEmit("chat", "gpt2", prompt);
    setChatHistory(chatHistory => [...chatHistory, prompt]);
    setPrompt("");
  }

  useEffect(() => {
    EventsOn("chat-response", ({ response }: { response: string }) => {
      setChatHistory(chatHistory => [...chatHistory, response]);
    });
  }, []);

  return <Page>
    <h1>Home</h1>

    <hr />

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
