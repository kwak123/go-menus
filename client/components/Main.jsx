import React from 'react';
import axios from 'axios';

import api from '../api/api'

class Main extends React.Component {
  state = {
    name: 'Loading',
    menuList: [],
  }

  componentDidMount() {
    api.getAllMenus()
      .then(menuList => {
        console.log(menuList)
        this.setState({ menuList })
      });
  }

  render() {
    return (
      <div className="main-page">
        <h1 className="main-page__header">Menus!</h1>
        {this.state.menuList.map(menu => <li><a href={`/app/${menu.id}`}>{menu.name}</a></li>)}
      </div>
    );
  }
}

export default Main;