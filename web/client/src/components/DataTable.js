import * as React from 'react';
import { DataGrid } from '@mui/x-data-grid';
import Paper from '@mui/material/Paper';

const columns = [
  { field: 'done', headerName: 'Выполнено', type: 'boolean',  flex: 1 },
  { field: 'summary', headerName: 'Задача', type: 'string', flex: 2 },
];

const paginationModel = { page: 0, pageSize: 5 };

export default function DataTable({ rows, selectedRows, onSelectionChange }) {
  return (
    <Paper sx={{
      height: 400,
      width: '60%',
      margin: '20px auto',
      padding: 2
    }}>
      <DataGrid
        rows={rows}
        columns={columns}
        initialState={{ pagination: { paginationModel } }}
        pageSizeOptions={[5]}
        rowsPerPageOptions={[5]}
        onRowSelectionModelChange={onSelectionChange}
        rowSelectionModel={selectedRows}
        sx={{
          border: 0,
          '& .MuiDataGrid-cell:focus': {
            outline: 'none'
          }
        }}
      />
    </Paper>
  );
}
