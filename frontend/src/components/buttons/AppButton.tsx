import React from 'react';
import './AppButtons.css';

interface AppButtonProps extends React.ComponentPropsWithRef<'button'> {
  onClick?: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  children?: React.ReactNode;
}

const AppButton = (props: AppButtonProps) => {
  const { onClick, children } = props;
  return (
    <button onClick={onClick} className="button">
      {children}
    </button>
  );
};

export default AppButton;
