import axios from 'axios';

const api = axios.create();

api.interceptors.response.use(res => res.data);

const getAllMenus = () => api.get('/api/')

const getMenu = menuId => api.get(`/api/${menuId}`);

const updateMenuItem = (menuId, item) => api.put(`/api/${menuId}`, item);

export default {
  getAllMenus,
  getMenu,
  updateMenuItem,
};
