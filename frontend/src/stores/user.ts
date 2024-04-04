import axios, { AxiosResponse } from 'axios';
import { create } from 'zustand';
import config from '../utils/config';

import { UserDTO, UserCreateDTO } from '../types/user';
import { handleAxiosError } from '../utils/errors/axios';
import useAuthStore from './auth';

const API_URL = config.env.apiUrl;

interface UserStoreState {
  users: UserDTO[];
  getAllUsers: (search?: string) => Promise<UserDTO[]>;
  getUserById: (id: string) => Promise<UserDTO>;
  createUser: (data: UserCreateDTO) => Promise<UserDTO>;
  updateUser: (id: string, data: UserCreateDTO) => Promise<UserDTO>;
  deleteUser: (id: string) => Promise<AxiosResponse>;
}

const useUserStore = create<UserStoreState>(() => ({
  users: [],
  getAllUsers: async (search: string = ''): Promise<UserDTO[]> => {
    try {
      const token = useAuthStore.getState().token;
      const response = await axios.get(
        API_URL + '/users' + `?search=${search}`,
        {
          headers: { Authorization: `Bearer ${token}` },
        },
      );
      return response.data.data as UserDTO[];
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  getUserById: async (id: string): Promise<UserDTO> => {
    try {
      const token = useAuthStore.getState().token;
      const response = await axios.get(API_URL + `/users/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      return response.data.data as UserDTO;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  createUser: async (data: UserCreateDTO): Promise<UserDTO> => {
    try {
      const token = useAuthStore.getState().token;
      const response = await axios.post(API_URL + '/users', data, {
        headers: { Authorization: `Bearer ${token}` },
      });
      return response.data.data as UserDTO;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  updateUser: async (id: string, data: UserCreateDTO): Promise<UserDTO> => {
    try {
      const token = useAuthStore.getState().token;
      const response = await axios.patch(API_URL + `/users/${id}`, data, {
        headers: { Authorization: `Bearer ${token}` },
      });
      return response.data.data as UserDTO;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  deleteUser: async (id: string): Promise<AxiosResponse> => {
    try {
      // TODO: if users deletes himself, force logout
      const token = useAuthStore.getState().token;
      const response = await axios.delete(API_URL + `/users/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      return response;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
}));

export default useUserStore;
