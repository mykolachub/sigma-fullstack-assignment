import React from 'react';
import ButtonSmall from '../buttons/ButtonSmall';
import useUserStore from '../../utils/store';
import { useNavigate } from 'react-router-dom';

const User = ({ email, id }) => {
  const { deleteUser, getAllUsers } = useUserStore();
  const navigate = useNavigate();

  const handleButtonSubmit = (event) => {
    deleteUser(id).then((res) => getAllUsers());
  };

  const handleUpdateUser = () => {
    navigate(`/update?id=${id}`);
  };

  return (
    <div className="user__wrapper flex w-full justify-between items-center cursor-pointer rounded-lg bg-gray-100 p-2 m-1">
      <div onClick={handleUpdateUser} className="flex-1">
        <h3>{email}</h3>
      </div>
      <ButtonSmall onClick={handleButtonSubmit}>Delete</ButtonSmall>
    </div>
  );
};

export default User;
