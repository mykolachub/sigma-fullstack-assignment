import React from 'react';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';
import AppButton from '../../components/buttons/AppButton';
import useAuthStore from '../../stores/auth';
import { SubmitHandler, useForm } from 'react-hook-form';

import useToastStore from '../../stores/toast';
import AppInput from '../../components/inputs/AppInput';
import AppInputAlert from '../../components/alerts/AppInputAlert';

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
    <div className="app__login_wrapper">
      <div className="app__login_container">
        <h1 className="app__login_headline">Login</h1>

        <form className="app__login_form" onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-5">
            <label htmlFor="email" className="app__login_lable">
              Your email
            </label>
            <AppInput
              {...register('email', { required: true })}
              aria-invalid={errors.email ? 'true' : 'false'}
              type="text"
              placeholder="name@domain.com"
            />
            {errors.email && <AppInputAlert />}
          </div>
          <div className="mb-5">
            <label htmlFor="password" className="app__login_label">
              Your password
            </label>
            <AppInput
              {...register('password', { required: true })}
              aria-invalid={errors.password ? 'true' : 'false'}
              type="password"
            />
            {errors.password && <AppInputAlert />}
          </div>

          <AppButton type="submit">Log in to Account</AppButton>
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
