import { useState } from 'react';
import useUserStore from '../../utils/store';
import { useNavigate } from 'react-router-dom';
import Button from '../../components/buttons/Button';

const CreateUser = () => {
  const { createUser } = useUserStore();
  const navigate = useNavigate();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleButtonSubmit = (event) => {
    if (!email.length || !password.length) return;
    createUser({ email, password }).then(() => navigate('/'));
  };

  return (
    <div className="home main h-full">
      <div className="main__content flex justify-center flex-col items-center h-full gap-10">
        <h1 className="text-center text-gray-900 text-xl font-bold">
          Create User
        </h1>

        <form className="max-w-sm mx-auto">
          <div className="mb-5">
            <label
              htmlFor="email"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Your email
            </label>
            <input
              type="email"
              id="email"
              name="email"
              onChange={handleEmailChange}
              value={email}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="name@flowbite.com"
              required
            />
          </div>
          <div className="mb-5">
            <label
              htmlFor="password"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Your password
            </label>
            <input
              type="text"
              id="password"
              name="passwprd"
              onChange={handlePasswordChange}
              value={password}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              required
            />
          </div>

          <Button onClick={handleButtonSubmit}>Create New User</Button>
        </form>
      </div>
    </div>
  );
};

export default CreateUser;
