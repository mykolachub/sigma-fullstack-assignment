import React from 'react';

interface AppButtonProps extends React.ComponentPropsWithRef<'button'> {
  onClick?: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  children?: React.ReactNode;
}

const AppButton = (props: AppButtonProps) => {
  const { onClick, children } = props;
  return (
    <button
      onClick={onClick}
      className="button w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5  dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
    >
      {children}
    </button>
  );
};

export default AppButton;
