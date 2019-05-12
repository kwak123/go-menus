import React from 'react';
import axios from 'axios';

import Menu from './menu/Menu';

import api from '../api/api'

class Main extends React.Component {
  state = {
    menuList: [],
  }

  componentDidMount() {
    api.getAllMenus()
      .then(menuList => this.setState({ menuList }));
  }

  render() {
    return (
      <div className="main-page">
        <h1 className="main-page__header">Menus!</h1>
        {this.state.menuList.map(menu => <Menu menu={menu} key={menu.id}></Menu>)}
      </div>
    );
  }
}

export default Main;