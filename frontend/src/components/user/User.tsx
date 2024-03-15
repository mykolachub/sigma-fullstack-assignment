import React from 'react';
import useUserStore from '../../stores/user';
import { useNavigate } from 'react-router-dom';
import AppButtonSmall from '../buttons/AppButtonSmall';

interface UserProps {
  email: string;
  id: string;
}

const User = ({ email, id }: UserProps) => {
  const { deleteUser, getAllUsers } = useUserStore();
  const navigate = useNavigate();

  const handleButtonSubmit = () => {
    deleteUser(id).then(() => getAllUsers());
  };

  const handleUpdateUser = () => {
    navigate(`/update?id=${id}`);
  };

  return (
    <div className="user__wrapper flex w-full justify-between items-center cursor-pointer rounded-lg bg-gray-100 p-2 m-1">
      <div onClick={handleUpdateUser} className="flex-1">
        <h3>{email}</h3>
      </div>
      <AppButtonSmall onClick={handleButtonSubmit}>Delete</AppButtonSmall>
    </div>
  );
};

export default User;
