import React from 'react';
import './AppButtons.css';

interface AppButtonProps extends React.ComponentPropsWithRef<'button'> {
  onClick?: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  children?: React.ReactNode;
}

const AppButtonSmall = (props: AppButtonProps) => {
  const { onClick, children } = props;
  return (
    <button onClick={onClick} className="button--small">
      {children}
    </button>
  );
};

export default AppButtonSmall;
