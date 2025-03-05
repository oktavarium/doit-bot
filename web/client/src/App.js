import './App.css';
import { useState, useEffect } from 'react';
import BasicTextFields from './components/BasicTextFields';
import DataTable from './components/DataTable';
import BasicButton from './components/BasicButton';
import { retrieveRawInitData } from '@telegram-apps/sdk';
import { Button, Paper } from '@mui/material';
import vanHeader from './assets/images/vanHeader.png';

// Отключаем WebSocket для production
if (process.env.NODE_ENV === 'production') {
  window.addEventListener('error', (event) => {
    if (event.message === 'Script error.' && event.filename === '') {
      event.preventDefault();
    }
  });
}

const initData = retrieveRawInitData();
console.log("initData ====>>>>>",initData, "<<<<<<<======");

function App() {
  const [tableData, setTableData] = useState([]);
  const [inputSummary, setInputSummary] = useState('');
  const [selectedRows, setSelectedRows] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  // Загрузка данных
  const fetchData = async () => {
    setIsLoading(true);
    try {
      console.log('Fetching from:', '/api/get_tasks_by_owner');
      const response = await fetch(`${process.env.REACT_APP_API_URL}/api/get_tasks_by_owner`, {
        method: 'POST',
        headers: {
          "Authorization": `tma ${initData}`,
          // "Authorization": "dbg 326804199",
          'Content-Type': 'application/json',
        },
      });
      if (!response.ok) {
        const text = await response.text();
        console.error('Response text:', text);
        throw new Error(`HTTP error! status: ${response.status}, text: ${text}`);
      }
      const data = await response.json();
      setTableData(data.tasks);
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  // Отправка новой задачи
  const handleSend = async () => {
    // if (!inputSummary.trim()) {
    //   alert('Пожалуйста, заполните все поля');
    //   return;
    // }

    try {
      const response = await fetch(`${process.env.REACT_APP_API_URL}/api/create_task`, {
        method: 'POST',
        headers: {
          "Authorization": `tma ${initData}`,
          // "Authorization": "dbg 326804199",
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          summary: inputSummary,
          done: false,
        }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      setInputSummary('');
      setSelectedRows([]);
      fetchData();
    } catch (error) {
      console.error('Error sending data:', error);
    }
  };

  // Обработчики для кнопок меню
  const handleComplete = async () => {
    if (selectedRows.length === 0) return;
    
    // Находим выбранную задачу
    const selectedTask = tableData.find(task => task.id === selectedRows[0]);
    if (!selectedTask) return;

    try {
      console.log("selectedRows", selectedRows);
      const response = await fetch(`${process.env.REACT_APP_API_URL}/api/update_task_by_id`, {
        method: 'POST',
        headers: {
          "Authorization": `tma ${initData}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          id: selectedRows[0],
          done: !selectedTask.done, // Инвертируем текущее состояние
        }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      fetchData();
    } catch (error) {
      console.error('Error updating task:', error);
    }
  };

  const handleDelete = async () => {
    if (selectedRows.length === 0) return;
    
    try {
      console.log("selectedRows", selectedRows);
      const response = await fetch(`${process.env.REACT_APP_API_URL}/api/delete_task_by_id`, {
        method: 'POST',
        headers: {
          "Authorization": `tma ${initData}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          id: selectedRows[0]
        }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      fetchData();
    } catch (error) {
      console.error('Error deleting task:', error);
    }
  };

  return (
    <div style={{ 
      display: 'flex', 
      flexDirection: 'column', 
      height: '100vh',
      position: 'fixed',
      top: 0,
      left: 0,
      right: 0,
      bottom: 0,
      overflow: 'hidden'
    }}>
      <div className="App" style={{ 
        flex: 1,
        display: 'flex',
        flexDirection: 'column',
        overflow: 'hidden'
      }}>
        <header className="App-header" style={{ 
          height: '50px', 
          display: 'flex', 
          alignItems: 'center',
          padding: '0 20px',
          fontSize: '14px',
          position: 'relative',
          backgroundColor: '#282c34',
          flexShrink: 0
        }}>
          <div style={{ 
            display: 'flex', 
            alignItems: 'center',
            width: '100%',
            position: 'relative'
          }}>
            <img 
              src={vanHeader} 
              alt="Van Header" 
              className="rotating-image"
              style={{ 
                height: '40px',
                width: 'auto',
                objectFit: 'contain'
              }} 
            />
            <p style={{ 
              margin: 0,
              position: 'absolute',
              left: '50%',
              transform: 'translateX(-50%)',
              width: 'auto',
              color: '#ffffff'
            }}>вэнский хедер</p>
          </div>
        </header>

        <div className='inputForm' style={{
          height: '60px',
          display: 'flex',
          alignItems: 'center',
          gap: '20px',
          padding: '10px 20px',
          position: 'relative',
          backgroundColor: '#fff',
          flexShrink: 0
        }}>
          <div style={{ flex: '3' }}>
            <BasicTextFields
              label='Название'
              value={inputSummary}
              onChange={(e) => setInputSummary(e.target.value)}
            />
          </div>
          <div style={{ flex: '1' }}>
            <BasicButton 
              onClick={handleSend} 
              disabled={isLoading}
            />
          </div>
        </div>

        <div style={{ 
          flex: 1,
          overflow: 'auto',
          WebkitOverflowScrolling: 'touch'
        }}>
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
        <Paper 
          elevation={3} 
          style={{
            position: 'fixed',
            bottom: 0,
            left: 0,
            right: 0,
            padding: '1rem',
            display: 'flex',
            justifyContent: 'space-around',
            backgroundColor: '#f5f5f5',
            zIndex: 1000
          }}
        >
          <Button
            variant="contained"
            color="primary"
            onClick={handleComplete}
          >
            {tableData.find(task => task.id === selectedRows[0])?.done ? 'Отменить' : 'Выполнить'}
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
    </div>
  );
}

export default App;
