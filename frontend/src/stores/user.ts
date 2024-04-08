import axios from 'axios';
import { create } from 'zustand';
import config from '../utils/config';

import { UserDTO, UserCreateDTO } from '../types/user';
import { handleAxiosError } from '../utils/errors/axios';
import useAuthStore from './auth';
import {
  CreateUserResponseData,
  DeleteUserResponseData,
  GetUserResponseData,
  GetUsersResposeData,
  ServerResponse,
  UpdateUserResponseData,
} from '../types/server';

const API_URL = config.env.apiUrl;

interface UserStoreState {
  users: UserDTO[];
  getAllUsers: (search?: string) => Promise<GetUsersResposeData>;
  getUserById: (id: string) => Promise<GetUserResponseData>;
  createUser: (body: UserCreateDTO) => Promise<CreateUserResponseData>;
  updateUser: (
    id: string,
    body: UserCreateDTO,
  ) => Promise<UpdateUserResponseData>;
  deleteUser: (id: string) => Promise<DeleteUserResponseData>;
}

const useUserStore = create<UserStoreState>(() => ({
  users: [],
  getAllUsers: async (search: string = ''): Promise<GetUsersResposeData> => {
    try {
      const token = useAuthStore.getState().token;
      const response = await axios.get(
        API_URL + '/users' + `?search=${search}`,
        {
          headers: { Authorization: `Bearer ${token}` },
        },
      );
      const { data } = response.data as ServerResponse<GetUsersResposeData>;
      return data;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  getUserById: async (id: string): Promise<GetUserResponseData> => {
    try {
      const token = useAuthStore.getState().token;
      const response = await axios.get(API_URL + `/users/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      const { data } = response.data as ServerResponse<GetUserResponseData>;

      return data;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  createUser: async (body: UserCreateDTO): Promise<CreateUserResponseData> => {
    try {
      const token = useAuthStore.getState().token;
      const response = await axios.post(API_URL + '/users', body, {
        headers: { Authorization: `Bearer ${token}` },
      });
      const { data } = response.data as ServerResponse<CreateUserResponseData>;
      return data;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  updateUser: async (
    id: string,
    body: UserCreateDTO,
  ): Promise<UpdateUserResponseData> => {
    try {
      const token = useAuthStore.getState().token;
      const response = await axios.patch(API_URL + `/users/${id}`, body, {
        headers: { Authorization: `Bearer ${token}` },
      });
      const { data } = response.data as ServerResponse<UpdateUserResponseData>;
      return data;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  deleteUser: async (id: string): Promise<DeleteUserResponseData> => {
    try {
      // TODO: if users deletes himself, force logout
      const token = useAuthStore.getState().token;
      const response = await axios.delete(API_URL + `/users/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      const { data } = response.data as ServerResponse<DeleteUserResponseData>;
      return data;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
}));

export default useUserStore;
