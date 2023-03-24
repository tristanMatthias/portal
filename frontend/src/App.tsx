import React from 'react';
import { RouterProvider } from 'react-router-dom';
import router from './router';
import './styles/base.scss';

export const App = () => {
  return <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
}
