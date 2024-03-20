import React from 'react';
import useUserStore from '../../stores/user';
import { useNavigate } from 'react-router-dom';
import AppButtonSmall from '../buttons/AppButtonSmall';
import useAuthStore from '../../stores/auth';
import useToastStore from '../../stores/toast';

interface UserProps {
  email: string;
  id: string;
  role: string;
}

const User = ({ email, id, role }: UserProps) => {
  const { user, logout } = useAuthStore();
  const isAdmin = user.role === 'admin';
  const isOwner = user.id === id;
  const canDelete = isAdmin || isOwner;

  const { deleteUser, getAllUsers } = useUserStore();
  const navigate = useNavigate();

  const { addToastError, addToastInfo } = useToastStore();

  const handleDeleteUser = async () => {
    if (!isAdmin && !isOwner) {
      addToastError('you do not have permission');
      return;
    }
    try {
      await deleteUser(id);
      const users = await getAllUsers();
      useUserStore.setState({ users: users });
      addToastInfo('user deleted');
      if (isOwner) {
        logout();
        navigate('/login');
      }
    } catch (error) {
      if (error instanceof Error) {
        addToastError(error.message);
      }
    }
  };

  const handleUpdateUser = () => {
    if (!isAdmin && !isOwner) {
      addToastError('you do not have permission');
      return;
    }
    navigate(`/update?id=${id}`);
  };

  return (
    <div
      className="user__wrapper flex w-full justify-between items-center cursor-pointer rounded-lg bg-gray-100 p-2 m-1"
      style={{ minHeight: '70px', opacity: isAdmin || isOwner ? '1' : '0.5' }}
    >
      <div onClick={handleUpdateUser} className="flex-1">
        <h3>
          {email} <i>{role}</i> {isAdmin}
        </h3>
      </div>
      {canDelete && (
        <AppButtonSmall onClick={handleDeleteUser}>Delete</AppButtonSmall>
      )}
    </div>
  );
};

export default User;
