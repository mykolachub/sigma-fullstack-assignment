import { Outlet, useNavigate } from 'react-router-dom';
import useAuthStore from '../stores/auth';

const Layout = () => {
  const { user, logout } = useAuthStore();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <main className="app__wrapper h-screen">
      {user.email && (
        <div className='className="fixed right-1/2 translate-x-1/2"'>
          <p>
            Welcome {user.email}! <i>{user.role}</i>
          </p>
          <button onClick={handleLogout}>Logout</button>
        </div>
      )}

      <Outlet />
    </main>
  );
};

export default Layout;