import React from 'react';

import MenuContext from '../../contexts/menu';

const Item = (props) => {
  const { item } = props;
  const { name, provider } = item;

  return (
    <MenuContext.Consumer>
      {({ updateMenuItem }) => (
        <li className="menu-item" onBlur={() => console.log('test')}>
          <div className="menu-item__name-field">
            <p>Item: </p>
            <input type="text" value={item.name}></input>
          </div>
          <div className="menu-item__provider-field">
            <p>Provider: </p>
            <input type="text" value={item.provider}></input>
          </div>
        </li>
      )}
    </MenuContext.Consumer>
  );
};

export default Item;
