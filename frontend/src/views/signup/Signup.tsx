import React from 'react';
import { useNavigate } from 'react-router-dom';
import AppButton from '../../components/buttons/AppButton';
import useAuthStore from '../../stores/auth';
import { SubmitHandler, useForm } from 'react-hook-form';

import { NavLink } from 'react-router-dom';
import useToastStore from '../../stores/toast';
import AppInput from '../../components/inputs/AppInput';
import AppInputAlert from '../../components/alerts/AppInputAlert';
import AppSelect from '../../components/selects/AppSelect';

enum RoleEnum {
  user = 'user',
  male = 'admin',
}

interface Inputs {
  email: string;
  role: RoleEnum;
  password: string;
}

const Signup = () => {
  const { signup } = useAuthStore();
  const navigate = useNavigate();

  const { addToastError, addToastInfo } = useToastStore();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    try {
      await signup(data);
      navigate('/login');
      addToastInfo('you successfully signed up');
    } catch (error) {
      if (error instanceof Error) {
        addToastError(error.message);
      }
    }
  };

  return (
    <div className="app__signup_wrapper">
      <div className="app__signup_container">
        <h1 className="app__signup_headline">Signup</h1>

        <form className="form" onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-5">
            <label htmlFor="email" className="app__signup_lable">
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
            <label htmlFor="email" className="app__signup_lable">
              Your role
            </label>
            <AppSelect {...register('role', { required: true })}>
              <option value="user">user</option>
              <option value="admin">admin</option>
            </AppSelect>
          </div>
          <div className="mb-5">
            <label htmlFor="password" className="app__signup_lable">
              Your password
            </label>
            <AppInput
              {...register('password', { required: true })}
              aria-invalid={errors.password ? 'true' : 'false'}
              type="password"
              className="input"
            />
            {errors.password && <AppInputAlert />}
          </div>

          <AppButton type="submit">Create Account</AppButton>
          <br />
          <NavLink to={'/login'} className="block text-center">
            I have an account
          </NavLink>
        </form>
      </div>
    </div>
  );
};

export default Signup;
