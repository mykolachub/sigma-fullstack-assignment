import useUserStore from '../../stores/user';
import React, { useState } from 'react';
import AppButton from '../../components/buttons/AppButton';
import { useForm, SubmitHandler } from 'react-hook-form';
import useToastStore from '../../stores/toast';
import AppInput from '../../components/inputs/AppInput';
import AppSelect from '../../components/selects/AppSelect';
import AppInputAlert from '../../components/alerts/AppInputAlert';

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
    <div className="app__signup_wrapper">
      <div className="app__signup_container">
        <h1 className="app__signup_headline">Update User</h1>

        <form className="app__signup_form" onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-5">
            <label htmlFor="email" className="lable">
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

          <AppButton type="submit">Update User</AppButton>
        </form>
      </div>
    </div>
  );
};

export default UpdateUser;
