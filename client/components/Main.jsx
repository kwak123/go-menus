import React from 'react';
import { Router, Link } from '@reach/router';

import MenuContext from '../contexts/menu';

import HomePage from './common/HomePage';
// import Menu from './menu/Menu';
import MenuContainer from './menu/MenuContainer';

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
    currentMenu: {
      id: '',
      name: '',
      itemList: [],
    },
    updateMenuItem: item => this.updateMenuItem(item),
    addMenuItem: () => this.addMenuItem(),
    deleteMenuItem: (itemId) => this.deleteMenuItem(itemId),
  }

  updateMenuItem = (item) => {
    const { currentMenuId } = this.state;
    return api.updateMenuItem(currentMenuId, item)
      .then((updatedMenu) => {
        // Ghettoooooo
        const newMenuList = this.state.menuList.map((menu) => {
          if (menu.id === updatedMenu.id) {
            return updatedMenu;
          }
          return { ...menu };
        });
        this.setState({ menuList: newMenuList })
      });
  }

  componentDidMount() {
    api.getAllMenus()
      .then((menuList) => this.refreshMenuData(menuList));
  }

  refreshMenuData = (menuList) => {
    const urlArray = window.location.href.split('/');
    const currentMenuId = urlArray.pop();
    const currentMenu = menuList.find(menu => menu.id === currentMenuId);
    this.setState({
      menuList,
      currentMenuId,
      currentMenu,
    });
  }

  addMenuItem = () => {
    return api.addMenuItem()
      .then((updatedMenu) => {
        // Ghettoooooo
        const newMenuList = this.state.menuList.map((menu) => {
          if (menu.id === updatedMenu.id) {
            return updatedMenu;
          }
          return { ...menu };
        });
        this.refreshMenuData(newMenuList);
      });
  }

  deleteMenuItem = (itemId) => {
    return api.deleteMenuItem(itemId)
      .then((updatedMenu) => {
        // Ghettoooooo
        const newMenuList = this.state.menuList.map((menu) => {
          if (menu.id === updatedMenu.id) {
            return updatedMenu;
          }
          return { ...menu };
        });
        this.refreshMenuData(newMenuList);
      });
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
                <MenuContainer path="/app/:menuId" />
              </Router>
            </div>
          </div>
        </div>
      </MenuContext.Provider>
    );
  }
}

export default Main;
