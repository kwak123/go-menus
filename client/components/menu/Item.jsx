import React, { useState } from 'react';

import MenuContext from '../../contexts/menu';

import { Button } from '@material-ui/core';

const Item = (props) => {
  const { item } = props;

  const [itemName, setItemName] = useState(item.name);
  const [itemProvider, setItemProvider] = useState(item.provider)

  return (
    <MenuContext.Consumer>
      {({ updateMenuItem, deleteMenuItem }) => (
        <li
          className="menu-item"
          onBlur={() => updateMenuItem({
            id: item.id,
            name: itemName,
            provider: itemProvider,
          })}
        >
          <div className="menu-item__name-field">
            <p>Item: </p>
            <input
              type="text"
              value={itemName}
              onChange={e => setItemName(e.target.value)}
            />
          </div>
          <div className="menu-item__provider-field">
            <p>Provider: </p>
            <input
              type="text"
              value={itemProvider}
              onChange={e => setItemProvider(e.target.value)}
            />
          </div>
          <Button
            onClick={() => deleteMenuItem(item)}
          >
            Delete
          </Button>
        </li>
      )}
    </MenuContext.Consumer>
  );
};

export default Item;
