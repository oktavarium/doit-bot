import { retrieveRawInitData } from '@telegram-apps/sdk'
const initData = retrieveRawInitData();

const getUrl = (url) => {
    if (process.env.REACT_APP_API_URL) {
        return `${process.env.REACT_APP_API_URL}${url}`;
    }
    return url;
}

// Универсальная функция для fetch запросов
export const getWithAuth = async (url, body = null) => {
    try {
        const response = await fetch(getUrl(url), {
            method: 'GET',
            headers: {
                "Authorization": `tma ${initData}`,
                // "Authorization": "dbg 326804199",
                'Content-Type': 'application/json',
            },
            ...(body && { body: JSON.stringify(body) })
        });
        console.log("response", response);
        let responseData;
        const responseText = await response.text();
        console.log("responseText", responseText);

        try {
            responseData = responseText ? JSON.parse(responseText) : {};
            console.log("responseData", responseData);
        } catch (e) {
            console.error('Ответ не является JSON:', responseText);
            throw new Error(`Некорректный формат ответа от сервера`);
        }

        if (!response.ok) {
            console.error('Ошибка запроса:', {
                status: response.status,
                statusText: response.statusText,
                response: responseData
            });
            throw new Error(responseData?.message || `Ошибка ${response.status}: ${response.statusText}`);
        }

        return responseData || {};
    } catch (error) {
        console.error(`Ошибка при выполнении запроса к ${url}:`, error);
        // Прокидываем оригинальную ошибку дальше
        throw error.message || 'Произошла неизвестная ошибка';
    }
}

export const postWithAuth = async (url, body = null) => {
    try {
        const response = await fetch(getUrl(url), {
            method: 'POST',
            headers: {
                "Authorization": `tma ${initData}`,
                // "Authorization": "dbg 326804199",
                'Content-Type': 'application/json',
            },
            ...(body && { body: JSON.stringify(body) })
        });
        console.log("response", response);
        let responseData;
        const responseText = await response.text();
        console.log("responseText", responseText);

        try {
            responseData = responseText ? JSON.parse(responseText) : {};
            console.log("responseData", responseData);
        } catch (e) {
            console.error('Ответ не является JSON:', responseText);
            throw new Error(`Некорректный формат ответа от сервера`);
        }

        if (!response.ok) {
            console.error('Ошибка запроса:', {
                status: response.status,
                statusText: response.statusText,
                response: responseData
            });
            throw new Error(responseData?.message || `Ошибка ${response.status}: ${response.statusText}`);
        }

        return responseData || {};
    } catch (error) {
        console.error(`Ошибка при выполнении запроса к ${url}:`, error);
        // Прокидываем оригинальную ошибку дальше
        throw error.message || 'Произошла неизвестная ошибка';
    }
}

