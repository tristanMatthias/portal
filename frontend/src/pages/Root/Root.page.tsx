import { PropsWithChildren } from "react";
import { Outlet } from "react-router-dom";
import './Root.page.scss';

export const PageRoot = ({ children }: PropsWithChildren) => <div id="page">
  <Outlet />
</div>
