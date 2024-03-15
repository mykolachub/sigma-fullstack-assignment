import React from 'react';

interface AppButtonProps extends React.ComponentPropsWithRef<'button'> {
  onClick?: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  children?: React.ReactNode;
}

const AppButtonSmall = (props: AppButtonProps) => {
  const { onClick, children } = props;
  return (
    <button
      onClick={onClick}
      className="button text-gray-900 hover:text-white border border-gray-800 hover:bg-gray-900 focus:ring-4 focus:outline-none focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:border-gray-600 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-gray-800"
    >
      {children}
    </button>
  );
};

export default AppButtonSmall;
