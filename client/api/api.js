import axios from 'axios';

const api = axios.create();

api.interceptors.response.use(res => res.data);

const getAllMenus = () => api.get('/api/')

const getMenu = menuId => api.get(`/api/${menuId}`);

const updateMenuItem = (menuId, item) => api.put(`/api/${menuId}`, item);

const addMenuItem = () => api.post('/api/add', {});

const deleteMenuItem = itemId => api.post('/api/delete', { id: itemId });

export default {
  getAllMenus,
  getMenu,
  updateMenuItem,
  addMenuItem,
  deleteMenuItem,
};
