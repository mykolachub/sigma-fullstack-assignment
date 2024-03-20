import React, { PropsWithChildren, useContext, useEffect } from 'react';
import { createContext, useState } from 'react';

import { UserDTO } from '../types/user';
import useAuthStore from '../stores/auth';
import axios, { AxiosResponse } from 'axios';

type AuthProviderProps = PropsWithChildren;

const AuthContext = createContext<Promise<AxiosResponse<ServerData>> | null>(
  null,
);

interface ServerData {
  data: UserDTO;
}

export default function AuthProvider({ children }: AuthProviderProps) {
  const { authed, setToken } = useAuthStore();

  const localToken = window.localStorage.getItem('access_token');
  const [userPromise, setUserPromise] = useState<Promise<
    AxiosResponse<ServerData>
  > | null>(() => {
    if (!localToken) return null;
    setToken(localToken);
    return axios.get<ServerData>('http://localhost:8080/api/me', {
      headers: { Authorization: `Bearer ${localToken}` },
    });
  });

  useEffect(() => {
    // authed has been changed in setAuthorization after login
    if (!localToken) {
      setUserPromise(null);
      return;
    }
    setToken(localToken);
    setUserPromise(
      axios.get<ServerData>('http://localhost:8080/api/me', {
        headers: { Authorization: `Bearer ${localToken}` },
      }),
    );
  }, [authed]);

  return (
    <AuthContext.Provider value={userPromise}>{children}</AuthContext.Provider>
  );
}

export const useAuth = () => {
  const context = useContext(AuthContext);

  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }

  return context;
};
