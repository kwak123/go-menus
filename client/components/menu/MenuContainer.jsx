import React from 'react';
import MenuContext from '../../contexts/menu';
import Menu from './Menu';

const MenuContainer = (menuId) => (
  <MenuContext.Consumer>
    {(context) => <Menu {...context} menuId={menuId} />}
  </MenuContext.Consumer>
);

export default MenuContainer;
