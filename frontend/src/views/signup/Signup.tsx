import React from 'react';
import { useNavigate } from 'react-router-dom';
import AppButton from '../../components/buttons/AppButton';
import useAuthStore from '../../stores/auth';
import { SubmitHandler, useForm } from 'react-hook-form';

import './Signup.css';
import { NavLink } from 'react-router-dom';
import useToastStore from '../../stores/toast';

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
    <div className="page__wrapper">
      <div className="container">
        <h1 className="headline">Signup</h1>

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
            <label htmlFor="email" className="lable">
              Your role
            </label>
            <select
              {...register('role', { required: true })}
              className="select"
            >
              <option value="user">user</option>
              <option value="admin">admin</option>
            </select>
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
