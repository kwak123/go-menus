import React from 'react';

import api from '../../api/api';

const Item = (props) => {
  const { item } = props;
  const { name, provider } = item;

  return (
    <li>
      <h3>Item: {name}</h3>
      <p>Provider: {provider}</p>
    </li>
  );
};

export default Item;
