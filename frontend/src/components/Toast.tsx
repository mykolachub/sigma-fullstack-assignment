import React from 'react';
import useToastStore, { ToastType } from '../stores/toast';

const Toast = () => {
  const { toasts, counter } = useToastStore();

  return (
    <>
      {counter > 0 && (
        <div className="fixed bottom-5 right-1/2 translate-x-1/2">
          {toasts.map(({ type, message }, idx) => {
            const toastColors = {
              [ToastType.info]: 'blue',
              [ToastType.warning]: 'orange',
              [ToastType.error]: 'red',
            };
            const toastColor = type ? toastColors[type] : 'black';

            return (
              <p key={`${message}${idx}`}>
                <span style={{ color: toastColor }}>{type}: </span>
                {message}
              </p>
            );
          })}
        </div>
      )}
    </>
  );
};

export default Toast;
