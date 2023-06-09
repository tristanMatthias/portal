import { useNavigate } from "react-router-dom";
import './Sidebar.scss';

export default function Sidebar() {
  const navigate = useNavigate();

  return <aside>
    <nav>
      <button>Resume</button>
      <button>Conversations</button>
      <button onClick={() => navigate('/models')}>Models</button>
      <button>Search</button>
      <button onClick={() => navigate('/settings')}>Settings</button>
    </nav>
  </aside>;
}
