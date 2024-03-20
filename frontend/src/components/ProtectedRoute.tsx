import { useEffect, useState } from 'react';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthProvider';
import React from 'react';
import useAuthStore from '../stores/auth';

const ProtectedRoute = () => {
  const { setUser, setAuthorization } = useAuthStore();
  const userPromise = useAuth();
  const navigate = useNavigate();
  const location = useLocation();

  const [authenticated, setAuthenticated] = useState(false);

  useEffect(() => {
    if (!userPromise) {
      setAuthenticated(false);
      // After login user gets redirected to the previous page
      navigate('/login', { replace: true, state: { from: location } });
      return;
    }
    userPromise
      .then((res) => {
        console.log(res.data);
        const { data: user } = res.data;
        setAuthenticated(user !== null);

        setUser(user);
        setAuthorization();
      })
      .catch((err: unknown) => console.log(err));
  }, [userPromise]);

  return authenticated !== null ? <Outlet /> : null;
};

export default ProtectedRoute;
