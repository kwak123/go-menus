import React from 'react';

import { Button } from '@material-ui/core';

import Item from './Item';

import api from '../../api/api';
import MenuContext from '../../contexts/menu';

class Menu extends React.Component {
  state = {
    name: 'Loading',
    itemList: [],
  }

  componentDidMount() {
    const { menuId } = this.props;
    api.getMenu(menuId)
      .then(menu => this.setState({
        name: menu.name,
        itemList: menu.itemList,
      }));
  }

  // TODO: This refetch can be solved
  componentDidUpdate() {
    const { menuId } = this.props;
    api.getMenu(menuId)
      .then(menu => this.setState({
        name: menu.name,
        itemList: menu.itemList,
      }));
  }

  render() {
    return (
        <div className="menu__container">
          <h2>{this.state.name}</h2>
          <ul>
            {this.state.itemList.map(item => <Item item={item} key={item.id}></Item>)}
          </ul>
          <MenuContext.Consumer>
            {({ addMenuItem }) => (
              <Button variant="contained" onClick={addMenuItem}>Add Item</Button>
            )}
          </MenuContext.Consumer>
        </div>
    );
  }
};

export default Menu;
