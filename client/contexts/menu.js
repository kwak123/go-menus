import React from 'react';

const menuState = {
  menuList: [],
  currentMenuId: '',
  updateMenuItem: () => {},
  addMenuItem: () => {},
};

const MenuContext = React.createContext(menuState);

export default MenuContext;
