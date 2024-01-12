import axios from 'axios';
import { create } from 'zustand';

const useUserStore = create((set, get) => ({
  users: [],
  getAllUsers: async () => {
    const response = await axios.get('http://localhost:8080/api/users');
    set({ users: await response.data.data });
  },
  getUserById: async (id) => {
    const response = await axios.get(`http://localhost:8080/api/user?id=${id}`);
    return response.data.data;
  },
  createUser: async (data) => {
    const response = await axios.post('http://localhost:8080/api/users', data);
    return response.data.data;
  },
  updateUser: async (id, data) => {
    const response = await axios.patch(
      `http://localhost:8080/api/users?id=${id}`,
      data
    );
    return response.data.data;
  },
  deleteUser: async (id) => {
    const response = await axios.delete(
      `http://localhost:8080/api/users?id=${id}`
    );
    return response;
  },
}));

export default useUserStore;
