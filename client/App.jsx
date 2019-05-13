import React from 'react';

import { Router, Link } from '@reach/router';

import Main from './components/Main.jsx';

import { AppBar, Toolbar, Typography } from '@material-ui/core';

class App extends React.Component {
  render() {
    return (
      <div class="app__container">
        <AppBar>
          <Toolbar>
            <Typography variant="h6" color="inherit">
              <Link to="/" className="header__title">Go Menu!</Link>
            </Typography>
          </Toolbar>
        </AppBar>
        <Main />
      </div>
    );
  }
}

export default App;
