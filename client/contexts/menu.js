import React from 'react';

const menuState = {
  menuList: [],
  currentMenuId: '',
  updateMenuItem: () => {},
};

const MenuContext = React.createContext(menuState);

export default MenuContext;
