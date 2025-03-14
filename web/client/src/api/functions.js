import { API_URL } from "./constants";
import { getWithAuth, postWithAuth } from "../utils/api.utils";

export const getAllTasks = () => {
    return getWithAuth(API_URL.getAllTasks);
}

export const createTask = (inputSummary) => {
    return postWithAuth(API_URL.createTask, {
        name: inputSummary,
        done: false,
    });
}

export const updateTask = (selectedRows, selectedTask) => {
    return postWithAuth(API_URL.updateTask, {
        id: selectedRows[0],
        done: !selectedTask.done,
    });
}

export const deleteTask = (selectedRows) => {
    return postWithAuth(API_URL.deleteTask, {
        id: selectedRows[0]
    });
}
