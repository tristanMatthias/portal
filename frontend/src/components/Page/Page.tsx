import { useNavigate } from "react-router-dom";
import Sidebar from "../Sidebar/Sidebar";
import "./Page.scss";

export interface PageProps {
  id: string;
  title: string;
  sidebar?: boolean;
  backButton?: {
    to: string;
    text?: string;
  }
  children: React.ReactNode;
}
export default function Page({
  sidebar = true,
  children,
  id,
  title,
  backButton,
}: PageProps) {
  const navigate = useNavigate();

  return <>
    {sidebar && <Sidebar />}
    <main id={id}>
      <header>
        {backButton && <button onClick={() => navigate(backButton.to)}>
          {backButton.text || "Back"}
        </button>}
        <h1>{title}</h1>
      </header>

      <section>
        {children}
      </section>
    </main>
  </>;
}
