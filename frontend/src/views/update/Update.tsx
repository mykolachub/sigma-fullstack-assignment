import useUserStore from '../../stores/user';
import React, { useState } from 'react';
import AppButton from '../../components/buttons/AppButton';
import { useForm, SubmitHandler } from 'react-hook-form';
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

const UpdateUser = () => {
  const { getUserById, updateUser } = useUserStore();
  const { addToastError, addToastInfo } = useToastStore();

  const [id] = useState<string>(() => {
    const queryParameters = new URLSearchParams(window.location.search);
    return queryParameters.get('id') || '';
  });

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Inputs>({
    defaultValues: async () => {
      const user = await getUserById(id);
      return user as Inputs;
    },
  });
  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    try {
      await updateUser(id, data);
      addToastInfo('user updated');
    } catch (error) {
      if (error instanceof Error) {
        addToastError(error.message);
      }
    }
  };

  return (
    <div className="page__wrapper">
      <div className="container">
        <h1 className="headline">Update User</h1>

        <form className="form" onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-5">
            <label htmlFor="email" className="lable">
              User email
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
              User role
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
              User password
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

          <AppButton type="submit">Update User</AppButton>
        </form>
      </div>
    </div>
  );
};

export default UpdateUser;
