import React from 'react';
import { Router, Link } from '@reach/router';

import HomePage from './common/HomePage';
import Menu from './menu/Menu';

import {
  List,
  ListItem,
  ListItemText,
} from '@material-ui/core';

import api from '../api/api'

class Main extends React.Component {
  state = {
    name: 'Loading',
    menuList: [],
  }

  componentDidMount() {
    api.getAllMenus()
      .then(menuList => this.setState({ menuList }));
  }

  render() {
    return (
      <div className="main-page">
        <div className="main-page__container">
          <div className="menu-list">
            <h2 className="menu-list__header">Menus</h2>
            <List component="nav">
              {this.state.menuList.map(menu => (
                <ListItem
                  button
                  component="a"
                  href={`/app/${menu.id}`}
                >
                  <ListItemText>{menu.name}</ListItemText>
                </ListItem>)
              )}
            </List>
          </div>
          <div className="main-page__detail">
            <Router>
              <HomePage path="/" />
              <Menu path="/app/:menuId" />
            </Router>
          </div>
        </div>
      </div>
    );
  }
}

export default Main;
