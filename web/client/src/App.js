import './App.css';
import { useState, useEffect } from 'react';
import BasicTextFields from './components/BasicTextFields';
import DataTable from './components/DataTable';
import BasicButton from './components/BasicButton';
import { retrieveLaunchParams, init, miniApp } from '@telegram-apps/sdk';

function App() {
  const [tableData, setTableData] = useState([]);
  const [inputSummary, setInputSummary] = useState('');
  const [inputDescription, setInputDescription] = useState('');
  const [selectedRows, setSelectedRows] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [initDataRaw, setInitDataRaw] = useState('');

  // Инициализация Telegram SDK и получение initDataRaw
  useEffect(() => {
    const initialize = async () => {
      try {
        await init();

        if (miniApp.ready.isAvailable()) {
          await miniApp.ready();
        }

        // Получаем initDataRaw после инициализации
        const { initDataRaw: rawData } = retrieveLaunchParams();
        if (rawData) {
          setInitDataRaw(rawData);
        } else {
          console.error('initDataRaw is empty');
        }
      } catch (error) {
        console.error('Ошибка инициализации:', error);
      }
    };

    initialize();
  }, []);

  // Загрузка данных
  const fetchData = async () => {
    if (!initDataRaw) {
      console.error('initDataRaw is empty, skipping fetchData');
      return;
    }

    setIsLoading(true);
    try {
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

  // Выполняем fetchData при изменении initDataRaw
  useEffect(() => {
    if (initDataRaw) {
      fetchData();
    }
  }, [initDataRaw]);

  // Отправка новой задачи
  const handleSend = async () => {
    if (!inputSummary.trim() || !inputDescription.trim()) {
      alert('Пожалуйста, заполните все поля');
      return;
    }

    if (!initDataRaw) {
      console.error('initDataRaw is empty, cannot send data');
      return;
    }

    try {
      await fetch('/api/create_task', {
        method: 'POST',
        headers: {
          "Authorization": `tma ${initDataRaw}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          summary: inputSummary,
          description: inputDescription,
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