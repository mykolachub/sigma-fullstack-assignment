import axios, { AxiosResponse } from 'axios';
import { create } from 'zustand';
import config from '../utils/config';

import { UserDTO, UserCreateDTO } from '../types/user';

const API_URL = config.env.apiKey;

interface UserStoreState {
  users: UserDTO[];
  getAllUsers: () => Promise<void>;
  getUserById: (id: string) => Promise<UserDTO>;
  createUser: (data: UserCreateDTO) => Promise<UserDTO>;
  updateUser: (id: string, data: UserCreateDTO) => Promise<UserDTO>;
  deleteUser: (id: string) => Promise<AxiosResponse>;
}

const useUserStore = create<UserStoreState>((set) => ({
  users: [],
  getAllUsers: async () => {
    const response = await axios.get(API_URL + '/users');
    set({ users: await response.data.data });
  },
  getUserById: async (id: string): Promise<UserDTO> => {
    const response = await axios.get(API_URL + `/user?id=${id}`);
    return response.data.data as UserDTO;
  },
  createUser: async (data: UserCreateDTO): Promise<UserDTO> => {
    const response = await axios.post(API_URL + '/users', data);
    return response.data.data as UserDTO;
  },
  updateUser: async (id: string, data: UserCreateDTO): Promise<UserDTO> => {
    const response = await axios.patch(API_URL + `/users?id=${id}`, data);
    return response.data.data as UserDTO;
  },
  deleteUser: async (id: string): Promise<AxiosResponse> => {
    const response = await axios.delete(API_URL + `/users?id=${id}`);
    return response;
  },
}));

export default useUserStore;
