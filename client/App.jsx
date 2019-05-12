import React from 'react';

import { Router, Link } from '@reach/router';

import Main from './components/Main.jsx';
import Menu from './components/menu/Menu';

import { AppBar, Toolbar, Typography } from '@material-ui/core';

class App extends React.Component {
  render() {
    return (
      <div>
        <AppBar>
          <Toolbar>
            <Typography variant="h6" color="inherit">
              Go Menu!
            </Typography>
          </Toolbar>
        </AppBar>
        <Router>
          <Main path="/" />
          <Menu path="/app/:menuId" />
        </Router>
      </div>
    );
  }
}

export default App;
