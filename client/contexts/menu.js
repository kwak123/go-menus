import React from 'react';

const menuState = {
  menuList: [],
  currentMenuId: '',
  currentMenu: {
    id: '',
    name: '',
    itemList: [],
  },
  updateMenuItem: () => {},
  addMenuItem: () => {},
  deleteMenuItem: () => {},
};

const MenuContext = React.createContext(menuState);

export default MenuContext;
