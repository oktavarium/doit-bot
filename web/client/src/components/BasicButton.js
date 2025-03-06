import * as React from 'react';
// import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';

export default function BasicButton({ onClick, disabled, style }) {
  return (
    <Button
      variant="contained"
      onClick={onClick}
      disabled={disabled}
      fullWidth
      sx={{
        height: '40px',
        padding: '0 10px',
        fontSize: '14px',
        ...style
      }}
    >
      Жми сюда
    </Button>
  );
}