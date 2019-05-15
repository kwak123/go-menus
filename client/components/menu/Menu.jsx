import React from 'react';

import { Button } from '@material-ui/core';

import Item from './Item';

const Menu = (props) => {
  const { currentMenu, addMenuItem } = props;
  const { name, location, itemList } = currentMenu;

  return (
    <div className="menu__container">
      <h2>{name}</h2>
      <h3>{location}</h3>
      <ul>
        {itemList.map(item => <Item item={item} key={item.id}></Item>)}
      </ul>
      <Button variant="contained" onClick={addMenuItem}>Add Item</Button>
    </div>
  );
};

export default Menu;
