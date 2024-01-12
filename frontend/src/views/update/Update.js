import { useEffect, useState } from 'react';
import useUserStore from '../../utils/store';
import { useNavigate } from 'react-router-dom';
import Button from '../../components/buttons/Button';

const UpdateUser = () => {
  const { getUserById, updateUser } = useUserStore();
  const navigate = useNavigate();

  const [id, setId] = useState('');
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  useEffect(() => {
    const queryParameters = new URLSearchParams(window.location.search);
    const id = queryParameters.get('id');
    setId(id);
    getUserById(id).then(({ email, password }) => {
      setEmail(email);
      setUsername(email);
      setPassword(password);
    });
  }, []);

  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleButtonSubmit = (event) => {
    if (!email.length || !password.length) return;
    updateUser(id, { email, password }).then((res) => navigate('/'));
  };

  return (
    <div className="home main h-full">
      <div className="main__content flex justify-center flex-col items-center h-full gap-10">
        <h1 className="text-center text-gray-900 text-xl font-bold">
          Change {username} User
        </h1>

        <form class="max-w-sm mx-auto">
          <div class="mb-5">
            <label
              for="email"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Change your email
            </label>
            <input
              type="email"
              id="email"
              name="email"
              onChange={handleEmailChange}
              value={email}
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="name@flowbite.com"
              required
            />
          </div>
          <div class="mb-5">
            <label
              for="password"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Change your password
            </label>
            <input
              type="text"
              id="password"
              name="passwprd"
              onChange={handlePasswordChange}
              value={password}
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              required
            />
          </div>

          <Button className="" onClick={handleButtonSubmit}>
            Change User
          </Button>
        </form>
      </div>
    </div>
  );
};

export default UpdateUser;
