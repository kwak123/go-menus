import React from 'react';

import { Router, Link } from '@reach/router';

import Main from './components/Main.jsx';
import Menu from './components/menu/Menu';

class App extends React.Component {
  render() {
    return (
      <Router>
        <Main path="/" />
        <Menu path="/app/:menuId" />
      </Router>
    );
  }
}

export default App;
