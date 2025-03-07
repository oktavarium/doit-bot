import TextField from '@mui/material/TextField';
import { useEffect } from 'react';
import { on } from '@telegram-apps/sdk';
import './BasicTextFields.scss';

export default function BasicTextFields({ label, value, onChange }) {
  useEffect(() => {
    let isKeyboardOpen = false;

    // Слушаем изменения viewport только при закрытии клавиатуры
    const removeViewportListener = on('viewport_changed', (event) => {
      if (isKeyboardOpen && event.isStateStable) {
        isKeyboardOpen = false;
        window.scrollTo(0, 0);
      }
    });

    // Обрабатываем фокус на инпуте
    const handleFocus = () => {
      isKeyboardOpen = true;
    };

    // Добавляем обработчик фокуса
    const input = document.querySelector('input');
    if (input) {
      input.addEventListener('focus', handleFocus);
    }

    return () => {
      removeViewportListener();
      if (input) {
        input.removeEventListener('focus', handleFocus);
      }
    };
  }, []);

  return (
    <TextField
      className="text-field"
      label={label}
      variant="outlined"
      value={value}
      onChange={onChange}
      size="small"
      fullWidth
      autoComplete="off"
    />
  );
}