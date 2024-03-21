import React, { forwardRef } from 'react';
import './AppSelect.css';

interface AppSelectProps extends React.ComponentPropsWithRef<'select'> {
  children: React.ReactNode;
}

const AppSelect = forwardRef<HTMLSelectElement, AppSelectProps>(
  (props: AppSelectProps, ref) => {
    const { children, ...rest } = props;
    return (
      <select {...rest} ref={ref} className="select">
        {children}
      </select>
    );
  },
);

AppSelect.displayName = 'AppSelect';

export default AppSelect;
