import { createRoot } from 'react-dom/client';
import './index.css';
import { createBrowserRouter, RouterProvider, NavLink } from 'react-router-dom';
import Dashboard from './pages/dashboard';
import Login from './pages/login';

const router = createBrowserRouter([
       {
              path: '/',
              element: <Dashboard />
       },
       {
              path: '/login',
              element: <Login/>
       }
])

let root = createRoot(document.getElementById('root') as Element);
root.render(<RouterProvider router={router} />);

export default router;