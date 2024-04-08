import { Outlet, useNavigate } from 'react-router-dom';
import useAuthStore from '../stores/auth';
import Toast from './Toast';
import React from 'react';

const Layout = () => {
  const { user, logout } = useAuthStore();
  console.dir({ user }, { depth: null });
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <main className="app__wrapper h-screen">
      {user && (
        <div className='className="fixed right-1/2 translate-x-1/2"'>
          <p>
            Welcome {user.email}! <i>{user.role}</i>
          </p>
          <button onClick={handleLogout}>Logout</button>
        </div>
      )}

      <Outlet />
      <Toast />
    </main>
  );
};

export default Layout;
