import TextField from '@mui/material/TextField';
import { useEffect } from 'react';
import { on } from '@telegram-apps/sdk';

export default function BasicTextFields({ label, value, onChange, style }) {
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
      label={label}
      variant="outlined"
      value={value}
      onChange={onChange}
      size="small"
      fullWidth
      autoComplete="off"
      sx={{
        '& .MuiOutlinedInput-root': {
          height: '40px',
        },
        '& .MuiOutlinedInput-input': {
          height: '40px',
          padding: '0 14px',
          '-webkit-appearance': 'none',
          '-webkit-tap-highlight-color': 'transparent'
        },
        '& .MuiInputLabel-root': {
          transform: 'translate(14px, -6px) scale(0.75)',
          backgroundColor: '#fff',
          padding: '0 4px'
        },
        ...style
      }}
    />
  );
}