import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import User from '../../components/user/User';
import useUserStore from '../../stores/user';
import AppButton from '../../components/buttons/AppButton';
import useAuthStore from '../../stores/auth';

const Home = () => {
  const { user } = useAuthStore();
  const isAdmin = user.role === 'admin';
  const { users, getAllUsers } = useUserStore((state) => state);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const users = await getAllUsers();
        useUserStore.setState({ users });
      } catch (error) {
        if (error instanceof Error) {
          console.log(error.message);
        }
      }
    };
    fetchUsers();
  }, []);

  return (
    <div className="home main">
      <div className="main__content">
        <div className="flex justify-center flex-col items-center gap-3 mb-20">
          <h1 className="text-center text-gray-900 text-xl font-bold mt-20">
            Users
          </h1>
          <Link
            to={isAdmin ? '/create' : '/'}
            className={isAdmin ? '' : 'opacity-50'}
          >
            <AppButton>Create User</AppButton>
          </Link>
        </div>

        {users.map(({ id, email, role }) => (
          <User key={id} id={id} email={email} role={role} />
        ))}
      </div>
    </div>
  );
};

export default Home;
