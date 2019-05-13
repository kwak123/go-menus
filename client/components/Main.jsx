import React from 'react';
import { Router, Link } from '@reach/router';

import MenuContext from '../contexts/menu';

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
    menuList: [],
    currentMenuId: '',
    updateMenuItem: item => this.updateMenuItem(item),
  }

  updateMenuItem = (item) => {
    const { currentMenuId } = this.state;
    return api.updateMenuItem(currentMenuId, item)
      .then(updatedMenu => {
        // Ghettoooooo
        const newMenuList = this.state.menuList.map((menu) => {
          if (menu.id === updatedMenu.id) {
            console.log('found');
            return updatedMenu;
          }
          return { ...menu };
        });
        console.log(item)
        console.log(JSON.parse(JSON.stringify(newMenuList)));
        this.setState({ menuList: newMenuList })
      });
  }

  componentDidMount() {
    this.setCurrentMenuId();
    api.getAllMenus()
      .then(menuList => this.setState({ menuList }));
  }

  setCurrentMenuId = () => {
    const urlArray = window.location.href.split('/');
    const menuId = urlArray.pop();
    this.setState({ currentMenuId: menuId });
  }

  render() {
    return (
      <MenuContext.Provider value={this.state}>
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
                    key={menu.id}
                    selected={this.state.currentMenuId === menu.id}
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
      </MenuContext.Provider>
    );
  }
}

export default Main;
