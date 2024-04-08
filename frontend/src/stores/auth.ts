import axios from 'axios';
import { create } from 'zustand';
import config from '../utils/config';

import { UserDTO, UserCreateDTO, UserLoginDTO } from '../types/user';
import { handleAxiosError } from '../utils/errors/axios';
import {
  SignupResponseData,
  LoginResponseData,
  ServerResponse,
  GetMeResponseData,
} from '../types/server';

const API_URL = config.env.apiUrl;

interface UserAuthState {
  user: UserDTO;
  authed: boolean;
  token: string | null;
  signup: (body: UserCreateDTO) => Promise<SignupResponseData>;
  login: (body: UserLoginDTO) => Promise<LoginResponseData>;
  setAuthorization: () => void;
  setUser: (data: UserDTO) => void;
  setToken: (token: string) => void;
  logout: () => void;
  me: () => Promise<UserDTO>;
}

const useAuthStore = create<UserAuthState>((set) => ({
  user: <UserDTO>{},
  authed: false,
  token: null,
  signup: async (body: UserCreateDTO): Promise<SignupResponseData> => {
    try {
      const response = await axios.post(API_URL + '/user/signup', body);
      const { data } = response.data as ServerResponse<SignupResponseData>;
      return data;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  login: async (body: UserLoginDTO): Promise<LoginResponseData> => {
    try {
      const response = await axios.post(API_URL + '/user/login', body);
      const { data } = response.data as ServerResponse<LoginResponseData>;
      localStorage.setItem('access_token', data.token);
      // TODO: get refresh token, set to localStorage
      return data;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
  setAuthorization: (): void => {
    const localToken = localStorage.getItem('access_token');
    if (!localToken) {
      set({
        user: <UserDTO>{},
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
    set({ user: <UserDTO>{}, authed: false, token: '' });
  },
  me: async (): Promise<UserDTO> => {
    try {
      const token = `Bearer ${useAuthStore.getState().token}`;
      const response = await axios.get(API_URL + '/users/me', {
        headers: { Authorization: token },
      });
      const { data } = response.data as ServerResponse<GetMeResponseData>;
      return data.user;
    } catch (error) {
      throw new Error(handleAxiosError(error));
    }
  },
}));

export default useAuthStore;
