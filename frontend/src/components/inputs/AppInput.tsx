import React, { forwardRef } from 'react';
import './AppInput.css';

interface AppInputProps extends React.ComponentPropsWithRef<'input'> {
  type: React.HTMLInputTypeAttribute;
  placeholder?: string;
}

const AppInput = forwardRef<HTMLInputElement, AppInputProps>(
  (props: AppInputProps, ref) => {
    const { placeholder, type, ...rest } = props;
    return (
      <input
        {...rest}
        ref={ref}
        type={type}
        className="input"
        placeholder={placeholder}
      />
    );
  },
);

AppInput.displayName = 'AppInput';

export default AppInput;
