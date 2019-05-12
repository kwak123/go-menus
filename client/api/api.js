import axios from 'axios';

const api = axios.create();

api.interceptors.response.use(res => res.data);

const getAllMenus = () => api.get('/api')

const getMenu = menuId => api.get(`/api/${menuId}`);

export default {
  getAllMenus,
  getMenu,
};
