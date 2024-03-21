import React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import AppButton from '../../components/buttons/AppButton';
import useUserStore from '../../stores/user';
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
const CreateUser = () => {
  const { createUser } = useUserStore();
  const { addToastError, addToastInfo } = useToastStore();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    try {
      await createUser(data);
      addToastInfo('user created');
    } catch (error) {
      if (error instanceof Error) {
        addToastError(error.message);
      }
    }
  };

  return (
    <div className="app__signup_wrapper">
      <div className="app__signup_container">
        <h1 className="app__signup_headline">Create New User</h1>

        <form className="app__signup_form" onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-5">
            <label htmlFor="email" className="app__signup_label">
              User email
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
            <label htmlFor="email" className="app__signup_label">
              User role
            </label>
            <AppSelect {...register('role', { required: true })}>
              <option value="user">user</option>
              <option value="admin">admin</option>
            </AppSelect>
          </div>
          <div className="mb-5">
            <label htmlFor="password" className="app__signup_label">
              User password
            </label>
            <AppInput
              {...register('password', { required: true })}
              aria-invalid={errors.password ? 'true' : 'false'}
              type="password"
              className="input"
            />
            {errors.password && <AppInputAlert />}
          </div>

          <AppButton type="submit">Create New User</AppButton>
        </form>
      </div>
    </div>
  );
};

export default CreateUser;
