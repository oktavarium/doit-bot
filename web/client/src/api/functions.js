import { API_URL } from "./constants";
import { fetchWithAuth } from "../utils/api.utils";

export const getAllTasks = () => {
    return fetchWithAuth(API_URL.getAllTasks);
}

export const createTask = (inputSummary) => {
    return fetchWithAuth(API_URL.createTask, {
        name: inputSummary,
        done: false,
    });
}   

export const updateTask = (selectedRows, selectedTask) => {
    return fetchWithAuth(API_URL.updateTask, {
        id: selectedRows[0],
        done: !selectedTask.done,
    });
}

export const deleteTask = (selectedRows) => {
    return fetchWithAuth(API_URL.deleteTask, {
        id: selectedRows[0]
    });
}
