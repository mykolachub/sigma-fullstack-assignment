import { Link } from 'react-router-dom';
import User from '../../components/user/User';
import useUserStore from '../../utils/store';
import { useEffect } from 'react';
import Button from '../../components/buttons/Button';

const Home = () => {
  const { users, getAllUsers } = useUserStore();

  useEffect(() => {
    getAllUsers();
  }, []);

  return (
    <div className="home main">
      <div className="main__content">
        <div className="flex justify-center flex-col items-center gap-3 mb-20">
          <h1 className="text-center text-gray-900 text-xl font-bold mt-20">
            Users
          </h1>
          <Link to={'/create'}>
            <Button className="main__create">Create User</Button>
          </Link>
        </div>

        {users.map(({ id, email }) => (
          <User key={id} id={id} email={email} />
        ))}
      </div>
    </div>
  );
};

export default Home;
