// A React router v6 component that renders the appropriate component based on the current URL

import { createBrowserRouter } from 'react-router-dom';
import PageError from './pages/Error/Error.page';
import PageHome from './pages/Home/Home.page';
import PageModels from './pages/Models/Models.page';
import PageNotFound from './pages/NotFound/NotFound.page';
import { PageRoot } from './pages/Root/Root.page';


export default createBrowserRouter([
  {
    path: "/",
    element: <PageRoot />,
    errorElement: <PageError />,
    children: [
      {
        path: "",
        element: <PageHome />,
      },
      {
        path: "/models",
        element: <PageModels />,
      },
      {
        path: "*",
        element: <PageNotFound />,
      },
    ]
  }
]);
