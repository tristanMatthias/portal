import { useState } from "react";
import { ActionChat } from "../../../wailsjs/go/chat/Chat";
import Page from "../../components/Page/Page";
import './Home.page.scss';

export default function PageHome() {

  const [prompt, setPrompt] = useState<string | null>(null);
  const [chatHistory, setChatHistory] = useState<string[]>([]);


  const send = () => {
    if (!prompt) return;
    ActionChat("gpt2", prompt).then(r => {
      setChatHistory(chatHistory => [...chatHistory, r.response]);
    });
    setChatHistory(chatHistory => [...chatHistory, prompt]);
    setPrompt("");
  }

  return <Page id="home" title="Home" >

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
