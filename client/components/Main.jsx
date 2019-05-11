import React from 'react';
import axios from 'axios';

import Menu from './menu/Menu';

class Main extends React.Component {
  state = {
    menu: {
      name: '',
      itemList: [],
    },
  }

  componentDidMount() {
    axios.get('/api')
      .then((res) => {
        this.setState({
          menu: res.data,
        });
      });
  }

  render() {
    return (
      <div>
        <h1>It's a menu!</h1>
        <Menu menu={this.state.menu}></Menu>
      </div>
    );
  }
}

export default Main;