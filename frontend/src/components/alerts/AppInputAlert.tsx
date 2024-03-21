import React from 'react';

import './AppAlerts.css';

interface AppInputAlertProps extends React.ComponentPropsWithRef<'span'> {
  message?: string;
}

const AppInputAlert = (props: AppInputAlertProps) => {
  const { message } = props;
  return (
    <span role="alert" className="app_input_alert">
      {message}
    </span>
  );
};

export default AppInputAlert;
