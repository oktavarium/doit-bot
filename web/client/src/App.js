import './App.scss';
import { useState, useEffect, useCallback } from 'react';
import BasicTextFields from './components/BasicTextField/BasicTextFields';
import DataTable from './components/DataTable/DataTable';
import BasicButton from './components/BasicButton/BasicButton';
import { Button, Paper, Snackbar, Alert } from '@mui/material';
import { getAllTasks, createTask, updateTask, deleteTask } from './api/functions';

// Отключаем WebSocket для production
if (process.env.NODE_ENV === 'production') {
  window.addEventListener('error', (event) => {
    if (event.message === 'Script error.' && event.filename === '') {
      event.preventDefault();
    }
  });
}

function App() {
  const [tableData, setTableData] = useState([]);
  const [inputSummary, setInputSummary] = useState('');
  const [selectedRows, setSelectedRows] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleError = useCallback((error) => {
    setError(error.toString());
    setTimeout(() => setError(null), 3000);
  }, []); // Нет зависимостей, так как setError стабилен

  // Загрузка данных
  const fetchData = useCallback(async () => {
    setIsLoading(true);
    try {
      const data = await getAllTasks();
      setTableData(data?.tasks || []);
    } catch (error) {
      handleError(error);
      setTableData([]); // Устанавливаем пустой массив в случае ошибки
    } finally {
      setIsLoading(false);
    }
  }, [handleError]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  // Отправка новой задачи
  const handleSend = async () => {
    if (!inputSummary.trim()) return; // Проверка на пустой ввод
    
    try {
      await createTask(inputSummary);
      setInputSummary('');
      setSelectedRows([]);
      fetchData();
    } catch (error) {
      handleError(error);
    }
  };

  // Обработчики для кнопок меню
  const handleComplete = async () => {
    if (selectedRows.length === 0) return;
    
    const selectedTask = tableData.find(task => task.id === selectedRows[0]);
    if (!selectedTask) return;

    try {
      await updateTask(selectedRows, selectedTask);
      fetchData();
    } catch (error) {
      handleError(error);
    }
  };

  const handleDelete = async () => {
    if (selectedRows.length === 0) return;
    
    try {
      await deleteTask(selectedRows);
      fetchData();
    } catch (error) {
      handleError(error);
    }
  };

  // Получаем статус выбранной задачи безопасно
  const getSelectedTaskStatus = () => {
    const selectedTask = tableData.find(task => task.id === selectedRows[0]);
    return selectedTask?.done ?? false;
  };

  return (
    <div className="app-container">
      <div className="app-content">
        <header className="app-header">
          <div className="app-header__container">
            <p className="app-header__title"></p>
          </div>
        </header>

        <div className="input-form">
          <div className="input-form__input-container">
            <BasicTextFields
              label='Название'
              value={inputSummary}
              onChange={(e) => setInputSummary(e.target.value)}
            />
          </div>
          <div className="input-form__button-container">
            <BasicButton 
              onClick={handleSend} 
              disabled={isLoading || !inputSummary.trim() || inputSummary.length > 30}
            />
          </div>
        </div>

        <div className="table-container">
          {isLoading ? (
            <p>Загрузка...</p>
          ) : (
            <DataTable
              rows={tableData}
              selectedRows={selectedRows}
              onSelectionChange={setSelectedRows}
            />
          )}
        </div>
      </div>
      
      {selectedRows.length > 0 && (
        <Paper className="bottom-menu" elevation={3}>
          <Button
            variant="contained"
            color="primary"
            onClick={handleComplete}
          >
            {getSelectedTaskStatus() ? 'Отменить' : 'Выполнить'}
          </Button>
          <Button
            variant="contained"
            color="error"
            onClick={handleDelete}
          >
            Удалить
          </Button>
        </Paper>
      )}

      <Snackbar 
        open={!!error} 
        autoHideDuration={3000} 
        onClose={() => setError(null)}
      >
        <Alert severity="error" onClose={() => setError(null)}>
          {error}
        </Alert>
      </Snackbar>
    </div>
  );
}

export default App;
