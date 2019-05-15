import axios from 'axios';

const api = axios.create();

api.interceptors.response.use(res => res.data);

const getAllMenus = () => api.get('/api/')

const getMenu = menuId => api.get(`/api/${menuId}`);

const updateMenuItem = (menuId, item) => api.put(`/api/`, { menuId, item });

const addMenuItem = (menuId) => api.post('/api/add', { menuId });

const deleteMenuItem = (menuId, item) => api.post('/api/delete', { menuID: menuId, item  });

export default {
  getAllMenus,
  getMenu,
  updateMenuItem,
  addMenuItem,
  deleteMenuItem,
};
