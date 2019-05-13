import React from 'react';

import api from '../../api/api';

import Item from './Item';

class Menu extends React.Component {
  state = {
    name: '',
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

  render() {
    return (
      <div>
        <h2>{this.state.name || "Menu"}</h2>
        <ul>
          {this.state.itemList.map(item => <Item item={item} key={item.id}></Item>)}
        </ul>
      </div>
    );
  }
};

export default Menu;
