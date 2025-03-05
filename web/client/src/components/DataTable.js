import * as React from 'react';
import { DataGrid } from '@mui/x-data-grid';
import Paper from '@mui/material/Paper';
import CheckIcon from '@mui/icons-material/Check';
import CloseIcon from '@mui/icons-material/Close';

const columns = [
  { 
    field: 'done', 
    headerName: 'Выполнено', 
    type: 'boolean',  
    flex: 1,
    renderCell: (params) => (
      <div style={{ 
        display: 'flex', 
        justifyContent: 'center', 
        width: '100%' 
      }}>
        {params.value ? (
          <CheckIcon style={{ color: '#4caf50' }} /> // зеленый цвет для галочки
        ) : (
          <CloseIcon style={{ color: '#f44336' }} /> // красный цвет для крестика
        )}
      </div>
    )
  },
  { 
    field: 'summary', 
    headerName: 'Задача', 
    type: 'string', 
    flex: 2,
    renderCell: (params) => (
      <div style={{ 
        textDecoration: params.row.done ? 'line-through' : 'none',
        color: params.row.done ? '#666' : 'inherit'
      }}>
        {params.value}
      </div>
    )
  },
];

export default function DataTable({ rows, selectedRows, onSelectionChange }) {
  console.log("rows", rows);
  return (
    <Paper sx={{
      margin: '0 20px',
      height: 'calc(100vh - 180px)',
      overflow: 'hidden',
      backgroundColor: 'transparent',
      boxShadow: 'none'
    }}>
      <DataGrid
        rows={rows}
        columns={columns}
        hideFooter={true}
        disableColumnMenu
        onRowSelectionModelChange={onSelectionChange}
        rowSelectionModel={selectedRows}
        sx={{
          border: 'none',
          '& .MuiDataGrid-cell:focus': {
            outline: 'none'
          },
          '& .MuiDataGrid-row:hover': {
            backgroundColor: 'rgba(25, 118, 210, 0.04)'
          },
          '& .MuiDataGrid-row.Mui-selected': {
            backgroundColor: 'rgba(25, 118, 210, 0.08)'
          },
          '& .MuiDataGrid-columnHeaders': {
            borderBottom: '1px solid rgba(224, 224, 224, 1)',
            borderTop: 'none',
            borderLeft: 'none',
            borderRight: 'none'
          },
          '& .MuiDataGrid-virtualScroller': {
            overflowY: 'scroll'
          },
          '& .MuiDataGrid-cell': {
            borderBottom: '1px solid rgba(224, 224, 224, 0.4)',
            borderTop: 'none',
            borderLeft: 'none',
            borderRight: 'none'
          }
        }}
      />
    </Paper>
  );
}
