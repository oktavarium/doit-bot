import { API_URL } from "./constants";
import { getWithAuth, postWithAuth, patchWithAuth, deleteWithAuth} from "../utils/api.utils";

export const getAllTasks = () => {
    return getWithAuth(API_URL.getAllTasks);
}

export const createTask = (inputSummary) => {
    return postWithAuth(API_URL.createTask, {
        name: inputSummary,
        status: false,
    });
}

export const updateTask = (selectedRows, selectedTask) => {
    return patchWithAuth(API_URL.updateTask + "/" + selectedRows[0], {
        status: !selectedTask.status,
    });
}

export const deleteTask = (selectedRows) => {
    return deleteWithAuth(API_URL.deleteTask+"/" + selectedRows[0], {
    });
}
