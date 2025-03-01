import './App.css';
import { useState } from 'react';
import BasicTextFields from './components/BasicTextFields';
import DataTable from './components/DataTable';
import BasicButton from './components/BasicButton';
import { retrieveLaunchParams } from '@telegram-apps/sdk';

function App() {
  const [tableData, setTableData] = useState([]);
  const [inputSummary, setInputSummary] = useState('');
  const [inputDescription, setInputDescription] = useState('');
  const [selectedRows, setSelectedRows] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  // Загрузка данных
  const fetchData = async () => {
    setIsLoading(true);
    try {
      const { initDataRaw } = retrieveLaunchParams();
      const response = await fetch('/api/get_tasks', {
        method: 'POST',
        headers: {
          "Authorization": `tma ${initDataRaw}`,
        },
      });
      const data = await response.json();
      setTableData(data);
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setIsLoading(false);
    }
  };


  // Отправка новой задачи
  const handleSend = async () => {
    if (!inputSummary.trim() || !inputDescription.trim()) {
      alert('Пожалуйста, заполните все поля');
      return;
    }

    try {
      const { initDataRaw } = retrieveLaunchParams();
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
      setInputDescription('');
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
          <BasicTextFields
            label='Описание'
            value={inputDescription}
            onChange={(e) => setInputDescription(e.target.value)}
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
