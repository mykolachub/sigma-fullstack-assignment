import axios from 'axios';
import { create } from 'zustand';
import config from '../utils/config';

import {
  UserDTO,
  UserCreateDTO,
  UserLoginDTO,
  UserTokenDTO,
} from '../types/user';
import { handleAxiosError } from '../utils/errors/axios';

const API_URL = config.env.apiKey;

interface UserAuthState {
  user: UserDTO;
  admin: boolean;
  authed: boolean;
  token: string | null;
  signup: (data: UserCreateDTO) => Promise<UserDTO>;
  login: (data: UserLoginDTO) => Promise<UserTokenDTO>;
  setAuthorization: () => void;
  setUser: (data: UserDTO) => void;
  setToken: (token: string) => void;
  logout: () => void;
  me: () => void;
}

const useAuthStore = create<UserAuthState>((set) => ({
  user: <UserDTO>{},
  admin: false,
  authed: false,
  token: null,
  signup: async (data: UserCreateDTO): Promise<UserDTO> => {
    try {
      const response = await axios.post(API_URL + '/signup', data);
      return response.data.data as UserDTO;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  login: async (data: UserLoginDTO): Promise<UserTokenDTO> => {
    try {
      const response = await axios.post(API_URL + '/login', data);
      const { token } = response.data;
      localStorage.setItem('access_token', token);
      // TODO: get refresh token, set to localStorage
      return { token } as UserTokenDTO;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  setAuthorization: (): void => {
    const localToken = localStorage.getItem('access_token');
    if (!localToken) {
      set({
        user: <UserDTO>{},
        admin: false,
        authed: false,
        token: null,
      });
      return;
    }
    set({ token: localToken, authed: true });
  },
  setUser: (data: UserDTO): void => set({ user: data, authed: true }),
  setToken: (token: string): void => set({ token: token }),
  logout: (): void => {
    localStorage.removeItem('access_token');
    set({ user: <UserDTO>{}, admin: false, authed: false, token: '' });
  },
  me: async () => {},
}));

export default useAuthStore;
