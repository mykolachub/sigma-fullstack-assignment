import React from 'react';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';
import AppButton from '../../components/buttons/AppButton';
import useAuthStore from '../../stores/auth';
import { SubmitHandler, useForm } from 'react-hook-form';

import './Login.css';
import useToastStore from '../../stores/toast';

interface Inputs {
  email: string;
  password: string;
}

const Login = () => {
  const { login, setAuthorization } = useAuthStore();
  const { addToastError, addToastInfo } = useToastStore();

  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from?.pathname || '/';

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    try {
      await login(data);
      setAuthorization();
      navigate(from, { replace: true });
      addToastInfo('you successfully logged in');
    } catch (error) {
      if (error instanceof Error) {
        addToastError(error.message);
      }
    }
  };

  return (
    <div className="page__wrapper">
      <div className="container">
        <h1 className="headline">Login</h1>

        <form className="form" onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-5">
            <label htmlFor="email" className="lable">
              Your email
            </label>
            <input
              {...register('email', { required: true })}
              aria-invalid={errors.email ? 'true' : 'false'}
              type="text"
              className="input"
              placeholder="name@domain.com"
            />
            {errors.email && (
              <span role="alert" className="input__alert">
                This field is required
              </span>
            )}
          </div>
          <div className="mb-5">
            <label htmlFor="password" className="lable">
              Your password
            </label>
            <input
              {...register('password', { required: true })}
              aria-invalid={errors.password ? 'true' : 'false'}
              type="password"
              className="input"
            />
            {errors.password && (
              <span role="alert" className="input__alert">
                This field is required
              </span>
            )}
          </div>

          <AppButton type="submit" className="">
            Log in to Account
          </AppButton>
          <br />
          <NavLink to={'/signup'} className="block text-center">
            Create new account
          </NavLink>
        </form>
      </div>
    </div>
  );
};

export default Login;
