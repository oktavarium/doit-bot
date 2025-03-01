import './App.css';
import { useState, useEffect } from 'react';
import BasicTextFields from './components/BasicTextFields';
import DataTable from './components/DataTable';
import BasicButton from './components/BasicButton';
import { retrieveRawInitData } from '@telegram-apps/sdk';

function App() {
  const [tableData, setTableData] = useState([]);
  const [inputSummary, setInputSummary] = useState('');
  const [selectedRows, setSelectedRows] = useState([]);
  const [isLoading, setIsLoading] = useState(false);



  // Загрузка данных
  const fetchData = async () => {
    setIsLoading(true);
    try {
      const initDataRaw  = retrieveRawInitData();
      const response = await fetch('/api/get_tasks', {
        method: 'POST',
        headers: {
          "Authorization": `tma ${initDataRaw}`,
        },
      });
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
    if (!inputSummary.trim()) {
      alert('Пожалуйста, заполните все поля');
      return;
    }

    try {
      const initDataRaw  = retrieveRawInitData();
      await fetch('/api/create_task', {
        method: 'POST',
        headers: {
          "Authorization": `tma ${initDataRaw}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          summary: inputSummary,
          done: false,
        }),
      });

      setInputSummary('');
      setSelectedRows([]);
      fetchData(); // Обновляем данные после отправки
    } catch (error) {
      console.error('Error sending data:', error);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <p>Тут типа хедер с кнопками</p>
      </header>
      <div className='inputForm'>
        <div className='inputFields'>
          <BasicTextFields
            label='Название'
            value={inputSummary}
            onChange={(e) => setInputSummary(e.target.value)}
          />
        </div>
        <BasicButton onClick={handleSend} disabled={isLoading} />
      </div>
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
  );
}

export default App;
