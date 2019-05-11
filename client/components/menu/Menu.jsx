import React from 'react';
import Item from './Item';

const Menu = (props) => {
  const { menu } = props;
  const { name, itemList } = menu;

  return (
    <div>
      <h2>{name}</h2>
      <ul>
        {itemList.map(item => <Item item={item} key={item.id}></Item>)}
      </ul>
    </div>
  )
};

export default Menu;
